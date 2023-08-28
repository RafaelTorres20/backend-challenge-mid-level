SHELL:=/bin/bash

# Load environment variables from .env
include .env

run:
	go run main.go server
