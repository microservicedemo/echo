FROM alpine:edge

ADD bundles/echo /usr/bin/echo

EXPOSE 9090
CMD /usr/bin/echo
