include $(GOROOT)/src/Make.inc

GOFMT = $(GOROOT)/bin/gofmt -tabwidth=4 -spaces=true -tabindent=false -w

TARG = date
GOFILES = daycount.go \
          date.go

GOTESTFILES = daycount_test.go

include $(GOROOT)/src/Make.pkg

format: $(GOFILES) $(GOTESTFILES)
	$(GOFMT) $?

