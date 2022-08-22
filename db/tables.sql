drop table if exists items,order_info, payments,deliveries, order_delivery, invalid_data;



create table deliveries
(
    id      serial primary key,
    name    varchar(30),
    phone   varchar(20),
    zip     varchar(20),
    city    varchar(20),
    address varchar(25),
    region  varchar(20),
    email   varchar(30)
);



create table payments
(
    transaction   varchar(50) primary key,
    request_id    varchar(50),
    currency      varchar(10),
    provider      varchar(20),
    amount        integer,
    payment_dt    bigint,
    bank          varchar(30),
    delivery_cost integer,
    goods_total   integer,
    custom_fee    integer

);

create table order_info
(
    order_uid          varchar(50) primary key references payments (transaction),
    track_number       varchar(50) unique,
    entry              varchar(20),
    locale             varchar(10),
    internal_signature varchar(50),
    customer_id        varchar(20),
    delivery_service   varchar(20),
    shardkey           varchar(15),
    sm_id              integer,
    date_created       timestamp,
    oof_shard          varchar(15)
);



create table items
(
    chrt_id      integer,
    track_number varchar(50) references order_info (track_number),
    price        integer,
    rid          varchar(50),
    name         varchar(20),
    sale         integer,
    size         varchar(10),
    total_price  integer,
    nm_id        integer,
    brand        varchar(20),
    status       integer
);

create table order_delivery
(
    order_uid   varchar(50) references order_info (order_uid),
    delivery_id int references deliveries (id)
);

create table invalid_data
(
    id        serial primary key,
    data      varchar,
    timestamp timestamp default now()
);