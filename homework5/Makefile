all: build test cover run
build: 
        $(GOBUILD) -o $(BINARY_NAME) -v
test: 
        $(GOTEST) -v ./...
cover:
        $(GOTEST) ./... -coverprofile=c.out && go tool cover -html=c.out
clean: 
        $(GOCLEAN)
        rm -f $(BINARY_NAME)
        rm -f c.out
run:
        $(GOBUILD) -o $(BINARY_NAME) -v ./...
        ./$(BINARY_NAME)
lint:
        golangci-lint run