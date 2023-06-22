learning go

steps to initialize go:
1. go mod init <package_name>
2. create <filename>.go
3. run: go run .

to update the location of packages, run edit command, run: go mod edit -replace example.com/greetings=../greetings
to update go.mod, run: go mod tidy