TIMESTAMP	:= $(shell /bin/date "+%F %T")
NAME		:= jing

default:
	@echo "no default target"; false

fmt:
	@go fmt $(CURDIR)/... &> /dev/null

test:
	@go test $(CURDIR)/...

github: fmt
	@git add .
	@git commit -m "$(TIMESTAMP)"
	@git push

.PHONY: usage fmt test github