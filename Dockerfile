FROM busybox:ubuntu-14.04

ADD ./gotoolbox /usr/bin/
ADD ./views /usr/bin/views
ADD ./public /usr/bin/public

EXPOSE 8080

ENTRYPOINT ["gotoolbox"]
