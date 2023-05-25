# Go project template

This is for my personal use but feel free to use it. It's for a command line tool. I won't be using it for pure libraries

To test, put your test files in ./test and run `go test -v ./test`

To run the command line, run `go run cmd/main.go`

To build, run `go build -C cmd -o go-project.exe` or change directory to cmd and run `go build`

You can create scripts for these commands to simplify it, but I wanted to keep this platform neutral, and I would write powershell scripts or Docker related files