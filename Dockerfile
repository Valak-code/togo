FROM golang:latest

RUN mkdir -p /app
WORKDIR /app
COPY . .
RUN mkdir -p bin
RUN go build -o ./bin/main main.go

EXPOSE 5050

CMD ["/app/bin/main"] 