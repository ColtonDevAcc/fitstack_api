start:
	docker-compose up -d
stop:
	docker-compose down
upgrade_all_dep:
	go list -u -m all
test_all:
	go test all