FROM golang:1.15.5-alpine3.12
WORKDIR /usr/src/app
COPY . .
RUN go mod init github.com/nikunicke/reaktor/warehouse_api
RUN go build ./cmd/warehouse_api
CMD ["./warehouse_api"]
