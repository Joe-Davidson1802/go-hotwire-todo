FROM golang:latest


ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

RUN go get -u github.com/a-h/templ/cmd

WORKDIR /src
COPY . /src

RUN go mod download
RUN go mod verify
RUN go mod vendor
RUN $(go env GOPATH)/bin/cmd generate
RUN go build -o /src/todo
EXPOSE 80
CMD ["/src/todo"]
