FROM golang:latest

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io

WORKDIR /build
COPY . /build
RUN go build .

EXPOSE 4567
ENTRYPOINT ["./jinkela"]