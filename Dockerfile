FROM golang:1.23-alpine

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o receipt-points-service .

CMD ["./receipt-points-service"]