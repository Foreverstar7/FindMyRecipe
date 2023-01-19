#!/bin/bash

# User Sign Up
psql findmyrecipe -c "select userSignup('testuser', 'testpass');"

# User Login
psql findmyrecipe -c "select userLogin('testuser');"

