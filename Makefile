TIMESTAMP	:= $(shell /bin/date "+%F %T")
NAME		:= jing
VERSION		:= v1.0.0

default:
	@echo "no default target"; false

fmt:
	@go fmt $(CURDIR)/... &> /dev/null

test:
	@go test $(CURDIR)/...

github: test fmt
	@git add .
	@git commit -m "$(TIMESTAMP)"
	@git push

.PHONY: usage fmt test github