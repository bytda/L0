package main

import "github.com/nats-io/stan.go"

func main() {

	sc, _ := stan.Connect("test-cluster", "client-1")
	defer sc.Close()
	sc.Publish("foo", []byte(model1))
	sc.Publish("foo", []byte(model2))
	sc.Publish("foo", []byte("{invalid_model}"))
	sc.Publish("foo", []byte(model3))
	sc.Publish("foo", []byte("invalid_model+another_invalid_model"))
	sc.Publish("foo", []byte(model4))

}

var model1 = `{
  "order_uid": "model1",
  "track_number": "WB",
  "entry": "WBIL",
  "delivery": {
    "name": "Test Testov",
    "phone": "+9720000000",
    "zip": "2639809",
    "city": "Kiryat Mozkin",
    "address": "Ploshad Mira 15",
    "region": "Kraiot",
    "email": "test@gmail.com"
  },
  "payment": {
    "transaction": "model1",
    "request_id": "",
    "currency": "USD",
    "provider": "wbpay",
    "amount": 1817,
    "payment_dt": 1637907727,
    "bank": "alpha",
    "delivery_cost": 1500,
    "goods_total": 317,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 9934930,
      "track_number": "WB",
      "price": 453,
      "rid": "ab4219087a764ae0btest",
      "name": "Mascaras",
      "sale": 30,
      "size": "0",
      "total_price": 317,
      "nm_id": 2389212,
      "brand": "Vivienne Sabo",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "test",
  "delivery_service": "meest",
  "shardkey": "9",
  "sm_id": 99,
  "date_created": "2021-11-26T06:22:19Z",
  "oof_shard": "1"
}
`

var model2 = `{
"order_uid": "model2",
"track_number": "test_track_number",
"entry": "WBIL",
"delivery": {
    "name": "Test Testov",
    "phone": "+9720000000",
    "zip": "2639809",
    "city": "Kiryat Mozkin",
    "address": "Ploshad Mira 15",
    "region": "Kraiot",
    "email": "test@gmail.com"
},
"payment": {
    "transaction": "model2",
    "request_id": "",
    "currency": "USD",
    "provider": "wbpay",
    "amount": 1817,
    "payment_dt": 1637907727,
    "bank": "alpha",
    "delivery_cost": 1500,
    "goods_total": 317,
    "custom_fee": 0
},
"items": [
    {
    "chrt_id": 9934930,
    "track_number": "test_track_number",
    "price": 453,
    "rid": "ab4219087a764ae0btest",
    "name": "Mascaras",
    "sale": 30,
    "size": "0",
    "total_price": 317,
    "nm_id": 2389212,
    "brand": "Vivienne Sabo",
    "status": 202
    }
],
"locale": "en",
"internal_signature": "",
"customer_id": "test",
"delivery_service": "meest",
"shardkey": "9",
"sm_id": 99,
"date_created": "2021-11-26T06:22:19Z",
"oof_shard": "1"
}`
var model3 = `{
  "order_uid": "model3",
  "track_number": "dsfgdsfg",
  "entry": "WBIL",
  "delivery": {
      "name": "Test2 Testov2",
      "phone": "+9720000000",
      "zip": "2639809",
      "city": "Kiryat Mozkin",
      "address": "Ploshad Mira 15",
      "region": "Kraiot",
      "email": "test@gmail.com"
  },
  "payment": {
      "transaction": "model3",
      "request_id": "",
      "currency": "USD",
      "provider": "wbpay",
      "amount": 1817,
      "payment_dt": 1637907727,
      "bank": "alpha",
      "delivery_cost": 1500,
      "goods_total": 317,
      "custom_fee": 0
  },
  "items": [
      {
      "chrt_id": 9934930,
      "track_number": "dsfgdsfg",
      "price": 453,
      "rid": "ab4219087a764ae0btest",
      "name": "Mascaras",
      "sale": 30,
      "size": "0",
      "total_price": 317,
      "nm_id": 2389212,
      "brand": "Vivienne Sabo",
      "status": 202
      }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "test",
  "delivery_service": "meest",
  "shardkey": "9",
  "sm_id": 99,
  "date_created": "2021-11-26T06:22:19Z",
  "oof_shard": "1"
  }`

var model4 = `{
    "order_uid": "model4",
    "track_number": "WBILMTESTTRACK",
    "entry": "WBIL",
    "delivery": {
        "name": "Test Testov",
        "phone": "+9720000000",
        "zip": "2639809",
        "city": "Kiryat Mozkin",
        "address": "Ploshad Mira 15",
        "region": "Kraiot",
        "email": "test@gmail.com"
    },
    "payment": {
        "transaction": "model4",
        "request_id": "",
        "currency": "USD",
        "provider": "wbpay",
        "amount": 1817,
        "payment_dt": 1637907727,
        "bank": "alpha",
        "delivery_cost": 1500,
        "goods_total": 317,
        "custom_fee": 0
    },
    "items": [
        {
        "chrt_id": 1,
        "track_number": "WBILMTESTTRACK",
        "price": 453,
        "rid": "ab4219087a764ae0btest",
        "name": "whatever",
        "sale": 30,
        "size": "0",
        "total_price": 317,
        "nm_id": 2389212,
        "brand": "Vivienne Sabo",
        "status": 202
        },
        {
        "chrt_id": 666,
        "track_number": "WBILMTESTTRACK",
        "price": 453,
        "rid": "ab4219087a764ae0btest",
        "name": "WhiteHat",
        "sale": 30,
        "size": "0",
        "total_price": 317,
        "nm_id": 2389212,
        "brand": "Vivienne Sabo",
        "status": 202
        },
        {
        "chrt_id": 431,
        "track_number": "WBILMTESTTRACK",
        "price": 453,
        "rid": "ab4219087a764ae0btest",
        "name": "RedHat",
        "sale": 30,
        "size": "0",
        "total_price": 317,
        "nm_id": 2389212,
        "brand": "Vivienne Sabo",
        "status": 202
        },
        {
        "chrt_id": 123,
        "track_number": "WBILMTESTTRACK",
        "price": 453,
        "rid": "ab4219087a764ae0btest",
        "name": "BlackHat",
        "sale": 30,
        "size": "0",
        "total_price": 317,
        "nm_id": 2389212,
        "brand": "Vivienne Sabo",
        "status": 202
        }
    ],
    "locale": "en",
    "internal_signature": "",
    "customer_id": "test",
    "delivery_service": "meest",
    "shardkey": "9",
    "sm_id": 99,
    "date_created": "2021-11-26T06:22:19Z",
    "oof_shard": "1"
    }`
