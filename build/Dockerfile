FROM golang:1.22-alpine as builder

ARG CGO_ENABLED=0
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build cmd/proxy -o /out/app

# multi stage build
FROM alpine:latest
COPY --from=builder /out/app /app
ENTRYPOINT ["/app"]