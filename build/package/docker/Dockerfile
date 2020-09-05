FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN apk update && apk upgrade && \
  apk add --no-cache bash git openssh
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/gocicd/index.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app . 
CMD ["./server"]