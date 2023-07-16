FROM golang:latest

WORKDIR /go/src/github.com/halo-lab-test-task
COPY . .

RUN go mod download
RUN go build -o test-task ./cmd/server/main.go
EXPOSE 8000
CMD ["./test-task"]
