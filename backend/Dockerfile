FROM golang:1.21

WORKDIR /app

COPY go.mod .
RUN go mod download || true

COPY . .

RUN go build -o main ./cmd/main.go

EXPOSE 8080

CMD ["./main"]