FROM golang

COPY . /app
WORKDIR /app

RUN go get
RUN go install napptest

ENTRYPOINT ["/go/bin/napptest"]

EXPOSE 8000