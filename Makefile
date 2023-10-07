.PHONY: api
# generate api
api:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
 	       --go-http_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
		   --go-errors_out=paths=source_relative:./api \
	       --openapi_out=fq_schema_naming=true,default_response=false:. \
	       shared/shared.proto
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