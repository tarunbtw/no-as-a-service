# build the binary
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o naas .

# run the binary
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/naas .

EXPOSE 8080

CMD ["./naas"]
