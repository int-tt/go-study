#!/bin/sh

go build -o mandelbrot ../../ch03/ex05 
go build -o encoder

./mandelbrot | ./encoder -f jpg > output.jpg
./mandelbrot | ./encoder -f png > output.png
./mandelbrot | ./encoder -f gif > output.gif