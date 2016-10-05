FROM alpine:3.4

ADD *.go /garto-go/

RUN apk update \
  && apk add bash git go \
  && export GOPATH=/gopath \
  && REPO_PATH="github.com/Financial-Times/garto-go" \
  && mkdir -p $GOPATH/src/${REPO_PATH} \
  && mv /garto-go/* $GOPATH/src/${REPO_PATH} \
  && cd $GOPATH/src/${REPO_PATH} \
  && go get \
  && go build \
  && mv garto-go /app \
  && apk del go git bzr \
  && rm -rf $GOPATH /var/cache/apk/* /garto-go \
  && mv /app /garto-go

CMD [ "/garto-go" ]
