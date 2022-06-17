FROM golang:1.16.6-alpine

WORKDIR /code

ADD . .
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GO111MODULE=on

RUN cd /code/discovery/client && go build -o /app/client

CMD /app/client
