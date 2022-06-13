FROM haproxy:1.8.30-alpine

RUN cd /etc/apk/keys \
    && wget http://alpine.adiscon.com/rsyslog@lists.adiscon.com-5a55e598.rsa.pub \
    && echo 'http://alpine.adiscon.com/3.7/stable' >> /etc/apk/repositories \
    && apk update && apk add --no-cache bash ca-certificates rsyslog \
    && touch /var/log/haproxy.log \
    && ln -sf /dev/stdout /var/log/haproxy.log

COPY haproxy.cfg /etc/haproxy.cfg
COPY rsyslog.conf /etc/rsyslog.conf

ADD docker-entrypoint-pre.sh /docker-entrypoint-pre.sh
ENTRYPOINT ["/docker-entrypoint-pre.sh"]


# no USER for backwards compatibility (to try to avoid breaking existing users)
CMD ["haproxy", "-f", "/etc/haproxy.cfg"]