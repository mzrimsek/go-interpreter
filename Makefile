GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
SRC_DIR=src/zip/
BINARY_NAME=main
TEST_FILE=test.zp

all: test build
install:
	cd $(SRC_DIR) && $(GOGET)
build: 
	cd $(SRC_DIR) && $(GOBUILD) -o $(BINARY_NAME) -v
test: 
	cd $(SRC_DIR) && $(GOTEST) -v ./...
clean: 
	cd $(SRC_DIR) && $(GOCLEAN) && rm -f $(BINARY_NAME)
run:
	cd $(SRC_DIR) && $(GOBUILD) -o $(BINARY_NAME) -v && ./$(BINARY_NAME)
file:
	cd $(SRC_DIR) && $(GOBUILD) -o $(BINARY_NAME) -v && ./$(BINARY_NAME) $(TEST_FILE)