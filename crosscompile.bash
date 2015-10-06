#!/bin/bash
##
# kd-go build-script
##
PACKAGE="kd-proxy"

type setopt >/dev/null 2>&1 && setopt shwordsplit
PLATFORMS="darwin/386 darwin/amd64 freebsd/386 freebsd/amd64 freebsd/arm/5 freebsd/arm/6 freebsd/arm/7 linux/386 linux/amd64 linux/arm/5 linux/arm/6 linux/arm/7 linux/arm/arm64 linux/ppc64 linux/ppc64le windows/386 windows/amd64 openbsd/386 openbsd/amd64"

FAILURES=""
for PLATFORM in $PLATFORMS; do
	GOOS=${PLATFORM%%/*}
	GOARCHWITHTYPE=${PLATFORM#*/}
	GOARCH=${GOARCHWITHTYPE%/*}
	GOARCHTYPE=${PLATFORM##*/}
	OUTPUT="build/$GOOS/${PACKAGE}-${GOARCH}"

	TYPE=""
	if [[ $GOARCHTYPE != $GOARCH ]] ; then
		TYPE="GO$(echo $GOARCH | tr '[a-z]' '[A-Z]')=${GOARCHTYPE}"
		OUTPUT="$OUTPUT-$GOARCHTYPE"
	fi

	# prepare cross-compile env for platform
	PREPARECMD="
		pushd ${GOROOT}/src/ > /dev/null ;
		GOPATH=${GOPATH}:${PWD}/../../../../ GOOS=${GOOS} GOARCH=${GOARCH} ${TYPE} ./make.bash --no-clean 2> /dev/null 1> /dev/null ;
		popd > /dev/null"

	# build for platform
	BUILDCMD="GOPATH=${GOPATH}:${PWD}/../../../../ GOOS=${GOOS} GOARCH=${GOARCH} ${TYPE} go build -o ${OUTPUT} *.go"

	echo "- preparing $PLATFORM"
	eval $PREPARECMD || RPEPFAILURES="$RPEPFAILURES $PLATFORM"
	echo "-- done"

	echo "- building $PLATFORM"
	eval $BUILDCMD || FAILURES="$FAILURES $PLATFORM"
	echo "-- done: $OUTPUT"
done

if [ "$RPEPFAILURES" != "" ]; then
	echo "*** prepare FAILED on $RPEPFAILURES ***"
fi

if [ "$FAILURES" != "" ]; then
	echo "*** build FAILED on $FAILURES ***"
fi