default: build

build:
	'$(CURDIR)/scripts/src.build.sh'

clean:
	'$(CURDIR)/scripts/src.clean.sh'

.PHONY: default build clean