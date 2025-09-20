install:
	go build -o ggit
	strip ggit
	cp ggit /usr/local/bin/ggit
	rm ggit
