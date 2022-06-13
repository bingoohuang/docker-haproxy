up:
# do docker-compose up --build to force a rebuild.
	docker-compose up --build
clean:
	docker-compose rm -fsv
t1:
	curl http://127.0.0.1:6000
t2:
	for i in $$(seq 1 20); do curl http://127.0.0.1:6000; done
t3:
	gurl :6000 -n10
