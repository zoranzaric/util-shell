SRC:= $(wildcard src/*.go)
BIN := $(notdir $(SRC:.go=))

all: $(BIN)
%: src/%.go
	mkdir -p build
	go build -o build/$@ $<
fmt:
	gofmt -tabs=true -s -w src
install:
	install -D build/string $(DESTDIR)/usr/bin/string
	install -D build/array $(DESTDIR)/usr/bin/array
	install -D build/math $(DESTDIR)/usr/bin/math
	install -D build/map $(DESTDIR)/usr/bin/map
