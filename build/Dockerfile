FROM golang:1.22-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /out/app ./cmd/proxy

FROM alpine:latest

COPY --from=builder /out/app /app

ENTRYPOINT ["/app"]