#!/bin/bash

# Create Tables
psql findmyrecipe -f './server/database/queries/sql/createTables.sql'
psql findmyrecipe -c "select createTables();"

# Define Drop Function
psql findmyrecipe -f './server/database/queries/sql/dropTables.sql'


psql findmyrecipe -c "alter sequence users_id_seq restart with 100;"
