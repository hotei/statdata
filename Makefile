# Makefile for statdata package

PROG = statdata
VERSION = 0.1.0
TARDIR = 	$(HOME)/Desktop/TarPit/
DATE = 	`date "+%Y-%m-%d.%H_%M_%S"`
DOCOUT = README-$(PROG)-godoc.md

all:
	go build

install:
	go build
	go tool vet .
	go tool vet -shadow .
	gofmt -w *.go
	go install
	godoc2md . > $(DOCOUT)
	godepgraph -md -p . >> $(DOCOUT)
	deadcode -md >> $(DOCOUT)
	cp README-$(PROG).md README.md
	cat $(DOCOUT) >> README.md
	cp README.md README2.md
#	cp $(PROG) $(HOME)/bin
	
neat:
	go fmt ./...

dead:
	deadcode > problems.dead

index:
	cindex .

clean:
	go clean ./...
	rm -f *~ problems.dead count.out README2.md
#	rm -f $(DOCOUT)

tar:
	echo $(TARDIR)$(PROG)_$(VERSION)_$(DATE).tar
	tar -ncvf $(TARDIR)$(PROG)_$(VERSION)_$(DATE).tar .

# Coverage test maker
cover:
	go test -covermode=count -coverprofile=count.out
	cover -html=count.out