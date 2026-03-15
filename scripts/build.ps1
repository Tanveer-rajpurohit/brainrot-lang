Write-Host "Building BrainRot Lang..." -ForegroundColor Yellow

New-Item -ItemType Directory -Force -Path dist | Out-Null

Write-Host "-> Windows" -ForegroundColor Cyan
$env:GOOS="windows"; $env:GOARCH="amd64"; go build -o dist/brainrot.exe main.go

Write-Host "-> Linux" -ForegroundColor Cyan
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o dist/brainrot-linux main.go

Write-Host "-> Mac Intel" -ForegroundColor Cyan
$env:GOOS="darwin"; $env:GOARCH="amd64"; go build -o dist/brainrot-mac main.go

Write-Host "-> Mac Apple Silicon" -ForegroundColor Cyan
$env:GOOS="darwin"; $env:GOARCH="arm64"; go build -o dist/brainrot-mac-arm main.go

# Reset env
$env:GOOS=""; $env:GOARCH=""

Write-Host "All builds complete -> dist/" -ForegroundColor Green