FROM golang:1.23.0-alpine3.20
LABEL authors="zen"
RUN mkdir /App
WORKDIR /app
COPY . .
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors4.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN apk add ffmpeg build-base mediainfo
RUN go mod tidy
RUN go mod vendor
RUN go build -o /usr/local/bin/AVTH main.go
ENTRYPOINT ["/usr/local/bin/AVTH"]

# docker build -t any_videos_to_h265:2 .