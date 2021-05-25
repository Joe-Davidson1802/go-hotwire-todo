FROM golang:latest


ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

RUN go get -u github.com/a-h/templ/cmd

WORKDIR /src
COPY go.mod /src/.
COPY go.sum /src/.

RUN go mod download

COPY . /src/.

RUN go mod verify
RUN go mod vendor

RUN $(go env GOPATH)/bin/cmd generate
RUN go build -o /src/todo
EXPOSE 80
CMD ["/src/todo"]
