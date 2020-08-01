FROM golang:latest as builder

ENV APP_HOME /go/src/its-rankings-scrapper

RUN mkdir -p $APP_HOME

WORKDIR $APP_HOME
COPY . .
COPY ./common/config/yaml/config.dist.yml ./common/config/yaml/config.yml

RUN go mod download
RUN go mod verify
RUN go get -u "github.com/qor/admin"
RUN go get -u -f github.com/qor/bindatafs/...
RUN CGO_ENABLED=0 go build -tags 'bindatafs' -o app

FROM alpine:latest

RUN apk --no-cache add ca-certificates

ENV APP_HOME /go/src/its-rankings-scrapper

RUN mkdir -p $APP_HOME

WORKDIR $APP_HOME

COPY --from=builder $APP_HOME/app $APP_HOME
COPY --from=builder $APP_HOME/common/config/yaml/config.dist.yml $APP_HOME/config.yml
RUN mkdir -p $APP_HOME/logs

EXPOSE 9000
CMD ["./app", "admin"]