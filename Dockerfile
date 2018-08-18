FROM golang:alpine
RUN apk add --no-cache git

WORKDIR /go/src/github.com/lovelaced/glitzz
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["glitzz run /data/config.json"]
