#
# Makefile for Bar2D2
#
# Use target 'setup' to download needed dependencies
#
# Use target 'build' or no target to build everything.
#

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
GPIO := gpio

# List building
ALL_LIST = $(DTO_LIST) $(PERSISTENCE) $(GPIO) $(SERVICES) $(UI)

CPU_TYPE = $(shell uname -m)
ifeq ($(CPU_TYPE),armv7l)
	GPIO_LIB_GET = $(GOGET) github.com/stianeikeland/go-rpio 
else
	GPIO_LIB_GET =
endif


BUILD_LIST = $(foreach int, $(ALL_LIST), $(int)_build) 
CLEAN_LIST = $(foreach int, $(ALL_LIST), $(int)_clean)
INSTALL_LIST = $(foreach int, $(ALL_LIST), $(int)_install)
TEST_LIST = $(foreach int, $(ALL_LIST), $(int)_test)
FMT_TEST = $(foreach int, $(ALL_LIST), $(int)_fmt)

all: $(info CPU is $(CPU_TYPE)) build


build: $(BUILD_LIST) 

clean: $(CLEAN_LIST)
install: $(INSTALL_LIST)
test: $(TEST_LIST)
fmt: $(FMT_TEST)

setup: 
	$(GOGET) github.com/mattn/go-sqlite3
	$(GPIO_LIB_GET)

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
    

    
