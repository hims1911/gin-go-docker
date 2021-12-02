# gin-go-docker
Example Project on Gin Go framework with docker container

### To Run without docker instance

```
    go get package_name
    go run src/main.go
```

### TO Run a Docker Container

```
    docker build . -t go-gin-container
    docker run -i -t -p 8080:8080 go-gin-container
```