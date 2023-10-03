.PHONY: api
# generate api
api:
	find app -type d -depth 1 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) api'

.PHONY: config
# generate config
config:
	find app -type d -depth 1 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) config'


.PHONY: generate
# generate wire
generate:
	find app -type d -depth 1 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) generate'

.PHONY: all
# all 
all:
	find app -type d -depth 1 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) all'

.PHONY: build
# build 
build:
	find app -type d -depth 1 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) build'

mail:
	docker-compose -f deploy/docker-compose/docker-compose.yaml up