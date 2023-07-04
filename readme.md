# learning go

steps to initialize go:
1. go mod init <package_name>
2. create <filename>.go
3. run: go run .

to update the location of packages, run edit command, run: go mod edit -replace example.com/greetings=../greetings
to update go.mod, run: go mod tidy

Tutorials:
1. https://go.dev/doc/tutorial/create-module
2. https://go.dev/doc/tutorial/call-module-code
3. https://go.dev/doc/tutorial/handle-errors
4. https://go.dev/doc/tutorial/random-greeting 
5. https://go.dev/doc/tutorial/greetings-multiple-people 