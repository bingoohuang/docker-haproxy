#!/bin/sh

# https://github.com/cynici/haproxy/blob/master/README.md

rm -f /var/run/rsyslogd.pid
rsyslogd

exec /usr/local/bin/docker-entrypoint.sh "$@"