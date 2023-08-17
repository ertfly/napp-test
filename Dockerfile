FROM golang

ADD . /app

WORKDIR /app

RUN go mod download
RUN go install napptest
ENTRYPOINT /go/bin/napptest
EXPOSE 8000