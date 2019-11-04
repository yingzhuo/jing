TIMESTAMP	:= $(shell /bin/date "+%F %T")
NAME		:= jing

no_default:
	@echo "no default target"; false

fmt:
	@go fmt $(CURDIR)/... &> /dev/null

test:
	@go test $(CURDIR)/...

github: fmt
	@git add .
	@git commit -m "$(TIMESTAMP)"
	@git push

.PHONY: no_default fmt test github