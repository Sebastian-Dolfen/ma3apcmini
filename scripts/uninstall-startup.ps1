Param(
    [string]$TaskName = 'APC Mini Bridge'
)

$ErrorActionPreference = 'Stop'

if (Get-ScheduledTask -TaskName $TaskName -ErrorAction SilentlyContinue) {
    Unregister-ScheduledTask -TaskName $TaskName -Confirm:$false | Out-Null
    Write-Host "Removed startup task '$TaskName'"
} else {
    Write-Host "Startup task '$TaskName' not found"
}
