# 基于GoFrameV2 的网站内容管理系统
> 前台演示地址：[demo-pc.linklink.fun](http://demo-pc.linklink.fun) （自动跳转移动端）
> 
> 移动演示地址：[demo-m.linklink.fun](http://demo-m.linklink.fun)
> 
> 后台演示地址：[demo-pc.linklink.fun/admin](http://demo-pc.linklink.fun/admin)
>
> 
> 用户名：<code>admin</code>
>
> 密码：<code>123456</code>

> 📱 新增部署在arm架构手机内的演示地址：
> 
> 前台：[mi8-gfcms.linklink.fun](http://mi8-gfcms.linklink.fun)
> 
> 后台：[mi8-gfcms.linklink.fun/admin](http://mi8-gfcms.linklink.fun/admin)
>
>
> 用户名：<code>admin</code>
>
> 密码：<code>123456</code>
>
> 部署环境为：小米8手机 -> LinuxDeploy容器 -> CentOS7
>
> 
> 手机不可能随时开机并运行项目，无法访问请提交issue联系我

> 安装步骤
> 
> 1、导入gfcms_demo.sql到数据库
>
> 2、复制配置文件到项目根目录，生产环境部署时将该配置文件放置到可执行文件同目录下
>
> ```shell
> cp manifest/config/config.yaml.example config.yaml
> ```
> 
> 3、修改config.yaml的数据库配置、redis配置
>
> 4、修改config.yaml域名绑定，mobileHost为移动端绑定域名，pcHost为PC端绑定域名，
> mobileHost留空不自动跳转移动端，mobileHost和pcHost都留空则使用IP加端口号访问
> 
> 5、下载演示站图片资源，将解压后的uploads文件夹放到项目resource/public文件夹下（不在乎无法显示图片的可以不下载）
> 
> 下载地址：http://pan.linklink.fun/#s/932wYqlg 访问密码：<code>C2Gq7</code>
> 
> 6、编译项目
> 
> 安装 [gf-cli] 工具（推荐、开发时必备）
> 
> ```shell
> wget -O gf https://github.com/gogf/gf/releases/latest/download/gf_$(go env GOOS)_$(go env GOARCH) && chmod +x gf && ./gf install -y && rm ./gf

> 已安装 [gf-cli] 工具
> ```shell
> gf run main.go
> ```
>
> 未安装工具
> ```shell
> go build -o ./main main.go
> ```
> 
> 7、启动项目
> 
> ```shell
> ./main
> ```

> 目前支持的模型：
> - 文章
> - 图集

> GitHub地址：https://github.com/demozx/gf_cms
>
> Gitee地址：https://gitee.com/demozx/gf_cms

> 截图
- 后台
![image](./README/backend.jpg)
![image](./README/channel.png)
- PC
![image](./README/pc.jpg)

- 移动
![image](./README/mobile.jpg)


## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=demozx/gf_cms&type=Date)](https://www.star-history.com/#demozx/gf_cms&Date)

