#!/bin/sh
#
# To use this script, install github-release via:
#   go get github.com/aktau/github-release
# Then verify GITHUB_TOKEN variable is set
#   export GITHUB_TOKEN=XXXXXXX (your personal token)
#
# Example:
# ./scripts/build-all.sh release v0.7.3
#
#
set -e
export GITHUB_USER=ecsteam
export GITHUB_REPO=cf_get_events

if [[ "$1" = "release" ]] ; then
	TAG="$2"
	: ${TAG:?"Usage: build_all.sh [release] [TAG]"}


	if ${github-release} info --tag $TAG > /dev/null 2>&1 ; then
		echo "$TAG exists, remove it or increment"
		exit 1
	else
		MAJOR=`echo $TAG | sed 's/^v//' | awk 'BEGIN {FS = "." } ; { printf $1;}'`
		MINOR=`echo $TAG | sed 's/^v//' | awk 'BEGIN {FS = "." } ; { printf $2;}'`
		BUILD=`echo $TAG | sed 's/^v//' | awk 'BEGIN {FS = "." } ; { printf $3;}'`

		`sed -i "" -e "1,/Major:.*/s/Major:.*/Major: $MAJOR,/" \
			-e "1,/Minor:.*/s/Minor:.*/Minor: $MINOR,/" \
			-e "1,/Build:.*/s/Build:.*/Build: $BUILD,/" cf_get_events.go`
	fi
fi

LINUX_FILENAME="get-events-plugin-linux"
MAC_FILENAME="get-events-plugin-darwin"
WIN_FILENAME="get-events-plugin.exe"

GOOS=linux GOARCH=amd64 go build -o $LINUX_FILENAME
LINUX64_SHA1=`cat $LINUX_FILENAME | openssl sha1`
mkdir -p bin/linux64
mv $LINUX_FILENAME bin/linux64

GOOS=darwin GOARCH=amd64 go build -o $MAC_FILENAME
OSX_SHA1=`cat $MAC_FILENAME | openssl sha1`
mkdir -p bin/osx
mv $MAC_FILENAME bin/osx

GOOS=windows GOARCH=amd64 go build -o $WIN_FILENAME
WIN64_SHA1=`cat $WIN_FILENAME | openssl sha1`
mkdir -p bin/win64
mv $WIN_FILENAME bin/win64

NOW=`TZ=UC date +'%Y-%m-%dT%TZ'`

cat repo-index.yml |
sed "s/__osx-sha1__/$OSX_SHA1/" |
sed "s/__win64-sha1__/$WIN64_SHA1/" |
sed "s/__linux64-sha1__/$LINUX64_SHA1/" |
sed "s/__TAG__/$TAG/" |
sed "s/__TODAY__/$NOW/" |
cat

export github_release="/usr/local/bin/github-release";
if [[ "$1" = "release" ]] ; then

	git commit -am "Build version $TAG"
	git push

	${github_release} release \
    --tag $TAG \
    --name "Cloud Foundry get events plugin $TAG" \
    --description "$TAG release"
#    --description "$TAG release - work in progress"

	${github_release} upload \
    --tag $TAG \
    --name "get-events-plugin-darwin" \
    --file bin/osx/get-events-plugin-darwin

	${github_release} upload \
    --tag $TAG \
    --name "get-events-plugin-linux" \
    --file bin/linux64/get-events-plugin-linux

	${github_release} upload \
    --tag $TAG \
    --name "get-events-plugin.exe" \
    --file bin/win64/get-events-plugin.exe

	#git commit -am "Build version $TAG"
	#git tag -a $TAG -m "Get Events Plugin v$TAG"
	#echo "Tagged release, 'git push --follow-tags' to push it to github, upload the binaries to github"
	echo "copy the output above to the cli repo you plan to deploy in"
fi
