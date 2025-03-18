FROM golang:latest

WORKDIR /usr/src/app

COPY . .
RUN go mod download

RUN go build -o server ./cmd/apollo-server/main.go

EXPOSE 5000

CMD ["./server"]
