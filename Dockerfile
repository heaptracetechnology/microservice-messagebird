FROM golang

RUN go get github.com/gorilla/mux

RUN go get github.com/messagebird/go-rest-api

RUN go get github.com/messagebird/go-rest-api/sms

WORKDIR /go/src/github.com/oms-services/messagebird

ADD . /go/src/github.com/oms-services/messagebird

RUN go install github.com/oms-services/messagebird

ENTRYPOINT messagebird

EXPOSE 3000