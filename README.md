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
1. 运行```BackEnd/main.go``来启动后端服务
2. 运行vue

在frontend文件夹下面使用如下命令启动
 ```shell
npm run serve
```
## 4.部署项目
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

