#!/bin/bash
set -e

# Download latest binary from Github

ARCH_UNAME=`uname -m`
if [[ "$ARCH_UNAME" == "x86_64" ]]; then
	ARCH="amd64"
else
	ARCH="386"
fi

EXT="tar.gz"

if [[ "$OSTYPE" == "linux-gnu" ]]; then
	OS="linux"
elif [[ "$OSTYPE" == "darwin"* ]]; then
	OS="darwin"
elif [[ "$OSTYPE" == "win32" ]] || [[ "$OSTYPE" == "cygwin" ]] || [[ "$OSTYPE" == "msys" ]] ; then
	OS="windows"
	EXT="zip"
else
	echo "No awless binary available for OS '$OSTYPE'. You may want to use go to install awless with 'go get -u github.com/wallix/awless'"
  exit
fi

LATEST_VERSION=`curl -s 'https://github.com/ACenterA/acenteracli/releases' | grep ".tar.gz" | grep 'href.*\/v' | head -n 1  | cut -d '"' -f 2 | sed -E 's~^.*download/v(.*)/.*tar.gz~\1~g'`

FILENAME=acentera-$OS-$ARCH.$EXT

DOWNLOAD_URL="https://github.com/ACenterA/acenteracli/releases/download/$LATEST_VERSION/$FILENAME"

echo "Downloading awless from $DOWNLOAD_URL"

if ! curl --fail -o $FILENAME -L $DOWNLOAD_URL; then
    exit
fi

echo ""
echo "extracting $FILENAME to ./acentera"

if [[ "$OS" == "windows" ]]; then
	echo 'y' | unzip $FILENAME 2>&1 > /dev/null
else
	tar -xzf $FILENAME
fi

echo "removing $FILENAME"
rm $FILENAME
chmod +x ./acentera

echo ""
echo "acentera successfully installed to ./acentera"
echo ""
echo "don't forget to add it to your path, with, for example, 'sudo mv awless /usr/local/bin/' "
echo ""
echo "then, for autocompletion, run:"
echo "    [bash] echo 'source <(acentera completion bash)' >> ~/.bashrc"
echo "    OR"
echo "    [zsh]  echo 'source <(acentera completion zsh)' >> ~/.zshrc"
