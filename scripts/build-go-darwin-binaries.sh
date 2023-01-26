#!/usr/bin/env bash

package=cmd/sedge/main.go
package_name=sedge

platforms=("darwin/amd64" "darwin/arm64")

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	output_name=$package_name'-'$VERSION'-'$GOOS'-'$GOARCH

    LDFLAGS="-X github.com/NethermindEth/sedge/internal/utils.Version=${VERSION}"

	env CGO_ENABLED=1 GOOS=$GOOS GOARCH=$GOARCH go build -ldflags "${LDFLAGS}" -o build/"$output_name" $package
	if [ $? -ne 0 ]; then
   		echo 'An error has occurred! Aborting the script execution...'
		exit 1
	fi
	echo "Generated ${output_name} file"
done