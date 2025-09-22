install:
	go build -o ggit
	strip --strip-all ggit
	cp ggit /usr/local/bin/ggit
	rm ggit
