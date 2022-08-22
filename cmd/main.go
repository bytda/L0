package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/nats-io/stan.go"
	"github.com/patrickmn/go-cache"
)

func main() {

	db := fmt.Sprintf("postgres://localhost:5432/postgres") //  подключение к базе данных

	conn, err := pgx.Connect(context.Background(), db)

	_, err = conn.Exec(context.Background(), sqlconfig) // очищаем базу данных

	log.Println("Проверка подключения к базе данных")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Соединение не установлено: %v\n", err)
		os.Exit(1)
	} else {
		log.Println("Соединение установлено")
	}

	defer conn.Close(context.Background())

	// Получаем все order_uid из БД и создаем кэш
	order_uid := GetOrderUid(conn)
	Cache := cache.New(-1, -1) // -1 - чтобы кэш не удалялся сам
	for i := range order_uid {
		data, _ := GetDataByUid(conn, order_uid[i])
		Cache.Set(order_uid[i], data, cache.NoExpiration)
	}
	log.Printf("Восстанавливаем данные из базы данных, количество элементов (%v)\n", len(Cache.Items()))

	sc, _ := stan.Connect("test-cluster", "client-1") // подключаемся к каналу натс
	defer sc.Close()

	var order OrderInfo

	sc.Subscribe("foo", func(m *stan.Msg) {
		err := json.Unmarshal(m.Data, &order)
		if err != nil {
			log.Println("Получены некорректные данные")
			InsertInvalidData(conn, string(m.Data))
		} else {
			log.Println("Получены корректные данные")
			Cache.Set(order.OrderUid, string(m.Data), cache.NoExpiration) //записываем данные в кэш
			log.Println("Новые данные записаны")
			InsertData(conn, order)
		}
	}, stan.DeliverAllAvailable())

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	log.Println("Http сервер запущен, чтобы подключиться - http://localhost:8080")
	r.LoadHTMLGlob("templates/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": "This is an index page...",
		})
	})
	r.POST("/result", func(c *gin.Context) {
		result, _ := Cache.Get(c.PostForm("order_uid"))

		if result == nil {
			c.PureJSON(http.StatusOK, "Записей по такому order_uid не найдено")
		} else {
			c.PureJSON(http.StatusOK, result)
		}
	})
	r.Run(":8080")

}
