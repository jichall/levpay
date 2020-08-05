CC=`which go`

SOURCES=$(wildcard src/*.go)
SOURCE_FOLDER=src

BIN_FOLDER=bin
BIN_NAME=hero.out

all:
	$(CC) build -o $(BIN_FOLDER)/$(BIN_NAME) $(SOURCES)

tests:
	$(CC) test -v $(SOURCE_FOLDER)


vars:
	@echo CC     	 = $(CC)
	@echo SOURCES 	 = $(SOURCES)
	@echo SRC_FOLDER = $(SOURCE_FOLDER)
	@echo BIN_FOLDER = $(BIN_FOLDER)