FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o ./tmp/main ./cmd/main.go 

EXPOSE 8080

CMD ["./tmp/main"]
