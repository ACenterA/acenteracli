#!/bin/sh
go get github.com/mitchellh/gox || true
gox -osarch="linux/amd64"
gox -osarch="windows/amd64"
gox -osarch="darwin/amd64"
mv awless_linux_amd64 acentera
tar -czvf acentera-linux-amd64.tar.gz acentera 
mv awless_darwin_amd64 acentera
tar -czvf acentera-darwin-amd64.tar.gz acentera 
mv awless_windows_amd64 acentera
tar -czvf acentera-windows-amd64.tar.gz acentera 
[ -e acentera ] && rm -f acentera


git tag v${VERSION}
git push
git push --tags
 
github-release release --security-token ${GITHUB_TOKEN} --user ACenterA --repo acenteracli --tag v${VERSION}

github-release upload --security-token ${GITHUB_TOKEN} --user ACenterA --repo acenteracli \
    --tag v${VERSION} --name acentera-linux-amd64.tar.gz --file acentera-linux-amd64.tar.gz

github-release upload --security-token ${GITHUB_TOKEN} --user ACenterA --repo acenteracli \
    --tag v${VERSION} --name acentera-darwin-amd64.tar.gz --file acentera-darwin-amd64.tar.gz

github-release upload --security-token ${GITHUB_TOKEN} --user ACenterA --repo acenteracli \
    --tag v${VERSION} --name acentera-windows-amd64.tar.gz --file acentera-windows-amd64.tar.gz
