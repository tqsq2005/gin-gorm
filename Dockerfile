FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/gin-gorm
COPY . $GOPATH/gin-gorm
RUN go build .

EXPOSE 8088
ENTRYPOINT ["./gin-gorm"]