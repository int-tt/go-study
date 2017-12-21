#!/bin/bash

go run main.go > image.png
open `pwd`/image.png
