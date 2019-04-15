FROM golang

RUN go get github.com/gorilla/mux

RUN go get github.com/messagebird/go-rest-api

RUN go get github.com/messagebird/go-rest-api/sms

WORKDIR /go/src/github.com/heaptracetechnology/microservice-messagebird

ADD . /go/src/github.com/heaptracetechnology/microservice-messagebird

RUN go install github.com/heaptracetechnology/microservice-messagebird

ENTRYPOINT microservice-messagebird

EXPOSE 3000