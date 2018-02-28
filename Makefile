BUILDFILE="posty"
build:
	go build -o ${BUILDFILE} 
	@echo "posty build complete"
build_linux:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${BUILDFILE}_linux
	@echo "creating linux 64bit executable"

