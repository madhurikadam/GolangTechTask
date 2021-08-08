FROM alpine
ADD . /
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY .builds/golangtechtask  ./
ENTRYPOINT ["./golangtechtask"]
