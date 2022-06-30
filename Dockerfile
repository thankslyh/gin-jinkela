FROM golang:latest

WORKDIR /build
COPY . /build
RUN go build .

EXPOSE 4567
ENTRYPOINT ["./jinkela"]