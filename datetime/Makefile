include $(GOROOT)/src/Make.inc

GOFMT = $(GOROOT)/bin/gofmt -tabwidth=4 -spaces=true -tabindent=false -w

TARG = datetime
GOFILES = ansidate.go \
          date.go \
          datetime.go

GOTESTFILES = ansidate_test.go

include $(GOROOT)/src/Make.pkg

format: $(GOFILES) $(GOTESTFILES)
	$(GOFMT) $?

