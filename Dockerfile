FROM golang:1.21.1-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build specifically for Linux AMD64 because it's what render.com has running
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./main.go

RUN chmod +x main

EXPOSE 1991

CMD ["./main"]
