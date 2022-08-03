
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

$version=Get-Content -Path "$PWD\configs\grpc.version" -ReadCount 1
$srcfile="protoc-$version-win64.zip"
$dstfile="$PWD\build\$srcfile"
$dstpath="$ENV:GOPATH"
$url="https://github.com/protocolbuffers/protobuf/releases/download/v$version/$srcfile"
$wc = New-Object System.Net.WebClient

New-Item -Path "$PWD\build" -ItemType Directory
Write-Output $wc.DownloadFileTaskAsync($url, $dstfile)
Expand-Archive -DestinationPath "$dstpath" -LiteralPath "$dstfile" -Force
Remove-Item -Path "$PWD\build" -Recurse -Force
