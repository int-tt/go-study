go build ../../ch01/fetch/
go build -o findlinks

./fetch http://golang.org | ./findlinks
