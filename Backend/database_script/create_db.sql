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
    name VARCHAR(100),
    warehouse_id VARCHAR(36),
    FOREIGN KEY (warehouse_id) REFERENCES warehouses(id)
);

CREATE TABLE shelve_types(
    id VARCHAR(36) PRIMARY KEY,
    type_name VARCHAR(100)
);

CREATE TABLE shelves(
    id VARCHAR(36) PRIMARY KEY,
    shelve_type_id VARCHAR(36),
    room_id VARCHAR(36),
    FOREIGN KEY (room_id) REFERENCES rooms(id),
    FOREIGN KEY (shelve_type_id) REFERENCES shelve_types(id)
);

CREATE TABLE item_status(
    id VARCHAR(36) PRIMARY KEY,
    status_name VARCHAR(100)
);

CREATE TABLE items(
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(100),
    description TEXT,
    class_one BOOLEAN,
    class_two BOOLEAN,
    class_three BOOLEAN,
    class_four BOOLEAN,
    damaged BOOLEAN,
    damaged_description TEXT,
    quantity INT,
    status_id VARCHAR(36),
    FOREIGN KEY (status_id) REFERENCES item_status(id)
);

Create table items_in_shelve(
    id VARCHAR(36) PRIMARY KEY,
    item_id VARCHAR(36) UNIQUE,
    shelve_id VARCHAR(36),
    FOREIGN KEY (item_id) REFERENCES items(id),
    FOREIGN KEY (shelve_id) REFERENCES shelves(id)
);

create table subjects(
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(100),
    description TEXT
);

create table item_subjects(
    id VARCHAR(36) PRIMARY KEY,
    item_id VARCHAR(36),
    subject_id VARCHAR(36),
    FOREIGN KEY (item_id) REFERENCES items(id),
    FOREIGN KEY (subject_id) REFERENCES subjects(id)
);

CREATE TABLE item_pictures(
    id VARCHAR(36) PRIMARY KEY,
    picture BLOB,
    item_id VARCHAR(36),
    FOREIGN KEY (item_id) REFERENCES items(id)
);

CREATE TABLE keywords_for_items(
    id VARCHAR(36) PRIMARY KEY,
    keyword VARCHAR(100),
    item_id VARCHAR(36),
    FOREIGN KEY (item_id) REFERENCES items(id)
);

create table user_types (
    id VARCHAR(36) PRIMARY KEY,
    type_name VARCHAR(100)
);

create table roles(
    id VARCHAR(36) PRIMARY KEY,
    role_name VARCHAR(100)
    /* Here come the roles (read table...) so that users can have multiple grants  */
);

CREATE TABLE users(
    id VARCHAR(36) PRIMARY KEY,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    username VARCHAR(100),
    email VARCHAR(100),
    password VARCHAR(100),
    job_title VARCHAR(100),
    phone_number VARCHAR(100),
    user_type_id VARCHAR(36),
    FOREIGN KEY (user_type_id) REFERENCES user_types(id)
);

create table user_roles(
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36),
    role_id VARCHAR(36),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (role_id) REFERENCES roles(id)
);