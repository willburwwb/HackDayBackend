FROM golang:1.20.2

ENV GOPROXY=http://goproxy.cn,direct
ENV GO111MODULE=on 

WORKDIR /app

COPY . .

RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o app .

EXPOSE 3333

CMD ["./app"]
