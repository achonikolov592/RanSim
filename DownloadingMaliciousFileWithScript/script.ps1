$url = "https://demo.wd.microsoft.com/Content/ransomware_testfile_unsigned.exe"

$dst = "C:/users/achon/onedrive/desktop/diplomna/rra/DownloadingMaliciousFileWithScript/out.exe"

Invoke-WebRequest -Uri $url -OutFile $dst