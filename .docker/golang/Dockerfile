FROM golang

ADD . /go/src/napptest

RUN go get
    
RUN go install napptest

ENTRYPOINT /go/bin/napptest

EXPOSE 80