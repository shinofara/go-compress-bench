bench:
	go test -bench . -benchmem

sample-data:
	sh makeTestFiles.sh
