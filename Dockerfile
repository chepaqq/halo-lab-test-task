FROM golang:alpine as build

WORKDIR /build
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o test-task /build/cmd/server/main.go

FROM scratch
COPY --from=build /build/test-task /test-task
EXPOSE 8000
CMD ["/test-task"]

