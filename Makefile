APPNAME = domainchecker

all:
	go build .

install: all
	go install .

uninstall:
	rm $(GOPATH)/bin/$(APPNAME)
