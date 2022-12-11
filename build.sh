#! /bin/bash
echo '创建upload_temp临时目录'
mkdir ./upload_temp
echo '将upload文件夹内容移动到upload_temp临时目录'
mv ./resource/public/upload/* ./upload_temp/
echo '开始打包'
gf build
echo '打包完成'
echo '将upload文件夹放回原处'
mv ./upload_temp/* ./resource/public/upload/
echo '删除upload_temp临时目录'
rm -rf ./upload_temp/
echo '全部完成'