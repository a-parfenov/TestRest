.PHONY: run clean

run:
	go run src/*.go

clean:
	rm -rf savedFiles
	rm -rf downloadFiles
