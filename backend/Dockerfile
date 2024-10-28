FROM golang:1.22

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ .
COPY .env .

RUN go build -o main ./cmd/main.go

EXPOSE ${BACKEND_PORT}

CMD ["./main"]
