.PHONY: run health
run:
	go run .

health:
	curl -fsS http://127.0.0.1:8081/health
