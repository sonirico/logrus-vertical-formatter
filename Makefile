# Variables
PACKAGE_NAME = logrus-vertical-formatter
SRC_DIR = ./
TESTS_DIR = ./

# Commands
GO = go
MKDIR = mkdir -p
RM = rm -rf

# Targets
.PHONY: all clean test fmt

all: clean test fmt

clean:
	$(RM) ./cover
	$(RM) ./doc
	$(RM) ./pkg

test:
	$(GO) test -race $(SRC_DIR)/...

fmt:
	$(GO) fmt $(SRC_DIR)/...
