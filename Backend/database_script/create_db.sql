DROP DATABASE IF EXISTS InventoryProDB;

CREATE DATABASE IF NOT EXISTS InventoryProDB;

USE InventoryProDB;

CREATE TABLE warehouses(
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(100),
    description TEXT
);

CREATE TABLE rooms(
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    warehouse_id VARCHAR(36),
    FOREIGN KEY (warehouse_id) REFERENCES warehouses(id)
);

CREATE TABLE shelves(
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    room_id VARCHAR(36),
    FOREIGN KEY (room_id) REFERENCES rooms(id)
);

/* ITEMS */
CREATE TABLE items(
    id VARCHAR(36) PRIMARY KEY,
    item_types ENUM('book', 'single_object', 'set_of_objects') NOT NULL,
    name VARCHAR(100),
    description TEXT,
    regular_shelf_id VARCHAR(36),
    class_one BOOLEAN,
    class_two BOOLEAN,
    class_three BOOLEAN,
    class_four BOOLEAN,
    damaged BOOLEAN,
    damaged_description TEXT,
    picture TEXT,
    hint_text TEXT,
    FOREIGN KEY (regular_shelf_id) REFERENCES shelves(id)
);

CREATE TABLE books(
    item_id VARCHAR(36) PRIMARY KEY,
    ISBN VARCHAR(50) UNIQUE NOT NULL,
    author VARCHAR(100),
    publisher VARCHAR(100),
    edition VARCHAR(100),
    FOREIGN KEY (item_id) REFERENCES items(id)
);

Create table single_object (
    item_id VARCHAR(36),
    PRIMARY KEY (item_id),
    FOREIGN KEY (item_id) REFERENCES items(id)
);

Create table sets_of_objects (
    item_id VARCHAR(36) PRIMARY KEY,
    total_objects INT,
    useful_objects INT,
    broken_objects INT,
    lost_objects INT,
    FOREIGN KEY (item_id) REFERENCES items(id)
);

create table subjects(
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(100) UNIQUE,
    description TEXT
);

create table item_subjects(
    item_id VARCHAR(36),
    subject_id VARCHAR(36),
    PRIMARY KEY (subject_id, item_id),
    FOREIGN KEY (item_id) REFERENCES items(id),
    FOREIGN KEY (subject_id) REFERENCES subjects(id)
);

CREATE TABLE keywords(
    id VARCHAR(36) PRIMARY KEY,
    keyword VARCHAR(100)
);

CREATE TABLE keywords_for_items(
    keyword_id VARCHAR(36),
    item_id VARCHAR(36),
    PRIMARY KEY (keyword_id, item_id),
    FOREIGN KEY (item_id) REFERENCES items(id),
    FOREIGN KEY (keyword_id) REFERENCES keywords(id)
);

/*  */
Create table items_in_shelf(
    item_id VARCHAR(36) UNIQUE,
    shelf_id VARCHAR(36),
    quantity INT,
    PRIMARY KEY (shelf_id, item_id),
    FOREIGN KEY (item_id) REFERENCES items(id),
    FOREIGN KEY (shelf_id) REFERENCES shelves(id)
);

create table user_types (
    id VARCHAR(36) PRIMARY KEY,
    type_name VARCHAR(100)
);

create table roles(
    id VARCHAR(36) PRIMARY KEY,
    role_name VARCHAR(100) UNIQUE
    /* Here come the roles (read table...) so that users can have multiple grants  */
);

CREATE TABLE users(
    id VARCHAR(36) PRIMARY KEY,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    username VARCHAR(100) UNIQUE,
    email VARCHAR(100),
    password VARCHAR(100),
    job_title VARCHAR(100),
    phone_number VARCHAR(100),
    user_type_id VARCHAR(36),
    profile_picture TEXT,
    registration_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    registration_accepted BOOLEAN DEFAULT FALSE,
    /* So admin can accept */
    is_active BOOLEAN,
    /* Point - immer wenn er das schnell regal leer, bekommt er punkte */
    FOREIGN KEY (user_type_id) REFERENCES user_types(id)
);

create table registration_codes (
    user_id VARCHAR(36),
    code VARCHAR(20) UNIQUE,
    PRIMARY KEY (user_id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

create table registration_requests(
    user_id VARCHAR(36),
    request_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

create table user_roles(
    user_id VARCHAR(36),
    role_id VARCHAR(36),
    PRIMARY KEY (user_id, role_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (role_id) REFERENCES roles(id)
);

CREATE TABLE quick_shelves (
    quick_shelf_id VARCHAR(36) PRIMARY KEY,
    room_id VARCHAR(36),
    /* So we know where shelf is */
    FOREIGN KEY (room_id) REFERENCES rooms(id)
);

Create table item_quick_shelf(
    item_id VARCHAR(36),
    user_id VARCHAR(36),
    quick_shelf_id VARCHAR(36),
    quantity INT,
    PRIMARY KEY (item_id, user_id, quick_shelf_id),
    Foreign KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (item_id) REFERENCES items(id),
    FOREIGN KEY (quick_shelf_id) REFERENCES quick_shelves(quick_shelf_id)
);

CREATE TABLE points (
    user_id VARCHAR(36),
    points INT DEFAULT 0,
    PRIMARY KEY (user_id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

/* Immer aktueller Stand */
create table user_items(
    user_id VARCHAR(36) NOT NULL,
    item_id VARCHAR(36) NOT NULL,
    quantity INT,
    transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, item_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (item_id) REFERENCES items(id)
);

/* Planning Reservations */
create table reservations (
    reservation_id VARCHAR(36) PRIMARY KEY,
    item_id VARCHAR(36),
    user_id VARCHAR(100),
    username VARCHAR(36),
    quantity INT,
    reservation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    time_from TIMESTAMP NOT NULL,
    time_to TIMESTAMP NOT NULL,
    is_cancelled BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (item_id) REFERENCES items(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

/* Transaction for history */
CREATE TABLE transactions (
    transaction_id VARCHAR(36) PRIMARY KEY,
    item_id VARCHAR(36) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    transaction_type ENUM(
        'borrow',
        'return',
        'place_in_quick_shelf',
        'transfer_request',
        'transfer_accepted',
        'reserve',
        'cancel_reservation',
        'report_lost',
        'report_damaged'
    ) NOT NULL,
    target_user_id VARCHAR(36),
    origin_user_id VARCHAR(36),
    transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    note TEXT
    /* For damage report */
);

/* Moving item from  */
Create table transfer_requests (
    transfer_request_id VARCHAR(36) PRIMARY KEY,
    item_id VARCHAR(36) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    target_user_id VARCHAR(36) NOT NULL,
    request_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_accepted BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (item_id) REFERENCES items(id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (target_user_id) REFERENCES users(id)
);