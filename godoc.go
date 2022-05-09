/*
Detect programming language of source files.
Go port of GitHub Linguist: https://github.com/github/linguist

Prerequisites:

    go get github.com/jteeuwen/go-bindata/go-bindata

Installation:

    mkdir -p $GOPATH/src/github.com/ianlewis/linguist
    git clone --depth=1 https://github.com/ianlewis/linguist $GOPATH/src/github.com/ianlewis/linguist
    go get -d github.com/ianlewis/linguist
    cd $GOPATH/src/github.com/ianlewis/linguist
    make
    l

Usage:

Please refer to the source code for the reference implementation at:

https://github.com/ianlewis/linguist/tree/master/cmd/l


See also:

https://github.com/ianlewis/linguist/tree/master/tokenizer
*/
package linguist
