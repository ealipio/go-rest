## REST API with golang
> using gorilla

verify that no one is using 8080 port, if that is the case then run:
```bash
go run -race main.go
```
if port 8080 is in use for other application just change the port number in **main.go**.