# ===================================
# SETTINGS
# ===================================

PROJECT := iteragit.iteratec.de/traze/goclient
MAIN := goclient.go
EXE := traze-goclient

# ===================================
# TARGETS
# ===================================

build:
	go build -o $(EXE) $(MAIN)

run:
	go run $(MAIN)

lint:
	golint $(PROJECT)

vet:
	go vet -v $(PROJECT)

clean:
	rm -f $(EXE)