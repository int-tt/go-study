#!/bin/sh
go build

./ex08 &

nc localhost 8000 < msg.text

#killall ex08
