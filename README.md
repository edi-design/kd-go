# kd-go
GO-Version of the KabelDeutschland streaming proxy (https://github.com/edi-design/kd-streaming-proxy)

## dependencies

* gorilla mux router (https://github.com/gorilla/mux)
 * `go get github.com/gorilla/mux`

## build

To build the source for various plattforms, simply run the build-script provided under `scripts/build.sh`.
It relies on a proper configured go-environment with GOPATH and GOROOT set up.

By now, I only tested the build process on OS X but it should not be a problem to run the commands on Linux.

## configuration

Place a `config.json` next to the binary and fill it with the following content, including your KabelDeutschland credentials.
```
{
  "Service": {
    "Username": "##USERNAME##",
    "Password": "##PASSWORD##",
    "Listen": ":8787"
  }
}
```

## run

The easiest way, is to run the binary without any params. It searches automatically for the `config.json` next to the binary.
`./kd_go`

### params

```
# ./kd_proxy -h
you need to set the following params:
  -c string
    	specifiy the config.json location, if not next to binary
  -h	display help message
  -no-cache
    	disables playlist caching
  -no-check-certificate
    	disable root CA check for HTTP requests
  -v	enable verbose mode to see more debug output.
  -version
    	shows the current version number.
```
