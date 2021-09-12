FROM golang:1.16.7
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN go build -o main ./cmd/bank-service/
CMD ["/app/main"]