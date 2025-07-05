SRC := ./main.go
BINARY := drink-selector-tui
OUTPUT := ./build/

build-prod: $(SRC)
	@GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(OUTPUT)/$(BINARY)-prod $(SRC)

build-dev: $(SRC)
	@GOOS=linux GOARCH=amd64 go build -o $(OUTPUT)/$(BINARY)-dev $(SRC)

# MacOS用ビルド
build-mac: $(SRC)
	@GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o $(OUTPUT)/$(BINARY)-mac $(SRC)

build-mac-dev: $(SRC)
	@GOOS=darwin GOARCH=arm64 go build -o $(OUTPUT)/$(BINARY)-mac-dev $(SRC)

clean:
	@rm -rf $(OUTPUT)/$(BINARY)*
