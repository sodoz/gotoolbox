FROM busybox:ubuntu-14.04

ADD ./gotoolbox /usr/bin/
ADD ./views /views
ADD ./public /public

RUN apt-get update -y && apt-get install —no-install-recommends -y -q  ca-certificates

EXPOSE 8080

ENTRYPOINT ["gotoolbox"]
