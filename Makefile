GO=go

all: vendor build test

vendor:
	${GO} mod tidy && ${GO} mod vendor

build:
	${GO} build -mod=vendor

test:
	${GO} test -mod=vendor -timeout=90s -p 1 -race -count=1 ./... -coverprofile cover.out

# doc:
# 	${GO} run -mod=vendor internal/im/main.go