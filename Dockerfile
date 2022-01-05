FROM golang:latest
WORKDIR /var/www/app

COPY . ./
RUN go mod download
RUN go build -o /app ./cmd/main.go

EXPOSE 8080

CMD ["/app"]
