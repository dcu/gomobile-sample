test: deps
	go test .

ios: deps
	gomobile bind -target=ios

deps:
	glide install

