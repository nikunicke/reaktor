FROM golang:1.15.5-alpine3.12
WORKDIR /usr/src/app
COPY . .
RUN go build ./cmd/api
CMD ["./api"]
