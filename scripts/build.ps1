Param(
    [ValidateSet('Release','Debug')]
    [string]$Configuration = 'Release',

    [ValidateSet('amd64','arm64','386')]
    [string]$Arch = 'amd64',

    [string]$OutputDir = 'dist'
)

$ErrorActionPreference = 'Stop'

# Resolve paths
$Root = Split-Path -Parent $PSScriptRoot
$Project = Join-Path $Root 'go\cmd\apc-mini-bridge'
$OutDir = Join-Path $Root $OutputDir
if (-not (Test-Path $OutDir)) { New-Item -Type Directory -Force -Path $OutDir | Out-Null }

# Environment for Windows build
$env:GOOS = 'windows'
$env:GOARCH = $Arch

# Build flags
$ldflags = ''
if ($Configuration -eq 'Release') { $ldflags = '-s -w' }

$exe = Join-Path $OutDir 'apc-mini-bridge.exe'

Write-Host "Building apc-mini-bridge ($Configuration, $Arch) -> $exe"

pushd $Project | Out-Null
try {
    go build -trimpath @('-ldflags', $ldflags) -o $exe .
}
finally {
    popd | Out-Null
}

Write-Host "Build complete: $exe"
