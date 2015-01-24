FROM busybox:ubuntu-14.04

ADD ./gotoolbox /usr/bin/
ADD ./views /views
ADD ./public /public

EXPOSE 8080

ENTRYPOINT ["gotoolbox"]
