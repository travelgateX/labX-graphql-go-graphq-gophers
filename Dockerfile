FROM golang:latest
LABEL Description="golang" Vendor="LabX XML Travelgate" Version="1.0.0"

COPY . /go/src/labX-graphql-go-graphq-gophers/
COPY ./resource /go/resource

RUN groupadd -g 666 go \
    && useradd -M -d /go -g go -u 666 -s /bin/bash go \
    && chown -R go:go /go

USER go
RUN cd /go/src/labX-graphql-go-graphq-gophers/cmd/task1 \
    && go get \
    && go build -v

EXPOSE 8080
CMD ["bash", "-c", "/go/bin/task1"]