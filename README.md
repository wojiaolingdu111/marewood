<div align=center>
<img src="http://xusenlin.com/assets/images/marewood_logo.png" width=283.5" height="221" />
</div>
<div align=center>
<img src="https://img.shields.io/badge/golang-1.12-blue"/>
<img src="https://img.shields.io/badge/gin-1.4.0-lightBlue"/>
<img src="https://img.shields.io/badge/jwt-go-3.2.0-lightBlue"/>
<img src="https://img.shields.io/badge/gorm-1.9.11-red"/>
<img src="https://img.shields.io/badge/react-16.9.0-brightgreen"/>
<img src="https://img.shields.io/badge/material-ui-4.4.0-green"/>
</div>



# MareWood
>MareWood 是一个轻量级的前端部署工具，使用了 GOLANG、GIN、GORM、JWT、SQLITE、REACE、MATERIAL-UI 开发，
不同于 Jenkins 的大而全，它很简单且只针对前端，你可以很灵活的配置各种部署环境。
如果你愿意，线上发布也可以是点击一下按钮这么简单的事情,当然也可以配置 WEBHOOK，提交 GIT 代码既自动发布。


# MareWood 在做什么
MareWood 可以帮你克隆前端仓库并使用你选择的工具安装依赖，并能在同一个仓库下通过不同的任务不同的打包命令打包出多个环境下的前端资源并放在 resources\webs 下，通过任务 ID 一一对应，并提供快捷访问。



# 使用

- 服务器请先安装 git、node、npm，建议同时安装 cnpm 和 yarn 以供选择。
- 修改 config.json 来配置 MareWood 启动的端口号、域名、和支持的前端工具。
- 编译 MareWood 并放入后台运行。
- 客户端在 public 目录，请安装依赖并修改正式请求地址然后运行 npm run build 即可。（请求地址=> /public/src/config/url.js）

其他：MareWood 提供 http 访问，建议 nginx 提供 https 静态文件服务器（指向resources\webs）这样可以灵活的选择 http 或者 https 去访问打包的前端项目。


那么，一个 React 或者 Vue 项目如何配置多个打包命令呢？
请移步http://xusenlin.com/article?path=%2f%e5%89%8d%e7%ab%af%e5%bc%80%e5%8f%91%2fREACT%e5%92%8cVUE%e7%9a%84%e5%a4%9a%e5%a5%97%e6%89%93%e5%8c%85%e7%8e%af%e5%a2%83.md

如果不明白也可以参考实际的项目

- https://github.com/xusenlin/VueMultiplePages.git
- https://github.com/xusenlin/VueElementUiAdmin.git

# 角色

角色目前分为超级管理员、管理员、开发者、项目记者。权限如下：

- 项目记者 - 只能查看所有内容，并且可以访问打包好的前端项目
- 开发者 - 拥有创建仓库、分类、任务和拉代码、切换分支、运行打包、删除依赖权限
- 管理员 - 可以删除仓库、分类、任务
- 超级管理员 -可以管理用户（提升和降级角色，删除用户）

注：等级高的角色拥有低级角色的全部权限。用户注册默认成为项目记者，如果注册的名字是 Admin 会自动成为超级管理员，注册名字不可重复。

# 更新日志
### V0.2
* 添加用户和权限控制模块
* 可以快速修改仓库、分类、任务描述
* 优化和重构了大部分前后端代码
* 优化了前端界面UI
* 添加了 docker 容器支持
* 修复了一部分代码的隐藏 BUG

### V0.1
* 提供了服务器克隆仓库选择分支打包访问的核心功能

# docker镜像
感谢 @shynome 提供docker镜像

构建镜像步骤
```
# 固化依赖
go mod vendor
# 打包镜像
docker build . -t marewood
# 临时运行容器, 退出后数据不保存
docker run --rm -ti -p 8081:8081 marewood
# 持久化运行, 有容器卷挂载数据可以保存下来
docker run -d --restart always -p 8081:8081 -v marewood:/marewood/data marewood
```

示例用法
```
docker run --rm -ti \
  -p 8090 \
  -e HttpPort=8090 \
  -e AddressUrl=http://127.0.0.1:8090 \
  -v $PWD/marewood-data:/marewood/data \
   shynome/marewood
```

# 预览

### Repository
![Repository](./preview/仓库.png)
### Category
![Category](./preview/分类.png)
### Job
![Job](./preview/任务.png)


# 番外

MareWood 翻译过来是 玛瑞·伍德 看上去像是一个人的名字，其实它是代表两个人，一个咒语。



# License

MIT © Richard McRichface
