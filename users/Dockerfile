FROM golang

RUN mkdir -p /go/src/github.com/ibadi-id/projectmain

ADD . /go/src/github.com/ibadi-id/projectmain

WORKDIR /go/src/github.com/ibadi-id/projectmain/users

RUN go mod tidy
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run /go/src/github.com/ibadi-id/projectmain/users/cmd  -watch github.com/ibadi-id/projectmain/users
