FROM golang:1.22-alpine as builder

WORKDIR /app

COPY ../go.mod /app/
COPY ../go.sum /app/
COPY . /app/

RUN go mod download

EXPOSE 5004

CMD ["go", "run", "/app/payment/main.go", "."]