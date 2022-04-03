FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /var/www
COPY . .

RUN go build -o app ./cmd/main.go

EXPOSE 3200
ENTRYPOINT [ "/var/www/app" ]
