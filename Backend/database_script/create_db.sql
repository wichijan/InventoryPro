DROP DATABASE IF EXISTS InventoryProDB;

CREATE DATABASE IF NOT EXISTS InventoryProDB;

USE InventoryProDB;

CREATE TABLE warehouse_types(
    id BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)) PRIMARY KEY,
    TYPE VARCHAR(100)
);

CREATE TABLE warehouses(
    id BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)) PRIMARY KEY,
    name VARCHAR(100),
    type_id BINARY(16),
    FOREIGN KEY (type_id) REFERENCES warehouse_types(id)
);

CREATE TABLE rooms(
    id BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)) PRIMARY KEY,
    warehouse_id BINARY(16),
    FOREIGN KEY (warehouse_id) REFERENCES warehouses(id)
);

CREATE TABLE shelves(
    id BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)) PRIMARY KEY,
    room_id BINARY(16),
    FOREIGN KEY (room_id) REFERENCES rooms(id)
);

CREATE TABLE item_status(
    id BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)) PRIMARY KEY,
    status_name VARCHAR(100)
);

CREATE TABLE items(
    id BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)) PRIMARY KEY,
    name VARCHAR(100),
    description TEXT,
    class_one BOOLEAN,
    class_two BOOLEAN,
    class_three BOOLEAN,
    class_four BOOLEAN,
    damaged BOOLEAN,
    damaged_description TEXT,
    quantity INT,
    shelf_id BINARY(16),
    status_id BINARY(16),
    FOREIGN KEY (status_id) REFERENCES item_status(id),
    FOREIGN KEY (shelf_id) REFERENCES shelves(id)
);

create table subjects(
    id BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)) PRIMARY KEY,
    name VARCHAR(100),
    description TEXT
);

create table item_subjects(
    id BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)) PRIMARY KEY,
    item_id BINARY(16),
    subject_id BINARY(16),
    FOREIGN KEY (item_id) REFERENCES items(id),
    FOREIGN KEY (subject_id) REFERENCES subjects(id)
);

CREATE TABLE item_pictures(
    id BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)) PRIMARY KEY,
    picture BLOB item_id BINARY(16),
    FOREIGN KEY (item_id) REFERENCES items(id)
);

CREATE TABLE keywords_for_items(
    id BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)) PRIMARY KEY,
    keyword VARCHAR(100),
    item_id BINARY(16),
    FOREIGN KEY (item_id) REFERENCES items(id)
);

create table user_types (
    id BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)) PRIMARY KEY,
    type_name VARCHAR(100)
);

create table roles(
    id BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)) PRIMARY KEY,
    roles_name VARCHAR(100)
    /* Here come the roles (read table...) so that users can have multiple grants  */
);

CREATE TABLE users(
    id BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)) PRIMARY KEY,
    username VARCHAR(100),
    password VARCHAR(100),
    email VARCHAR(100),
    role VARCHAR(100),
    user_type_id BINARY(16),
    FOREIGN KEY (user_type_id) REFERENCES user_types(id)
);

create table user_roles(
    id BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)) PRIMARY KEY,
    user_id BINARY(16),
    role_id BINARY(16),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (role_id) REFERENCES roles(id)
);