FROM golang:1.22

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ .

RUN mkdir -p /app/supabase/migrations

RUN go build -o main ./cmd/main.go

EXPOSE 8080

CMD ["./main"]
