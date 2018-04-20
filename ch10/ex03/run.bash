#!/bin/bash

go build -o fetch ../../ch01/fetch
./fetch http://gopl.io/ch1/helloworld?go-get=1