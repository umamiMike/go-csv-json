BUILDFILE="go-csv-json"
PHONY="build"
build:
	go build -o ${BUILDFILE} 
	@echo "${BUILDFILE} build complete"
build-shrink:
	go build -o ${BUILDFILE} 
	upx --brute ${BUILDFILE}
	@echo "${BUILDFILE}  build complete"
build-linux:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./linux/${BUILDFILE}
	upx --brute ./linux/${BUILDFILE}
	@echo "creating linux 64bit executable"

