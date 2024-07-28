deploy:
	- go build -o allo main.go
	- cp -f allo ~/cmd/allo