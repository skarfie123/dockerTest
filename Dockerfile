#1st build
FROM golang:1.10
WORKDIR /go/src/github.com/skarfie123/dockertest

COPY main ./main
COPY msg ./msg
COPY vendor ./vendor

RUN CGO_ENABLED=0 GOOS=linux go install ./main

#2nd Stage

FROM alpine:latest
WORKDIR /app/

COPY --from=0 /go/bin/main ./binary

CMD ./binary