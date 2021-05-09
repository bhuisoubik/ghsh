# Install on windows

$f_url = "https://raw.githubusercontent.com/soubikbhuiwk007/soubikbhuiwk007/master/dist/ghsh/ghsh.exe"

$out_path = "c:\GithubShell\ghsh.exe" # executable file
$out_dir = "C:\GithubShell" # Directory for installation

# See if $out_dir doesn't exists
if ((Test-Path $out_dir) -bxor 1) {
    Write-Output "Creating $out_dir ..."
    New-Item -Path "c:\" -Name "GithubShell" -ItemType "directory" | Out-Null
    Write-Output "Completed"
}

# See if $out_file doesn't exist
if ((Test-Path $out_path) -bxor 1) {
    New-Item -ItemType "file" -Path $out_path | Out-Null
}

# Download the file / Also works as a update if you already have one installed
Write-Output "Fetching executable..."
Invoke-WebRequest $f_url -OutFile $out_path
Write-Output "Completed"

# add to path
[System.Environment]::SetEnvironmentVariable("Path", $env:Path + ";$out_dir", "User")
Write-Output "Added to Path"