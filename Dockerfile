FROM golang:1.21

WORKDIR /usr/src/app

COPY . .

RUN go build -v -o /usr/local/bin/app ./net_pro/http_ex.go

CMD ["app"]

#docker run -it --rm -p 3000:3000 --name test gopro:1