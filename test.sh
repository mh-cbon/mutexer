#!/bin/sh

set -ex

rm `which mutexer`
go install


mutexer - demo/Tomate:TomateSync | grep "Name(it" || exit 1;
mutexer - demo/Tomate:TomateSync | grep "package main" || exit 1;
mutexer -p nop - demo/Tomate:TomateSync | grep "package nop" || exit 1;

rm -fr gen_test
mutexer demo/Tomate:gen_test/TomateSync || exit 1;
ls -al gen_test | grep "tomatesync.go" || exit 1;
cat gen_test/tomatesync.go | grep "Name(it" || exit 1;
cat gen_test/tomatesync.go | grep "package gen_test" || exit 1;
rm -fr gen_test

rm -fr demo/tomates.go
rm -fr demo/tomatessync.go
go generate demo/main.go
ls -al demo | grep "tomates.go" || exit 1;
ls -al demo | grep "tomatessync.go" || exit 1;
cat demo/tomatessync.go | grep "package main" || exit 1;
cat demo/tomatessync.go | grep "NewTomatesSync(" || exit 1;
go run demo/*.go | grep "Hello world!" || exit 1;
# rm -fr demo/tomates.go # keep it for demo

rm -fr demo/gen
go generate github.com/mh-cbon/mutexer/demo
ls -al demo | grep "tomatessync.go" || exit 1;
cat demo/tomatessync.go | grep "package main" || exit 1;
go run demo/*.go | grep "Hello world!" || exit 1;
# rm -fr demo/gen # keep it for demo

# go test


echo ""
echo "ALL GOOD!"
