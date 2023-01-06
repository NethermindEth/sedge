$package = "cmd\sedge\main.go"
$packageName = "sedge"

$platforms = @("windows/amd64", "windows/arm64")

if (-not(Test-Path "build")) {
    New-Item "build" -ItemType Directory | Out-Null
}

$VERSION = "0.6.0"

foreach ($platform in $platforms) {
    $parts = $platform.Split("/")
    $GOOS = $parts[0]
    $GOARCH = $parts[1]
    $outputName = "$packageName-$VERSION-$GOOS-$GOARCH.exe"
    $ldflags = "-X github.com/NethermindEth/sedge/internal/utils.Version=$VERSION"

    docker buildx build --platform=$GOOS/$GOARCH -t nethermindeth/sedge:$VERSION-$GOOS-$GOARCH --build-arg TARGETOS=$GOOS --build-arg TARGETARCH=$GOARCH --build-arg LDFLAGS=$LDFLAGS --build-arg OUTPUT_NAME=$outputName --build-arg PACKAGE=$package --load . -f scripts/Dockerfile
    docker create --name sedge nethermindeth/sedge:$VERSION-$GOOS-$GOARCH
    docker cp sedge:/sedge ./build/$outputName
    docker rm -f sedge
    
    if ($? -ne $true) {
        Write-Output "An error has occurred! Aborting the script execution..."
        Exit 1
    }
    Write-Output "Generated $outputName file"
}


