
# Go parameters
GOCMD=go
GOGET=$(GOCMD) get
GOBUILD=$(GOCMD) build 
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
GODEP=$(GOTEST) -i
GOFMT=gofmt -w

# Package lists
TOPLEVEL_PKG := bar2d2
DTO_LIST := dto
PERSISTENCE := persistence
SERVICES := services
UI := ui

# List building
ALL_LIST = $(DTO_LIST) $(PERSISTENCE) $(SERVICES) $(UI)

BUILD_LIST = $(foreach int, $(ALL_LIST), $(int)_build)
CLEAN_LIST = $(foreach int, $(ALL_LIST), $(int)_clean)
INSTALL_LIST = $(foreach int, $(ALL_LIST), $(int)_install)
TEST_LIST = $(foreach int, $(ALL_LIST), $(int)_test)
FMT_TEST = $(foreach int, $(ALL_LIST), $(int)_fmt)

CPU_TYPE = $(shell uname -m)

all: $(info CPU is $(CPU_TYPE)) build

build: $(BUILD_LIST)
clean: $(CLEAN_LIST)
install: $(INSTALL_LIST)
test: $(TEST_LIST)
fmt: $(FMT_TEST)

setup: 
	$(GOGET) github.com/mattn/go-sqlite3

$(BUILD_LIST): %_build:  
	cd ./$*; $(GOBUILD) .
$(CLEAN_LIST): %_clean:
	$(GOCLEAN) ./$*
$(INSTALL_LIST): %_install:
	$(GOINSTALL) ./$*
$(TEST_LIST): %_test:
	$(GOTEST) ./$*
$(FMT_TEST): %_fmt:
	$(GOFMT) ./$*
    

    