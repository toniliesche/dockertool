SHELL := bash

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	GOOS := "linux"
endif
ifeq ($(UNAME_S),Darwin)
	GOOS := "darwin"
endif

UNAME_M := $(shell uname -m)
ifeq ($(UNAME_M),aarch64)
	GOARCH := "arm64"
endif
ifeq ($(UNAME_M),x86_64)
	GOARCH := "amd64"
endif

ifneq ("$(wildcard $(CURDIR)/build.properties)","")
	include $(CURDIR)/build.properties
endif

PROFILE ?= generic
ifneq ("$(wildcard $(CURDIR)/profiles/profile.${PROFILE}.properties)","")
	include $(CURDIR)/profiles/profile.${PROFILE}.properties
else
	PROFILE = generic
	include $(CURDIR)/profiles/profile.generic.properties
endif

BUILD_DATE := $(shell date '+%Y-%m-%d %H:%M:%S')
BUILD_VERSION := ${build.version.major}.${build.version.minor}.${build.version.bugfix}-"DEV"
HASH := $(shell git rev-parse --short HEAD)
COMMIT_DATE := $(shell git show -s --format=%cd --date=format:'%y-%m-%d' $(HASH))
COMMIT := ${HASH} (${COMMIT_DATE})

build: build-info
	@echo Starting build
	@echo

	@GOOS=${GOOS} GOARCH=${GOARCH} go build -v -o dist/vmtool_${build.options.channel}_${GOOS}_${GOARCH} -ldflags="\
		-X 'github.com/toniliesche/dockertool/modules/build.AuthorName=${build.author.name}' \
		-X 'github.com/toniliesche/dockertool/modules/build.AuthorMail=${build.author.email}' \
		-X 'github.com/toniliesche/dockertool/modules/build.Channel=${build.options.channel}' \
		-X 'github.com/toniliesche/dockertool/modules/build.Copyright=${build.copyright.name}' \
		-X 'github.com/toniliesche/dockertool/modules/build.CopyrightYear=${build.copyright.year}' \
		-X 'github.com/toniliesche/dockertool/modules/build.Sudo=${build.options.sudo_enabled}' \
		-X 'github.com/toniliesche/dockertool/modules/build.SudoUser=${build.options.sudo_user}' \
		-X 'github.com/toniliesche/dockertool/modules/build.Commit=${COMMIT}' \
		-X 'github.com/toniliesche/dockertool/modules/build.Date=${BUILD_DATE}' \
		-X 'github.com/toniliesche/dockertool/modules/build.Version=${BUILD_VERSION}'" \
		cmd/cli.go

build-info:
	@echo
	@echo Finished preparations
	@echo
	@echo Summary:
	@echo ========
	@echo
	@echo Author: ${build.author.name}
	@echo Author E-Mail: ${build.author.email}
	@echo
	@echo Build Version: ${BUILD_VERSION}
	@echo "Build Commit: ${COMMIT}"
	@echo Build Date: ${BUILD_DATE}
	@echo
	@echo Build OS: ${GOOS}
	@echo Build Arch: ${GOARCH}
	@echo
	@echo Release Channel: ${build.options.channel}
	@echo
	@echo Sudo: ${build.options.sudo_enabled}
	@echo Sudo User: ${build.options.sudo_user}
	@echo

build-release-%:
	@$(MAKE) -s build-release PROFILE=$*

build-rc-%:
	@$(MAKE) -s build-rc PROFILE=$*

build-patch-%:
	@$(MAKE) -s build-patch PROFILE=$*

build-%: resolve-version-% build
	@echo "Finished build of dist/vmtool_${build.options.channel}_${GOOS}_${GOARCH}"
	@echo

resolve-version-release:
	@$(eval BUILD_VERSION := ${build.version.major}.${build.version.minor}.${build.version.bugfix})

resolve-version-rc:
	@$(eval BUILD_VERSION := ${build.version.major}.${build.version.minor}.${build.version.bugfix}-RC${build.version.candidate})

resolve-version-hotfix: ;
	@$(eval BUILD_VERSION := ${build.version.major}.${build.version.minor}.${build.version.bugfix}-p${build.version.patch})

clean:
	@rm -rf dist/vmtool*

run: run-m

run-%:
	@dist/vmtool_${build.options.channel}_${GOOS}_${GOARCH} $*

increase-%: update-% write-properties
	@echo updated build.properties file

update-major:
	@$(eval build.version.major := $(shell echo $$(($(build.version.major) + 1))))
	@$(eval build.version.minor := 0)
	@$(eval build.version.bugfix := 0)
	@$(eval build.version.candidate := 1)
	@$(eval build.version.patch := 1)
	@echo new major version: ${build.version.major}

update-minor:
	@$(eval build.version.minor := $(shell echo $$(($(build.version.minor) + 1))))
	@$(eval build.version.bugfix := 0)
	@$(eval build.version.candidate := 1)
	@$(eval build.version.patch := 1)
	@echo new minor version: ${build.version.minor}

update-bugfix:
	@$(eval build.version.bugfix := $(shell echo $$(($(build.version.bugfix) + 1))))
	@$(eval build.version.candidate := 1)
	@$(eval build.version.patch := 1)
	@echo new bugfix version: ${build.version.bugfix}

update-rc:
	@$(eval build.version.candidate := $(shell echo $$(($(build.version.candidate) + 1))))
	@echo new rc version: ${build.version.candidate}

update-patch:
	@$(eval build.version.patch := $(shell echo $$(($(build.version.patch) + 1))))
	@echo new patch version: ${build.version.patch}

write-properties:
	@echo "build.version.major=${build.version.major}" > build.properties.tmp
	@echo "build.version.minor=${build.version.minor}" >> build.properties.tmp
	@echo "build.version.bugfix=${build.version.bugfix}" >> build.properties.tmp
	@echo "build.version.candidate=${build.version.candidate}" >> build.properties.tmp
	@echo "build.version.patch=${build.version.patch}" >> build.properties.tmp
	@rm build.properties
	@mv build.properties.tmp build.properties
