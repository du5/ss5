FROM golang:alpine

RUN apk add git \
    && go get github.com/du5/ss5 

FROM alpine

COPY --from=0 /go/bin/ss5 /usr/bin/ss5

ENTRYPOINT [ "ss5" ]