Param(
    [string]$ExePath = (Join-Path (Split-Path -Parent $PSScriptRoot) 'dist\apc-mini-bridge.exe'),
    [string]$TaskName = 'APC Mini Bridge',
    [string]$WorkingDir = (Split-Path -Parent $ExePath),
    [string]$ConfigPath = (Join-Path (Split-Path -Parent $PSScriptRoot) 'config.yaml'),
    [int]$DelaySeconds = 5
)

$ErrorActionPreference = 'Stop'

if (-not (Test-Path $ExePath)) {
    throw "Executable not found: $ExePath. Run scripts/build.ps1 first."
}

# Build the arguments (quote config path)
$Arguments = @('-config', '"' + $ConfigPath + '"') -join ' '

# Create Scheduled Task to run at user logon (highest privileges, delayed start)
$action = New-ScheduledTaskAction -Execute $ExePath -Argument $Arguments -WorkingDirectory $WorkingDir
$trigger = New-ScheduledTaskTrigger -AtLogOn
$principal = New-ScheduledTaskPrincipal -UserId $env:USERNAME -RunLevel Highest

# Optional delay wrapper via OnLogon is not directly supported; use repetition pattern or simply rely on app internal readiness
# Register or update task
if (Get-ScheduledTask -TaskName $TaskName -ErrorAction SilentlyContinue) {
    Unregister-ScheduledTask -TaskName $TaskName -Confirm:$false | Out-Null
}
Register-ScheduledTask -TaskName $TaskName -Action $action -Trigger $trigger -Principal $principal -Description 'Launch APC Mini Bridge at user logon' | Out-Null

Write-Host "Installed startup task '$TaskName' launching: $ExePath $Arguments"
