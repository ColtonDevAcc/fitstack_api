.PHONY: postgres adminer migrate

postgres:
	docker run --rm -ti --network host -e POSTGRES_PASWORD=secret postgres

adminer:
	docker run --rm -ti --network host adminer

migrate:
	migrate -srouce file://migrations \ 
					-database postgres://postgres:secret@localhost/postgres?sslmode=disable up

upgrade_all_dep:
	go list -u -m all