FROM golang:latest
WORKDIR /usr/src/app
COPY ../test2_serij.go .
COPY go.mod .
RUN go mod tidy
RUN go build -o test2_serij .
CMD ["./test2_serij"]
