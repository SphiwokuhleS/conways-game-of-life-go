#!/bin/bash

sqlite3 api/conways.db "CREATE TABLE worlds (
id         INTEGER     NOT NULL PRIMARY KEY AUTOINCREMENT,
created_at DATETIME,
updated_at DATETIME,
deleted_at DATETIME,
name       TEXT        UNIQUE,
grid       TEXT,
epoch      INTEGER);"