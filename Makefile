TIMESTAMP	:= $(shell /bin/date "+%F %T")
NAME		:= jing
VERSION		:= v1.0.0

fmt:
	@go fmt $(CURDIR)/... &> /dev/null

clean:
	@true

github: clean fmt
	@git add .
	@git commit -m "$(TIMESTAMP)"
	@git push

.PHONY: usage fmt clean github
