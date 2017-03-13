#!/bin/bash

PWD=$(pwd)
GO_EXT="lukehoban.Go"
SETTINGS_DIR=".vscode"
SETTINGS_FILE="$SETTINGS_DIR/settings.json"
BIN_DIR="bin"

install_tool() {
	if [ ! -f $BIN_DIR/$1 ]; then
		echo "$1 not installed...installing..."
		go get -u -v $2
	else
		echo "$1 already installed..."
	fi
}

DID_INSTALL_EXT=0
INSTALLED_EXTS=`code --list-extensions`
if [[ ! $INSTALLED_EXTS  =~ .*$GO_EXT*. ]]; then
	echo "$GO_EXT not installed...installing..."
	code --install-extension $GO_EXT
	DID_INSTALL_EXT=1
else 
	echo "$GO_EXT already installed..."
fi

echo -e "\nConfiguring extension settings..."
if [ ! -d $SETTINGS_DIR ]; then
	mkdir $SETTINGS_DIR
fi

if [ ! -f $SETTINGS_FILE ]; then
	echo "{" >> $SETTINGS_FILE
	echo "	\"go.formatTool\": \"goreturns\"," >> $SETTINGS_FILE
	echo "	\"go.docsTool\": \"gogetdoc\"," >> $SETTINGS_FILE
	echo "	\"go.lintOnSave\": true," >> $SETTINGS_FILE
	echo "	\"go.gopath\": \"$PWD\"" >> $SETTINGS_FILE
	echo "}" >> $SETTINGS_FILE
fi

echo -e "\nInstalling tools for extension..."
install_tool "gocode"     "github.com/nsf/gocode"
install_tool "godef"      "github.com/rogpeppe/godef"
install_tool "gogetdoc"   "github.com/zmb3/gogetdoc"
install_tool "golint"     "github.com/golang/lint/golint"
install_tool "go-outline" "github.com/lukehoban/go-outline"
install_tool "goreturns"  "sourcegraph.com/sqs/goreturns"
install_tool "gorename"   "golang.org/x/tools/cmd/gorename"
install_tool "gopkgs"     "github.com/tpng/gopkgs"
install_tool "go-symbols" "github.com/newhook/go-symbols"
install_tool "guru"       "golang.org/x/tools/cmd/guru"
install_tool "gotests"    "github.com/cweill/gotests/..."

echo -e "\nEnvironment setup complete"
if [ "$DID_INSTALL_EXT" = 1 ]; then
	echo "Restart Visual Studio Code to activate $GO_EXT"
fi
