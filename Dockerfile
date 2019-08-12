FROM golang:1.12-alpine3.10 as builder
LABEL maintainer="jxlwqq <jxlwqq@gmail.com>"
WORKDIR /app
ENV GOPROXY https://goproxy.io
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.10
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]