
English | [简体中文](./README.md) 

<div align=center>
<img src="https://user-images.githubusercontent.com/78396698/203823434-c90a37a2-5711-45bb-9936-1d7cb7a29b4e.png" width="430" height="380" />
</div>
<div align=center>
<img src="https://img.shields.io/badge/golang-1.16-blue"/>
<img src="https://img.shields.io/badge/hertz-0.4.0-lightBlue"/>
<img src="https://img.shields.io/badge/vue-3.2.25-brightgreen"/>
<img src="https://img.shields.io/badge/element--plus-2.0.1-green"/>
<img src="https://img.shields.io/badge/gorm-1.22.5-red"/>
</div>		


START--

**First of thanks:** At present, the source code of the business part of Hertz vue admin open source project refers to the open source project [Gin Vue Admin](https://www.gin-vue-admin.com) And [Eladmin](https://eladmin.vip/ )。

**Important statement:** Since Hertz vue admin is being restructured, it is only open source for non-commercial use before the completion of the reconstruction. If you need to use it commercially, please move to the open source project [Gin Vue Admin]() According to its [flipped aurora team](https://github.com/flipped-aurora ) Declaration. (At the same time, Hertz vue admin team decided to provide some open source commercial products after the reconstruction of Hertz vue admin is completed.)

## 1. Basic introduction

### 1.1 Project introduction

Hertz vue admin is a new Golang microservice HTTP framework based on Vue3 and bytedance team [Hertz](https://www.cloudwego.io/zh/docs/hertz/ ) It is a development basic platform that separates the front end from the back end of the whole stack, integrating functions such as jwt authentication, dynamic routing, dynamic menu, casbin authentication, form generator, code generator, etc.

### 1.2 Online experience

Online address：http://1.117.77.53:8080/index.html

username：admin

password：123456



## 2. Project operation

### 2.1 Environment configuration

```
- node version > v16.8.3
- golang version >= v1.16
```

### 2.2 server

Try to put Hertz-vue-admin in the GOPATH directory

```
# Enter the server folder
cd server

# Use go mod and install the go dependency package
go generate

# compile
go build -o server main.go (windows compilego: go build -o server.exe main.go )

# Run
./server (windows运行命令为 server.exe)
```

### 2 .3、web

```
# Enter the web folder
cd web

# Installation Dependencies
npm install

# Run
npm run serve
```

### 2.3 swagger Automation API Documentation

```
# Installation Dependencies
go get -u github.com/swaggo/swag/cmd/swag

# Enter the server folder
cd server

# Run
swag init
```

After executing the above command, ` docs go`, `swagger. json`, `swagger. Yaml ` The three files are updated. After starting the go service, enter them in the browser http://localhost:8888/swagger/index.html You can view the swagger document



## 3. Technical selection

- Web: use [Vue](https://vuejs.org/ ) based  [Element](https://github.com/ElemeFE/element) Build the base page.

- Server：use ByteDance Cloudwego's Golang microservice HTTP framework  — [Hertz](https://www.cloudwego.io/zh/docs/hertz/ ).

- Database: adopt the MySql >(5.7) version database engine InnoDB, and use [gorm]( http://gorm.cn/) Implement the basic operation of the database.

- Caching: Use Redis to record the jwt token of the current active user and implement the restrictions on multiple logins.

- API document: use Swagger to build an automatic document.。

- Configuration file: using [fsnotify](https://github.com/fsnotify/fsnotify )  And [viper](https://github.com/spf13/viper) Implement the configuration file in the format of yaml.。

- Log: using [zap]( https://github.com/uber-go/zap) Implement logging.

  

## 4. Project architecture

### 4.1 web architecture

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



### 4.2 server architecture

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
