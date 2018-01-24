FROM golang:1.9
ADD . /go/src/go_test_travis
WORKDIR /go/src/go_test_travis
RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

FROM alpine:3.7
WORKDIR /root/
COPY --from=0 /go/src/go_test_travis/app .
CMD ["./app"]