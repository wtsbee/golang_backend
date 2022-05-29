FROM golang:1.14

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

RUN go mod init example
RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql

EXPOSE 8080

# CMD ["go", "run", "main.go"]