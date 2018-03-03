BUILDFILE="posty"
build:
	go build -o ${BUILDFILE} 
	@echo "posty build complete"
build-shrink:
	go build -o ${BUILDFILE} 
	upx --brute ${BUILDFILE}
	@echo "posty build complete"
build-linux:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./linux/${BUILDFILE}
	upx --brute ./linux/${BUILDFILE}
	@echo "creating linux 64bit executable"

