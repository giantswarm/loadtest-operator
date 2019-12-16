FROM alpine:3.10

RUN apk add --no-cache ca-certificates

ADD ./loadtest-operator /loadtest-operator

ENTRYPOINT ["/loadtest-operator"]
