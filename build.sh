#!/bin/bash
#
# 快速部署脚本
#

go test -v -coverprofile=./test/testReport.out
go tool cover -html=./test/testReport.out -o=./test/testReport.html

VERSION=$1
echo $VERSION
if [! -n "$VERSION"];then
    echo "请输入tab版本号";
    exit 1;
else
    echo "tag版本号为$VERSION"
fi;

git add .
git commit -m "fix:上传新版本"
git tag $VERSION
git push origin master
git push origin $VERSION