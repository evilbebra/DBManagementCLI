FROM golang:1.20

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Собираем Go приложение
RUN go build -o main ./cmd

ENV PORT=8080
CMD ["./main"]