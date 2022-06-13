# docker-haproxy-nginx

## Usage

- start: `make up`
- clean: `make clean`

## Test

```sh
bingoobjca@bogon docker-haproxy % for i in $(seq 1 20); do curl http://127.0.0.1:8080; done
From nginx2
From nginx3
From nginx1
From nginx2
From nginx3
From nginx1
From nginx2
From nginx3
From nginx1
From nginx2
From nginx3
From nginx1
From nginx2
From nginx3
From nginx1
From nginx2
From nginx3
From nginx1
From nginx2
From nginx3
```

## HAProxy versions

[versions](http://www.haproxy.org/)

Branch  | Release date | End of life                   | Latest version | Changelog
--------|--------------|-------------------------------|----------------|-----------
2.7-dev | ~2022-Q4     | 2024-Q1 (dev » stable)        | 2.7-dev0       | 2022/05/31
2.6     | 2022-05-31   | 2027-Q2 (LTS)                 | 2.6.0          | 2022/05/31
2.5     | 2021-11-23   | 2023-Q1 (stable)              | 2.5.7          | 2022/05/13
2.4     | 2021-05-14   | 2026-Q2 (LTS)                 | 2.4.17         | 2022/05/13
2.2     | 2020-07-07   | 2025-Q2 (LTS)                 | 2.2.24         | 2022/05/13
2.0     | 2019-06-16   | 2024-Q2 (critical fixes only) | 2.0.29         | 2022/05/13
1.8     | 2017-11-26   | 2022-Q4 (critical fixes only) | 1.8.30         | 2021/04/12

- version 2.6 : May, 31st, 2022  the latest long-term supported release.
- version 2.5 : runtime server addition/removal, runtime CA/CRL updates, native HTTP client, simplified HTTPS logging, default TCP/HTTP rulesets, JWT validation, and more
- version 2.4 : syslog and DNS over TCP, multi-threaded Lua, full sharing of idle conns, lower latency, server-side dynamic SSL update, Opentracing, WebSocket over H2, atomic maps, Vary support, new debugging tools, even - more user-friendly CLI and configuration, lots of cleanups
- version 2.3 : syslog forwarding, better idle conn management, improved balancing with large queues, simplified SSL managment, more stats metrics, stricter config checking by default, general performance improvements
- version 2.2 : runtime certificate additions, improved idle connection management, logging over TCP, HTTP "return" directive, errorfile templates, TLSv1.2 by default, extensible health-checks
- version 2.1 : improved I/Os and multi-threading, FastCGI, runtime certificate updates, HTX-only, improved debugging, removal of obsolete keywords
- version 2.0 : gRPC, layer 7 retries, process manager, SSL peers, log load balancing/sampling, end-to-end TCP fast-open, automatic settings (maxconn, threads, HTTP reuse, pools), ...
- version 1.9 : improved multi-threading, end-to-end HTTP/2, connection pools, queue priority control, stdout logging, ...
- version 1.8 : multi-threading, HTTP/2, cache, on-the fly server addition/removal, seamless reloads, DNS SRV, hardware SSL engines, ...
- version 1.7 : added server hot reconfiguration, content processing agents, multi-type certs, ...
- version 1.6 : added DNS resolution support, HTTP connection multiplexing, full stick-table replication, stateless compression, ...
- version 1.5 : added SSL, IPv6, keep-alive, DDoS protection, ...
