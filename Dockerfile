FROM golang

ADD . /go/src/TapTalk-BE
WORKDIR /go/src/TapTalk-BE

RUN go get -d -v .
RUN go install -v .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]