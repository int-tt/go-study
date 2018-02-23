#!/bin/bash

go build -o bin/clock clock/main.go
go build -o bin/clockwall clockwall/main.go

TZ=US/Eastern bin/clock -port 8010 &
TZ=Asia/Tokyo bin/clock -port 8020 &
TZ=Europe/London bin/clock -port 8030 &

trap 'killall clock' 2
bin/clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030
