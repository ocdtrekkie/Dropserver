cd "$(dirname "$0")"/../ || exit

if [ "$1" == "debug" ]; then
	echo Compiling for debug...
	go build -gcflags="all=-N -l" -o bin/ds-host cmd/ds-host/ds-host.go
else
	go build -o bin/ds-host cmd/ds-host/ds-host.go
fi
