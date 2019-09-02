#!/bin/sh
VERSION=$(cat config/version.go  | grep 'Version.*=.*v[0-9]' | cut -d '"' -f 2 | sed -r 's/v//g')
export VERSION
go get github.com/mitchellh/gox || true
gox -osarch="linux/amd64"
gox -osarch="windows/amd64"
gox -osarch="darwin/amd64"
mv *_linux_amd64 acentera
tar -czvf acentera-linux-amd64.tar.gz acentera 
mv *_darwin_amd64 acentera
tar -czvf acentera-darwin-amd64.tar.gz acentera 
mv acenteracli_windows_amd64.exe acentera.exe
zip -y acentera-windows-amd64.zip acentera.exe

[ -e acentera ] && rm -f acentera
[ -e acentera.exe ] && rm -f acentera.exe

git tag v${VERSION}
git push
git push --tags
 
github-release release --security-token ${GITHUB_TOKEN} --user ACenterA --repo acenteracli --tag v${VERSION}

github-release upload --security-token ${GITHUB_TOKEN} --user ACenterA --repo acenteracli \
    --tag v${VERSION} --name acentera-linux-amd64.tar.gz --file acentera-linux-amd64.tar.gz

github-release upload --security-token ${GITHUB_TOKEN} --user ACenterA --repo acenteracli \
    --tag v${VERSION} --name acentera-darwin-amd64.tar.gz --file acentera-darwin-amd64.tar.gz

github-release upload --security-token ${GITHUB_TOKEN} --user ACenterA --repo acenteracli \
    --tag v${VERSION} --name acentera-windows-amd64.zip --file acentera-windows-amd64.zip
