FROM golang:1.23-alpine
RUN go install github.com/air-verse/air@latest
WORKDIR /app
ENTRYPOINT ["air"]