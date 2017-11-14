#!/usr/bin/env bash
set -e

name="$1"
if [[ -z $name ]]; then
	echo "Usage: $0 name version"
	exit 1
fi

version="$2"
if [[ -z $version ]]; then
	echo "Usage: $0 name version"
	exit 1
fi

for os in darwin linux windows; do
	folder=release/$name-$version-$os-amd64
	rm -rf $folder
	mkdir -p $folder
	bin="$name"
	if [ "$os" == "windows" ]; then
		bin="${name}.exe"
	fi
	echo "building ${os}..."

	GOOS=${os} GOARCH=amd64 go build -a -tags 'netgo' -o $folder/$bin .
	cp LICENSE README.md $folder

	(
		cd release
		zip -rq $name-$version-$os-amd64.zip $name-$version-$os-amd64
	)
	if [ $os = "linux" ]
	then
		folder=release/$name-$version-$os-arm64
		rm -rf $folder
		mkdir -p $folder
		GOOS=${os} GOARCH=arm64 go build -a -tags 'netgo' -o $folder/$bin .
		cp LICENSE README.md $folder

		(
			cd release
			zip -rq $name-$version-$os-arm64.zip $name-$version-$os-arm64
		)
	fi
done

echo "done"

