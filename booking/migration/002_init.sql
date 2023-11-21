CREATE TABLE IF NOT EXISTS "apartment"
(
    "id"             serial PRIMARY KEY,
    "phone"           VARCHAR(50)         NOT NULL,
    "address"         VARCHAR(255)        NOT NULL,
    "category"        VARCHAR(255)        NOT NULL,
    "rating"          float(25)           NOT NULL
);

CREATE TABLE IF NOT EXISTS "bookCalendar"
(
    "id"                serial PRIMARY KEY,
    "apartment_id"      int    NOT NULL,
    "time"              date   NOT NULL
);
