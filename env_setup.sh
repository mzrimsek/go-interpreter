#!/bin/bash
PWD=$(pwd)
SETTINGS_DIR=".vscode"
SETTINGS_FILE="$SETTINGS_DIR/settings.json"
BIN_DIR="bin"

echo "Setting up settings.json..."
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
echo "Done"

echo "Installing tools..."
if [ ! -f $BIN_DIR/gocode ]; then
	go get -u -v github.com/nsf/gocode
fi

if [ ! -f $BIN_DIR/godef ]; then
	go get -u -v github.com/rogpeppe/godef
fi

if [ ! -f $BIN_DIR/gogetdoc ]; then
	go get -u -v github.com/zmb3/gogetdoc
fi

if [ ! -f $BIN_DIR/golint ]; then
	go get -u -v github.com/golang/lint/golint
fi

if [ ! -f $BIN_DIR/go-outline ]; then
	go get -u -v github.com/lukehoban/go-outline
fi

if [ ! -f $BIN_DIR/goreturns ]; then
	go get -u -v sourcegraph.com/sqs/goreturns
fi

if [ ! -f $BIN_DIR/gorename ]; then
	go get -u -v golang.org/x/tools/cmd/gorename
fi

if [ ! -f $BIN_DIR/gopkgs ]; then
	go get -u -v github.com/tpng/gopkgs
fi

if [ ! -f $BIN_DIR/go-symbols ]; then
	go get -u -v github.com/newhook/go-symbols
fi

if [ ! -f $BIN_DIR/guru ]; then 
	go get -u -v golang.org/x/tools/cmd/guru
fi

if [ ! -f $BIN_DIR/gotests ]; then
	go get -u -v github.com/cweill/gotests/...
fi
echo "Done"
