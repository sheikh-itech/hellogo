FROM golang:1.4
EXPOSE 8888
WORKDIR /go/src/app
COPY . /go/src/app
RUN ["/go/src/app/bin/build"]
CMD /go/src/app/bin/run
