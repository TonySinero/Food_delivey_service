CREATE TABLE order_status
(
    id          INT PRIMARY KEY,
    description VARCHAR
);

CREATE TABLE orders
(
    id                UUID             PRIMARY KEY       DEFAULT gen_random_uuid(),
    customer_id       INT              NOT NULL,
    restaurant_id     UUID             NOT NULL,
    courier_agency_id UUID,
    courier_id        UUID,
    payment_id        UUID,
    status_id         INT              NOT NULL,
    address           VARCHAR          NOT NULL,
    cost              DECIMAL,
    required_time     TIMESTAMP,
    fact_time         TIMESTAMP,
    created_at        TIMESTAMP        NOT NULL DEFAULT (now() AT TIME ZONE 'utc')
);

CREATE TABLE order_dishes
(
    id         UUID PRIMARY KEY                                       DEFAULT gen_random_uuid(),
    order_id   UUID references orders (id) on delete cascade not null,
    dish_id    UUID                                          NOT NULL,
    amount     INT,
    created_at TIMESTAMP                                     NOT NULL DEFAULT (now() AT TIME ZONE 'utc')
);

CREATE TABLE order_feedback
(
    id         UUID PRIMARY KEY                                       DEFAULT gen_random_uuid(),
    order_id   UUID references orders (id) on delete cascade not null,
    feedback   VARCHAR                                       NOT NULL,
    rating     integer                                       NOT NULL,
    created_at TIMESTAMP                                     NOT NULL DEFAULT (now() AT TIME ZONE 'utc')
);


CREATE TABLE customers
(
    id           INT PRIMARY KEY NOT NULL,
    full_name    VARCHAR          NOT NULL,
    address      VARCHAR,
    phone_number NUMERIC          NOT NULL,
    birthday     DATE
);

INSERT INTO order_status(id, description)
VALUES (1, 'New'),
       (2, 'In progress'),
       (3, 'Ready for delivery'),
       (4, 'Completed'),
       (5, 'Canceled');

INSERT INTO customers(id, full_name, address, phone_number, birthday)
VALUES ('1', 'Иван Иванов', 'ул. Новоуфимская 11, Минск', '375291111111', '1955-12-17'),
       ('2', 'Петя Петров', 'ул. Соломенная 23, Минск', '375291111112', '2000-12-17'),
       ('3', 'Катя Костевич', 'ул. Мелиоративная 38, Минск', '375291111113',
        '1967-11-05'),
       ('4', 'Зоя Заидовна', 'пер. Собинова 24, Минск', '375291111114', '1989-02-11'),
       ('5', 'Кирилл Кирилловив', 'пр. Сморговский 59, Минск', '375291111115',
        '2001-06-10');

INSERT INTO orders (id, customer_id, restaurant_id, status_id, address, cost, required_time)
VALUES ('01fb44e3-5f18-0000-80a1-d8b4e8a22f1b', '1',
        '02fb44e3-5f18-45eb-80a1-d8b4e8a22f1b', 1, 'ул. Новоуфимская 11, Минск', 32, '2016-06-22 19:10:25-07'),
       ('01fb44e3-5f18-0001-80a1-d8b4e8a22f1b', '2',
        '02fb44e3-5f18-45eb-80a1-d8b4e8a22f1b', 1, 'улица Притыцкого 29', 15, '2020-06-22 19:10:25-07'),
       ('01fb44e3-5f18-0002-80a1-d8b4e8a22f1b','3',
        '02fb44e3-5f18-45eb-80a1-d8b4e8a22f1b', 1, 'улица Притыцкого 30', 15.3, '2022-02-22 19:10:25-07');

INSERT INTO order_dishes (order_id, dish_id, amount)
VALUES ('01fb44e3-5f18-0000-80a1-d8b4e8a22f1b', '01fb44e3-0000-45eb-80a1-d8b4e8a22f1b', 2),
       ('01fb44e3-5f18-0000-80a1-d8b4e8a22f1b', '01fb44e3-0001-45eb-80a1-d8b4e8a22f1b', 2),
       ('01fb44e3-5f18-0001-80a1-d8b4e8a22f1b', '01fb44e3-0002-45eb-80a1-d8b4e8a22f1b', 3),
       ('01fb44e3-5f18-0001-80a1-d8b4e8a22f1b', '01fb44e3-0003-45eb-80a1-d8b4e8a22f1b', 1),
       ('01fb44e3-5f18-0002-80a1-d8b4e8a22f1b', '01fb44e3-0004-45eb-80a1-d8b4e8a22f1b', 1),
       ('01fb44e3-5f18-0002-80a1-d8b4e8a22f1b', '01fb44e3-0005-45eb-80a1-d8b4e8a22f1b', 1),
       ('01fb44e3-5f18-0002-80a1-d8b4e8a22f1b', '01fb44e3-0006-45eb-80a1-d8b4e8a22f1b', 1);
