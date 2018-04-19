FROM microsoft/windowsservercore:1709
SHELL ["powershell", "-command"]
RUN mkdir /windows-dns-test
WORKDIR /windows-dns-test
ADD windows-dns-test.exe .
ADD tail.ps1 .
ENTRYPOINT /windows-dns-test/windows-dns-test.exe