#!/bin/bash
#
# 快速部署脚本
#

go test -v -coverprofile=./test/testReport.out
go tool cover -html=./test/testReport.out -o=./test/testReport.html

VERSION=$1

git add .
git commit -m "fix:上传新版本"
git tag $VERSION
git push origin master
git push origin $VERSION