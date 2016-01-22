FROM golang:1.5.3
MAINTAINER xjimmyshcn@gmail.com

RUN wget http://mirrors.163.com/.help/sources.list.jessie -O /etc/apt/sources.list
RUN apt-get update
RUN apt-get install -y pkg-config patch libpcre3-dev

RUN mkdir -p /root/gopath

ENV GOROOT_BOOTSTRAP /usr/local/go
ENV GOPATH /root/gopath

#GOROOT and GOPATH
VOLUME /go
VOLUME /root/gopath

WORKDIR /go/src/
CMD ["./all.bash"]
