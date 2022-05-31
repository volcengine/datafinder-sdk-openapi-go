#!/bin/sh

rm -rf release/gosdk*
mkdir release/gosdk

cp -rf LICENSE release/gosdk/
cp -rf README.md release/gosdk/
cp -rf consts release/gosdk/
cp -rf dslcontent release/gosdk/
cp -rf go.mod release/gosdk/
cp -rf go.sum release/gosdk/
cp -rf util release/gosdk/

cd release
zip -r gosdk.zip gosdk/*

rm -rf gosdk

cd ../
