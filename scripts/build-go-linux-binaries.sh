#!/usr/bin/env bash
set -x
package=cmd/sedge/main.go
package_name=sedge

platforms=("linux/amd64" "linux/arm64")

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	output_name=$package_name'-v'$VERSION'-'$GOOS'-'$GOARCH

  LDFLAGS="-X github.com/NethermindEth/sedge/internal/utils.Version=v${VERSION}"

  docker buildx build --platform="$GOOS"/"$GOARCH" -t nethermindeth/sedge:"$VERSION"-"$GOOS"-"$GOARCH" --build-arg TARGETOS="$GOOS" --build-arg TARGETARCH="$GOARCH" --build-arg LDFLAGS="$LDFLAGS" --build-arg OUTPUT_NAME="$output_name" --build-arg PACKAGE="$package" --load scripts 
  docker create --name sedge nethermindeth/sedge:"$VERSION"-"$GOOS"-"$GOARCH"
  docker cp sedge:/sedge ./build/"$output_name"
  docker rm -f sedge

	if [ $? -ne 0 ]; then
   		echo 'An error has occurred! Aborting the script execution...'
		exit 1
	fi
	echo "Generated ${output_name} file"
done