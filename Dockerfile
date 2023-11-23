# Dockerfile
FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o fetch .

CMD ["./fetch"]
