FROM golang:1.21.5-alpine AS builder

WORKDIR /image-messer

COPY . .

RUN go mod download
RUN go build -o image-messer

EXPOSE 8080

CMD ["./image-messer"]