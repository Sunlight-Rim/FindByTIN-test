## Build
FROM golang:latest

WORKDIR /findbytin

COPY . .

RUN go build -o server cmd/main.go

## Deploy
EXPOSE 8080

RUN chmod a+x server

CMD ["./server"]