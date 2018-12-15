FROM golang:alpine

ENV DATA_PATH /data

WORKDIR /go/src/github.com/bobrnor/highloadcup2018
COPY . .

RUN go install -v ./...

EXPOSE 80

CMD ["sh", "run.sh"]