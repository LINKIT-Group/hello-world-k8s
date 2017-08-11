FROM golang:onbuild

WORKDIR /APP

ADD main.go .
RUN go build -o main .

CMD ["/APP/main"]
