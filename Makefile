BUILD_FILES = $(shell go list -f '{{range .GoFiles}}{{$$.Dir}}/{{.}}\
{{end}}' ./...)

GH_VERSION ?= $(shell git describe --tags 2>/dev/null || git rev-parse --short HEAD)
DATE_FMT = +%d-%m-%Y

ifdef SOURCE_DATE_EPOCH
    BUILD_DATE ?= $(shell date -u -d "@$(SOURCE_DATE_EPOCH)" "$(DATE_FMT)" 2>/dev/null || date -u -r "$(SOURCE_DATE_EPOCH)" "$(DATE_FMT)" 2>/dev/null || date -u "$(DATE_FMT)")
else
    BUILD_DATE ?= $(shell date "$(DATE_FMT)")
endif

GO_LDFLAGS := -X github.com/csothen/go-gen/internal/build.Version=$(GH_VERSION)
GO_LDFLAGS := -X github.com/csothen/go-gen/internal/build.Date=$(BUILD_DATE)

bin/go-gen: $(BUILD_FILES)
	go build -trimpath -ldflags "${GO_LDFLAGS}" -o "$@" .
