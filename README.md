# ustb-campus贝壳校园墙网页版
## 1.项目简介
本网站是基于微信小程序《贝壳校园墙》进行的改编，使用golang+vue进行了重构，**贝壳校园墙**小程序上能看到的信息本网页版都可以看到。

本网站可以作为golang和vue3的学习项目，使用前后端分离架构，后端基于gin框架。
## 2.项目开发发前的准备
### 2.1. 下载go语言的开发环境
### 2.2. 下载node.js
#### 2.2.1 配置环境
为了提高node.js的下载速度，建议切换国内镜像源，通常默认官方源是 ```https://registry.npmjs.org```，可以在cmd命令行下通过下面的命令查看
```shell
npm config get registry
```
在国内使用官方下载依赖往往速度慢，易出错，因此我们选择切换国内镜像源，根据需求或喜好选择下列命令中的一条
```shell
//切换淘宝源
npm config set registry https://registry.npm.taobao.org

//切换腾讯源
npm config set registry http://mirrors.cloud.tencent.com/npm/

//切换阿里云源
npm config set registry https://registry.npmmirror.com
```
#### 2.2.2 下载vue的构建工具
```shell
npm install -g vue-cli
```
如果下载失败，可以通过下面的命令删除后重新下载
```shell
npm uninstall vue-cli -g
```
#### 2.2.3 使用vue构建项目
```shell
vue init webpack FrontEnd
```
构建项目时需要选择相应的配置，我的选择如下
```text
? Project name frontend
? Project description 贝壳校园墙前端vue
? Author 我不是大佬zvj
? Vue build standalone
? Install vue-router? Yes
? Use ESLint to lint your code? No
? Set up unit tests No
? Setup e2e tests with Nightwatch? No
? Should we run `npm install` for you after the project has been created? (recommended) npm
```
#### 2.2.4 项目中使用到的依赖
```text
npm install axios --save
```
## 3.本地运行项目
### 3.1 下载代码
通过git将代码下载到本地

```shell
git clone https://github.com/U202142209/ustb-campus.git
```
### 3.2 启动go代码
1. 首先进入```BackEnd``文件夹
2. 运行以下命令来下载项目所需的所有依赖：
```shell
go mod download
```
这将根据 go.mod 文件中列出的依赖项下载并安装所有必需的包。如果您还希望将这些依赖项保存到 vendor 目录中，可以运行以下命令：   
```go mod vendor```
### 3.3 启动前端vue
在frontend文件夹下面使用如下命令启动
 ```shell
npm run serve
```
## 4.生产环境部署项目
### 4.1 安装go语言的开发环境
**下载go**
```shell
sudo aria2c  https://golang.google.cn/dl/go1.21.3.linux-amd64.tar.gz
```
**解压**
```shell
shdo tar -C /usr/local -xzf go1.21.3.linux-amd64.tar.gz
```
**测试是否安装了go**
```shell
go version
```
### 4.2 运行go后端代码
首先在```go.mod```所在的文件夹下执行下面的命令安装依赖
```text
go mod tidy
go mod download
```            
然后**编译代码** ```go build -o ustb-campus main.go```
这将在当前目录下生成一个名为 ustb-campus 的可执行文件

**封装systemctl服务**
```text
sudo vim /etc/systemd/system/ustbcampus.service
```
```text
[Unit]
Description=ustbcampus Service
After=network.target

[Service]
ExecStart=/src/ustb-campus/BackEnd/ustb-campus
WorkingDirectory=/src/ustb-campus/BackEnd
User=root
Group=root
Restart=always

[Install]
WantedBy=multi-user.target
```
**启动、停止服务**
```text
sudo systemctl enable ustbcampus 
sudo systemctl start  ustbcampus
sudo systemctl disable  ustbcampus 
```
**nginx配置**
```text
server {
    listen 80;
    server_name ustb.campus.work.wobushidalao.top;

    location / {
        proxy_pass http://localhost:8000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```
### 4.2 部署前端vue3
首先编译项目代码
```shell
npm run build
```
然后将编译生成的 ```dist```文件夹上传到服务器，dist文件夹里面包括```js```、```css```、```index.html```等信息，我们生产环境只需要这些就可以了，因为vue已经帮我们将所有的依赖全部集成到这里面了

我的nginx配置文件如下，大家只需要改一下配置的域名和dist文件夹的绝对路径就可以了
```text
server {
    listen 80;
    server_name ustb.campus.wobushidalao.top;

    location /proxy_url/ {
        proxy_pass http://ustb.campus.work.wobushidalao.top/;
    }

    location / {
        root /src/ustb-campus/frontend/dist;
        try_files $uri $uri/ /index.html;
    }
}
```
## 5.注意
## 6.常见问题
全局安装Vue CLI（用于创建和管理Vue项目）
```shell
npm install -g @vue/cli
```
创建一个新的Vue项目
```
vue create my-project
```
## 7. 获取帮助
```text
@Author ;@我不是大佬
@Email  ;2869210303@qq.com
@Qq     ;2869210303
@wx     ;safeseaa
@github ;https://github.com/U202142209
@blog   ;https://blog.csdn.net/V123456789987654
```
