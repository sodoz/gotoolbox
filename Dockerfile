FROM busybox:ubuntu-14.04

ADD ./gotoolbox /usr/bin/

EXPOSE 8080

ENTRYPOINT ["gotoolbox"]
