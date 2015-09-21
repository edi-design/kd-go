#!/bin/bash

# linux, desktop
pushd $GOROOT/src/
GOPATH=$GOPATH:$PWD/../ GOOS=linux GOARCH=amd64 ./make.bash --no-clean 2> /dev/null 1> /dev/null
GOPATH=$GOPATH:$PWD/../ GOOS=linux GOARCH=386 ./make.bash --no-clean 2> /dev/null 1> /dev/null
popd
GOPATH=$GOPATH:$PWD/../ GOOS=linux GOARCH=amd64 go build -o ../build/linux/kd_proxy_64 ../src/kd.go
GOPATH=$GOPATH:$PWD/../ GOOS=linux GOARCH=386 go build -o ../build/linux/kd_proxy_32 ../src/kd.go

# windows desktop
pushd $GOROOT/src/
GOPATH=$GOPATH:$PWD/../ GOOS=windows GOARCH=amd64 ./make.bash --no-clean 2> /dev/null 1> /dev/null
GOPATH=$GOPATH:$PWD/../ GOOS=windows GOARCH=386 ./make.bash --no-clean 2> /dev/null 1> /dev/null
popd
GOPATH=$GOPATH:$PWD/../ GOOS=windows GOARCH=amd64 go build -o ../build/windows/kd_proxy_64.exe ../src/kd.go
GOPATH=$GOPATH:$PWD/../ GOOS=windows GOARCH=386 go build -o ../build/windows/kd_proxy_32.exe ../src/kd.go

# mac
pushd $GOROOT/src/
GOPATH=$GOPATH:$PWD/../ GOOS=darwin GOARCH=amd64 ./make.bash --no-clean 2> /dev/null 1> /dev/null
popd
GOPATH=$GOPATH:$PWD/../ GOOS=darwin  GOARCH=amd64 go build -o ../build/mac/kd_proxy ../src/kd.go

# linux armv5
pushd $GOROOT/src/
GOPATH=$GOPATH:$PWD/../ GOOS=linux  GOARCH=arm GOARM=5 ./make.bash --no-clean 2> /dev/null 1> /dev/null
popd
GOPATH=$GOPATH:$PWD/../ GOOS=linux  GOARCH=arm GOARM=5 go build -o ../build/arm/kd_proxy_arm_v5 ../src/kd.go

# linux armv6
pushd $GOROOT/src/
GOPATH=$GOPATH:$PWD/../ GOOS=linux  GOARCH=arm GOARM=6 ./make.bash --no-clean 2> /dev/null 1> /dev/null
popd
GOPATH=$GOPATH:$PWD/../ GOOS=linux  GOARCH=arm GOARM=6 go build -o ../build/arm/kd_proxy_arm_v6 ../src/kd.go

# linux armv7
pushd $GOROOT/src/
GOPATH=$GOPATH:$PWD/../ GOOS=linux  GOARCH=arm GOARM=7 ./make.bash --no-clean 2> /dev/null 1> /dev/null
popd
GOPATH=$GOPATH:$PWD/../ GOOS=linux  GOARCH=arm GOARM=7 go build -o ../build/arm/kd_proxy_arm_v7 ../src/kd.go

#power-pc 64bit
pushd $GOROOT/src/
GOPATH=$GOPATH:$PWD/../ GOOS=linux  GOARCH=ppc64 ./make.bash --no-clean 2> /dev/null 1> /dev/null
popd
GOPATH=$GOPATH:$PWD/../ GOOS=linux  GOARCH=ppc64 go build -o ../build/ppc/kd_proxy_64 ../src/kd.go

#power-pc 64bit low energy
pushd $GOROOT/src/
GOPATH=$GOPATH:$PWD/../ GOOS=linux  GOARCH=ppc64le ./make.bash --no-clean 2> /dev/null 1> /dev/null
popd
GOPATH=$GOPATH:$PWD/../ GOOS=linux  GOARCH=ppc64le go build -o ../build/ppc/kd_proxy_64_le ../src/kd.go