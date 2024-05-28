#!/bin/bash
source .env

export DB_HOST=$DB_HOST
export DB_USER=$DB_USER
export DB_PASSWORD=$DB_PASSWORD

rm ../uploads/*

mariadb -u $DB_USER --password=$DB_PASSWORD -h $DB_HOST < ./database_script/create_db.sql
mariadb -u $DB_USER --password=$DB_PASSWORD -h $DB_HOST < ./database_script/insert_db.sql