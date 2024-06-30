FROM golang:1.22-bookworm as builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go mod download

RUN go build -o main .

FROM busybox:1.36-glibc

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]