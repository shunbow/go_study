FROM golang:1.19.2-alpine3.16

ENV GOPATH /go
ENV GO111MODULE on
ENV ROOT=/go/src/app
RUN apk update && \
    apk --no-cache add git

RUN mkdir ${ROOT}
WORKDIR ${ROOT}

ADD ../.. ${ROOT}

RUN go mod tidy && \
    go install github.com/cosmtrek/air@v1.40.4 && \
    go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go get -u github.com/go-sql-driver/mysql
CMD ["air", "-c", ".air.toml"]
