run:
	weaver generate .
	SERVICEWEAVER_CONFIG=config.toml go run .

run-single:
	make run

run-multi:
	go build .
	weaver multi deploy config.toml