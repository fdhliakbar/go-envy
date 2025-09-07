FROM golang:1.24.4

WORKDIR /app
COPY . .

RUN go build -o main ./examples/main.go

CMD ["./main"]
