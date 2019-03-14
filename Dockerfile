## build binary
FROM golang:1.11-alpine
WORKDIR /build
COPY main.go .
RUN go build -o /build/go-app

## copy binary to app container
FROM alpine:latest
WORKDIR /app
COPY --from=0 /build/go-app .
EXPOSE 3000
CMD ./go-app
