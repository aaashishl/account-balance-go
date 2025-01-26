#!/bin/bash

read -p "Enter account address:" address

go run ./main.go "$address"