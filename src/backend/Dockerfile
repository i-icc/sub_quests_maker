FROM golang:latest

RUN apt-get update && apt-get install -y vim
ENV HOME="/backend"
ENV GOPATH=/usr
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN mkdir -p ${HOME}
WORKDIR ${HOME}
COPY go.mod .
RUN go mod download
RUN go mod tidy
COPY . ${HOME}

CMD [ "go", "run", "server.go" ]