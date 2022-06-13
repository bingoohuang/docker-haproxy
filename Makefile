up:
# do docker-compose up --build to force a rebuild.
	docker-compose up --build
clean:
	docker-compose rm -fsv
test:
	for i in $$(seq 1 20); do curl http://127.0.0.1:8080; done
test2:
	gurl :8080 -n10
