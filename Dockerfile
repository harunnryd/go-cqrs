FROM golang:latest

MAINTAINER harunwols@gmail.com

RUN go get -u github.com/kardianos/govendor

RUN go get github.com/onsi/ginkgo/ginkgo

RUN go get github.com/onsi/gomega

RUN go get github.com/Shopify/sarama

RUN go get github.com/bsm/sarama-cluster

RUN go get github.com/jinzhu/copier

RUN go get github.com/jinzhu/gorm

RUN go get github.com/jinzhu/gorm/dialects/postgres

RUN go get github.com/labstack/echo

RUN go get github.com/satori/go.uuid

RUN mkdir -p /go/src/quizes

COPY .  /go/src/quizes/

WORKDIR /go/src/quizes

RUN govendor sync

RUN go get -d -v ./...

RUN go install -v ./...

RUN go build -o quizes .

EXPOSE 8080

CMD ["quizes"]

