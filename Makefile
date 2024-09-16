deploy-main:
	- go build -o allo main.go
	- cp -f allo-main ~/cmd/allo
deploy-cli:
	- go build -o allo cmd/cli/main.go
	- cp -f allo ~/cmd/allo

debug-bubbletea:
	- tail -f debug.log