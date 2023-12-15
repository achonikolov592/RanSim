$url = "https://demo.wd.microsoft.com/Content/ransomware_testfile_unsigned.exe"

$dst = "C:/users/acho/desktop/diplomna/DownloadingMaliciousFileWithScript/out.exe"

Invoke-WebRequest -Uri $url -OutFile $dst