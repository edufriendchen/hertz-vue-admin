[English](./README-en.md) | 简体中文

<div align=center>
<img src="https://user-images.githubusercontent.com/78396698/203823434-c90a37a2-5711-45bb-9936-1d7cb7a29b4e.png" width="430" height="380" />
</div>
<div align=center>
<img src="https://img.shields.io/badge/golang-1.16-blue"/>
<img src="https://img.shields.io/badge/gin-1.7.0-lightBlue"/>
<img src="https://img.shields.io/badge/vue-3.2.25-brightgreen"/>
<img src="https://img.shields.io/badge/element--plus-2.0.1-green"/>
<img src="https://img.shields.io/badge/gorm-1.22.5-red"/>
</div>		


START--

**首先鸣谢:** 目前Hertz-vue-admin开源项目业务部分源码借鉴开源项目[Gin-Vue-Admin](https://www.gin-vue-admin.com)和[Eladmin](https://eladmin.vip/)。

**重要声明:** 由于Hertz-vue-admin正在重构，在重构完成之前只对非商用开源，如需商用请移步开源项目[Gin-Vue-Admin](https://www.gin-vue-admin.com)，依据其[flipped-aurora团队](https://github.com/flipped-aurora)声明。 （同时Hertz-vue-admin团队我们决定重构完成后，开源项目Hertz-vue-admin重构完成后将提供部分开源商用）



## 1. 基本介绍

### 1.1 项目介绍

Hertz-vue-admin是一个基于Vue3和字节跳动团队最新Golang微服务HTTP框架[Hertz](https://www.cloudwego.io/zh/docs/hertz/)开发的全栈前后端分离的开发基础平台，集成jwt鉴权，动态路由， 动态菜单，casbin鉴权，表单生成器，代码生成器等功能。

### 1.2 在线体验

在线地址：http://1.117.77.53:8080/index.html

测试用户名：admin

测试密码：123456



## 2.项目运行

### 2.1 环境配置

```
- node版本 > v16.8.3
- golang版本 >= v1.16
```

### 2.2 server项目

尽量将Hertz-Vue-Admin放在GOPATH目录下

```
# 进入server文件夹
cd server

# 使用 go mod 并安装go依赖包
go generate

# 编译 
go build -o server main.go (windows编译命令为go build -o server.exe main.go )

# 运行二进制
./server (windows运行命令为 server.exe)
```

### 2 .3、web项目

```
# 进入web文件夹
cd web

# 安装依赖
npm install

# 启动web项目
npm run serve
```

### 2.3 swagger自动化API文档

```
#导入依赖
go get -u github.com/swaggo/swag/cmd/swag

# 进入server文件夹
cd server

#启动
swag init
```

执行上面的命令后，server目录下会出现docs文件夹里的 `docs.go`, `swagger.json`, `swagger.yaml` 三个文件更新，启动go服务之后, 在浏览器输入 http://localhost:8888/swagger/index.html 即可查看swagger文档



## 3.技术选型

- 前端：使用基于 [Vue](https://vuejs.org/) 的 [Element](https://github.com/ElemeFE/element) 构建基础页面。

- 后端：使用字节跳动Cloudwego的Golang微服务HTTP框架。

- 数据库：采用`MySql` > (5.7) 版本 数据库引擎 InnoDB，使用 [gorm](http://gorm.cn/) 实现对数据库的基本操作。

- 缓存：使用`Redis`实现记录当前活跃用户的`jwt`令牌并实现多点登录限制。

- API文档：使用`Swagger`构建自动化文档。

- 配置文件：使用 [fsnotify](https://github.com/fsnotify/fsnotify) 和 [viper](https://github.com/spf13/viper) 实现`yaml`格式的配置文件。

- 日志：使用 [zap](https://github.com/uber-go/zap) 实现日志记录。

  

## 4. 项目架构

### 4.1 前端结构

```
web
 ├── babel.config.js
 ├── Dockerfile
 ├── favicon.ico
 ├── index.html                  -- 主页面
 ├── limit.js                    -- 助手代码
 ├── package.json                -- 包管理器代码
 ├── src                         -- 源代码
 │   ├── api                    -- api 组
 │   ├── App.vue                -- 主页面
 │   ├── assets                 -- 静态资源
 │   ├── components             -- 全局组件
 │   ├── core                   -- 组件包
 │   │   ├── config.js         -- 网站配置文件
 │   │   ├── hertz-vue-admin.js  -- 注册欢迎文件
 │   │   └── global.js         -- 统一导入文件
 │   ├── directive              -- v-auth 注册文件
 │   ├── main.js                -- 主文件
 │   ├── permission.js          -- 路由中间件
 │   ├── pinia                  -- pinia 状态管理器，取代vuex
 │   │   ├── index.js          -- 入口文件
 │   │   └── modules           -- modules
 │   │       ├── dictionary.js
 │   │       ├── router.js
 │   │       └── user.js
 │   ├── router                 -- 路由声明文件
 │   │   └── index.js
 │   ├── style                  -- 全局样式
 │   │   ├── base.scss
 │   │   ├── basics.scss
 │   │   ├── element_visiable.scss  -- 此处可以全局覆盖 element-plus 样式
 │   │   ├── iconfont.css           -- 顶部几个icon的样式文件
 │   │   ├── main.scss
 │   │   ├── mobile.scss
 │   │   └── newLogin.scss
 │   ├── utils                  -- 方法包库
 │   │   ├── asyncRouter.js    -- 动态路由相关
 │   │   ├── btnAuth.js        -- 动态权限按钮相关
 │   │   ├── bus.js            -- 全局mitt声明文件
 │   │   ├── date.js           -- 日期相关
 │   │   ├── dictionary.js     -- 获取字典方法 
 │   │   ├── downloadImg.js    -- 下载图片方法
 │   │   ├── format.js         -- 格式整理相关
 │   │   ├── image.js          -- 图片相关方法
 │   │   ├── page.js           -- 设置页面标题
 │   │   ├── request.js        -- 统一请求文件
 │   │   └── stringFun.js      -- 字符串文件
 |   ├── view                   -- 主要view代码
 |   |   ├── about              -- 关于我们
 |   |   ├── dashboard          -- 面板
 |   |   ├── error              -- 错误
 |   |   ├── example            -- 上传案例
 |   |   ├── iconList           -- icon列表
 |   |   ├── init               -- 初始化数据  
 |   |   ├── layout             -- layout约束页面 
 |   |   |   ├── aside          -- 侧边栏
 |   |   |   ├── bottomInfo     -- bottomInfo
 |   |   |   ├── screenfull     -- 全屏设置
 |   |   |   ├── setting        -- 系统设置
 |   |   |   └── index.vue      -- base 约束
 |   |   ├── login              --登录 
 |   |   ├── person             --个人中心 
 |   |   ├── superAdmin         -- 超级管理员操作
 |   |   ├── system             -- 系统检测页面
 |   |   ├── systemTools        -- 系统配置相关页面
 |   |   └── routerHolder.vue   -- page 入口页面 
 ├── vite.config.js             -- vite 配置文件
 └── yarn.lock
```



### 4.2 后端结构

```
├── server
        ├── api             (api层)
        │   └── v1          (v1版本接口)
        ├── config          (配置包)
        ├── core            (核心文件)
        ├── docs            (swagger文档目录)
        ├── global          (全局对象)                    
        ├── initialize      (初始化)                        
        │   └── internal    (初始化内部函数)                            
        ├── middleware      (中间件层)                        
        ├── model           (模型层)                    
        │   ├── request     (入参结构体)                        
        │   └── response    (出参结构体)                            
        ├── packfile        (静态文件打包)                        
        ├── resource        (静态资源文件夹)                        
        │   ├── excel       (excel导入导出默认路径)                        
        │   ├── page        (表单生成器)                        
        │   └── template    (模板)                            
        ├── router          (路由层)                    
        ├── service         (service层)                    
        ├── source          (source层)                    
        └── utils           (工具包)                    
            ├── timer       (定时器接口封装)                        
            └── upload      (oss接口封装)                        
    
            web
        ├── babel.config.js
        ├── Dockerfile
        ├── favicon.ico
        ├── index.html                 -- 主页面
        ├── limit.js                   -- 助手代码
        ├── package.json               -- 包管理器代码
        ├── src                        -- 源代码
        │   ├── api                    -- api 组
        │   ├── App.vue                -- 主页面
        │   ├── assets                 -- 静态资源
        │   ├── components             -- 全局组件
        │   ├── core                   -- gva 组件包
        │   │   ├── config.js          -- gva网站配置文件
        │   │   ├── gin-vue-admin.js   -- 注册欢迎文件
        │   │   └── global.js          -- 统一导入文件
        │   ├── directive              -- v-auth 注册文件
        │   ├── main.js                -- 主文件
        │   ├── permission.js          -- 路由中间件
        │   ├── pinia                  -- pinia 状态管理器，取代vuex
        │   │   ├── index.js           -- 入口文件
        │   │   └── modules            -- modules
        │   │       ├── dictionary.js
        │   │       ├── router.js
        │   │       └── user.js
        │   ├── router                 -- 路由声明文件
        │   │   └── index.js
        │   ├── style                  -- 全局样式
        │   │   ├── base.scss
        │   │   ├── basics.scss
        │   │   ├── element_visiable.scss  -- 此处可以全局覆盖 element-plus 样式
        │   │   ├── iconfont.css           -- 顶部几个icon的样式文件
        │   │   ├── main.scss
        │   │   ├── mobile.scss
        │   │   └── newLogin.scss
        │   ├── utils                  -- 方法包库
        │   │   ├── asyncRouter.js     -- 动态路由相关
        │   │   ├── btnAuth.js         -- 动态权限按钮相关
        │   │   ├── bus.js             -- 全局mitt声明文件
        │   │   ├── date.js            -- 日期相关
        │   │   ├── dictionary.js      -- 获取字典方法 
        │   │   ├── downloadImg.js     -- 下载图片方法
        │   │   ├── format.js          -- 格式整理相关
        │   │   ├── image.js           -- 图片相关方法
        │   │   ├── page.js            -- 设置页面标题
        │   │   ├── request.js         -- 请求
        │   │   └── stringFun.js       -- 字符串文件
        |   ├── view -- 主要view代码
        |   |   ├── about -- 关于我们
        |   |   ├── dashboard -- 面板
        |   |   ├── error -- 错误
        |   |   ├── example --上传案例
        |   |   ├── iconList -- icon列表
        |   |   ├── init -- 初始化数据  
        |   |   |   ├── index -- 新版本
        |   |   |   ├── init -- 旧版本
        |   |   ├── layout  --  layout约束页面 
        |   |   |   ├── aside 
        |   |   |   ├── bottomInfo     -- bottomInfo
        |   |   |   ├── screenfull     -- 全屏设置
        |   |   |   ├── setting        -- 系统设置
        |   |   |   └── index.vue      -- base 约束
        |   |   ├── login              --登录 
        |   |   ├── person             --个人中心 
        |   |   ├── superAdmin         -- 超级管理员操作
        |   |   ├── system             -- 系统检测页面
        |   |   ├── systemTools        -- 系统配置相关页面
        |   |   └── routerHolder.vue   -- page 入口页面 
        ├── vite.config.js             -- vite 配置文件
        └── yarn.lock
```
