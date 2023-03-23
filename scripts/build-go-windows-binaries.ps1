$package = "cmd/sedge/main.go"
$packageName = "sedge"

$platforms = @("windows/amd64")

if (-not(Test-Path "build")) {
    New-Item "build" -ItemType Directory | Out-Null
}

foreach ($platform in $platforms) {
    $parts = $platform.Split("/")
    $GOOS = $parts[0]
    $GOARCH = $parts[1]
    $outputName = "$packageName-$VERSION-$GOOS-$GOARCH.exe"
    $ldflags = "-X github.com/NethermindEth/sedge/internal/utils.Version=$VERSION -extldflags -static"

    $env:GOOS = $GOOS ; $env:GOARCH = $GOARCH ; $env:CGO_ENABLED = 1 ; go build -ldflags "${ldflags}" -o .\build\$outputName $package 
    
    if ($? -ne $true) {
        Write-Output "An error has occurred! Aborting the script execution..."
        Exit 1
    }
    Write-Output "Generated $outputName file"
}


