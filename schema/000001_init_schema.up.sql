
CREATE TABLE IF NOT EXISTS home  (
id SERIAL NOT NULL  UNIQUE,
post_title VARCHAR(255)  NULL,
post_title_ru VARCHAR(255)  NULL,
post_img_path VARCHAR(255) NOT NULL DEFAULT '' ,
post_img_url VARCHAR(255) NULL,
post_body TEXT  NULL,
post_body_ru TEXT  NULL,
post_date TIMESTAMP  NULL DEFAULT (NOW()),
created_at TIMESTAMP  DEFAULT (NOW()),
updated_at TIMESTAMP NULL,
deleted_at TIMESTAMP NULL
);
CREATE TABLE IF NOT EXISTS admin(
    id SERIAL NOT NULL  UNIQUE,
    user_name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP  DEFAULT (NOW()),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL

);
CREATE TABLE IF NOT EXISTS contact (
id SERIAL NOT NULL  UNIQUE,
first_name VARCHAR(255) NOT NULL,
last_name VARCHAR(255)NOT  NULL ,
phone_number VARCHAR(255)NOT  NULL ,
type_service VARCHAR(255)NOT  NULL ,
text TEXT NOT NULL,
created_at TIMESTAMP  DEFAULT (NOW()),
updated_at TIMESTAMP NULL,
deleted_at TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS news (
id SERIAL NOT NULL  UNIQUE,
post_title VARCHAR(255)  NULL,
post_title_ru VARCHAR(255)  NULL,
post_img_path VARCHAR(255) NOT NULL DEFAULT '' ,
post_img_url VARCHAR(255) NULL,
post_body TEXT  NULL,
post_body_ru TEXT  NULL,
post_date TIMESTAMP  NULL DEFAULT (NOW()),
created_at TIMESTAMP  DEFAULT (NOW()),
updated_at TIMESTAMP NULL,
deleted_at TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS services (
id SERIAL NOT NULL  UNIQUE,
post_title VARCHAR(255)  NULL,
post_title_ru VARCHAR(255)  NULL,
post_img_path VARCHAR(255) NOT NULL DEFAULT '' ,
post_img_url VARCHAR(255) NULL,
post_body TEXT  NULL,
post_body_ru TEXT  NULL,
post_date TIMESTAMP  NULL DEFAULT (NOW()),
price  VARCHAR(255) NULL ,
created_at TIMESTAMP  DEFAULT (NOW()),
updated_at TIMESTAMP NULL,
deleted_at TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS tables (
id SERIAL NOT NULL  UNIQUE,
post_title VARCHAR(255)  NULL,
post_title_ru VARCHAR(255)  NULL,
post_img_path VARCHAR(255) NOT NULL DEFAULT '' ,
post_img_url VARCHAR(255) NULL,
post_body TEXT  NULL,
post_body_ru TEXT  NULL,
post_date TIMESTAMP  NULL DEFAULT (NOW()),
date VARCHAR(255) NULL,
format VARCHAR(255) NULL,
price  VARCHAR(255) NULL ,
duration  VARCHAR(255) NULL ,
created_at TIMESTAMP  DEFAULT (NOW()),
updated_at TIMESTAMP NULL,
deleted_at TIMESTAMP NULL
);


CREATE TABLE IF NOT EXISTS course (
id SERIAL NOT NULL  UNIQUE,
title VARCHAR(255)  NULL,
title_ru VARCHAR(255)  NULL,
price  VARCHAR(255) NULL ,
body VARCHAR(255) NULL,
body_ru TEXT  NULL,
duration  VARCHAR(255) NULL ,
term  VARCHAR(255) NULL,
format VARCHAR(255) NULL,
created_at TIMESTAMP  DEFAULT (NOW()),
updated_at TIMESTAMP NULL,
deleted_at TIMESTAMP NULL
);