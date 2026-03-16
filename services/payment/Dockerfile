FROM golang:1.25-alpine

WORKDIR /app

# copy root module files
COPY go.mod go.sum ./

RUN go mod download

# copy the entire project
COPY . .

# build the payment service
RUN go build -o payment ./services/payment

CMD ["./payment"]