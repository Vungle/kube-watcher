FROM alpine:3.3
ADD bin/kube-monitor /usr/bin/kube-monitor
RUN chmod +x /usr/bin/kube-monitor \
  && apk add --update -t deps ca-certificates \
  && apk del --purge deps \
  && rm /var/cache/apk/*
CMD "kube-monitor"
