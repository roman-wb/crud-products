FROM golang:alpine as builder
RUN apk --no-cache add git
WORKDIR /app
COPY . .
RUN GOOS=linux go build -ldflags "-s -w" -o bin/server cli/server/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/bin/server .
CMD ["./server"]
