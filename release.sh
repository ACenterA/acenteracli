#!/bin/sh
set -x
VERSION=$(cat config/version.go  | grep 'Version.*=.*v[0-9]' | cut -d '"' -f 2 | sed -r 's/v//g')
export VERSION
echo "Creating release v${VERSION}"
go get github.com/mitchellh/gox || true
go get -u -v github.com/aktau/github-release || true
ls -latrh
gox -osarch="linux/amd64"
ls -latrh
mv *_linux_amd64 acentera
ls -latrh
tar -czvf acentera-linux-amd64.tar.gz acentera 
rm -f acentera
gox -osarch="darwin/amd64"
ls -latrh
mv *_darwin_amd64 acentera
ls -latrh
tar -czvf acentera-darwin-amd64.tar.gz acentera 
ls -latrh
rm -f acentera
gox -osarch="darwin/arm64"
ls -latrh
mv *_darwin_arm64 acentera
ls -latrh
tar -czvf acentera-darwin-arm64.tar.gz acentera 
ls -latrh
rm -f acentera

gox -osarch="linux/arm64"
ls -latrh
mv *_linux_arm64 acentera
tar -czvf acentera-linux-arm64.tar.gz acentera 
rm -f acentera

gox -osarch="windows/amd64"
ls -latrh
mv *_windows_amd64.exe acentera.exe
zip -y acentera-windows-amd64.zip acentera.exe
rm -f acentera.exe

[ -e acentera ] && rm -f acentera
[ -e acentera.exe ] && rm -f acentera.exe

GITCOMMIT=$(git rev-parse HEAD)
git tag v${VERSION}
git push
git push --tags


mkdir -p out/
mv *.tar.gz out/.
mv *.zip out/.
 
# github-release release --security-token ${GITHUB_TOKEN} --user ACenterA --repo acenteracli --tag v${VERSION}
# 
# github-release upload --security-token ${GITHUB_TOKEN} --user ACenterA --repo acenteracli \
#     --tag v${VERSION} --name acentera-linux-amd64.tar.gz --file acentera-linux-amd64.tar.gz
# 
# github-release upload --security-token ${GITHUB_TOKEN} --user ACenterA --repo acenteracli \
#     --tag v${VERSION} --name acentera-darwin-amd64.tar.gz --file acentera-darwin-amd64.tar.gz
# 
# github-release upload --security-token ${GITHUB_TOKEN} --user ACenterA --repo acenteracli \
#     --tag v${VERSION} --name acentera-windows-amd64.zip --file acentera-windows-amd64.zip
# 
# github-release upload --security-token ${GITHUB_TOKEN} --user ACenterA --repo acenteracli \
#     --tag v${VERSION} --name acentera-linux-arm64.tar.gz --file acentera-linux-arm64.tar.gz
# github-release upload --security-token ${GITHUB_TOKEN} --user ACenterA --repo acenteracli \
# 
ghr -t "${GITHUB_TOKEN}" -u ACenterA -r acenteracli -c "${GITCOMMIT}" -n v${VERSION} -delete -replace  -generatenotes  v${VERSION} out/
