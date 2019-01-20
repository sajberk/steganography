
# Happy hacking!

.PHONY: clean test install

install:
	go install

test:
	go test ./...

clean:
	rm images/new_ponyo.jpg