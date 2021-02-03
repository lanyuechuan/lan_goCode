# lan_goCode
Submitting go code every day is a gesture,only persistence.

一、VsCode Go的正确安装方式
1、先从扩展管理中安装Go插件

2、安装Go插件所依赖的go tools（记得先go env -w GO111MODULE=on  和 go env -w GOPROXY=goproxy.cn)
按ctrl+shift+p 调出命令面板，输入go install tools 选Go: Install/Update Tools

点击全选下载

然后再进入包中go mod init ,至于第三方包，在代码中正常导入，导入格式是项目地址然后一定要在命令行运行main文件
二、VsCode Go的文件或项目打断点
只需要鼠标点击进入main文件（无论是单文件还是包文件总有入口文件），然后F5，即可开始打断点
