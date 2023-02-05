
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

**Important statement:** The new Hertz-vue-admin is under reconstruction. Welcome to use the Friendly Open Source Project [Gin Vue Admin]() , according to the statement of its [flipped aurora team](https://github.com/flipped-aurora ).

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
 ├── index.html                  
 ├── limit.js
 ├── package.json              
 ├── src                         
 │   ├── api              
 │   ├── App.vue
 │   ├── assets               
 │   ├── components
 │   ├── core                   
 │   │   ├── config.js       
 │   │   ├── hertz-vue-admin.js  
 │   │   └── global.js         
 │   ├── directive             
 │   ├── main.js                
 │   ├── permission.js          
 │   ├── pinia                  
 │   │   ├── index.js          
 │   │   └── modules         
 │   │       ├── dictionary.js
 │   │       ├── router.js
 │   │       └── user.js
 │   ├── router                 
 │   │   └── index.js
 │   ├── style                 
 │   │   ├── base.scss
 │   │   ├── basics.scss
 │   │   ├── element_visiable.scss  
 │   │   ├── iconfont.css         
 │   │   ├── main.scss
 │   │   ├── mobile.scss
 │   │   └── newLogin.scss
 │   ├── utils                 
 │   │   ├── asyncRouter.js    
 │   │   ├── btnAuth.js      
 │   │   ├── bus.js           
 │   │   ├── date.js          
 │   │   ├── dictionary.js     
 │   │   ├── downloadImg.js  
 │   │   ├── format.js         
 │   │   ├── image.js          
 │   │   ├── page.js           
 │   │   ├── request.js        
 │   │   └── stringFun.js      
 |   ├── view                  
 |   |   ├── about              
 |   |   ├── dashboard         
 |   |   ├── error              
 |   |   ├── example           
 |   |   ├── iconList          
 |   |   ├── init              
 |   |   ├── layout             
 |   |   |   ├── aside          
 |   |   |   ├── bottomInfo     
 |   |   |   ├── screenfull     
 |   |   |   ├── setting        
 |   |   |   └── index.vue      
 |   |   ├── login             
 |   |   ├── person             
 |   |   ├── superAdmin         
 |   |   ├── system             
 |   |   ├── systemTools        
 |   |   └── routerHolder.vue   
 ├── vite.config.js             
 └── yarn.lock


```



### 4.2 server architecture

```
├── server
        ├── api            
        │   └── v1          
        ├── config         
        ├── core            
        ├── docs            
        ├── global                            
        ├── initialize                          
        │   └── internal                             
        ├── middleware                            
        ├── model                             
        │   ├── request                            
        │   └── response                               
        ├── packfile                            
        ├── resource                               
        │   ├── excel                            
        │   ├── page                             
        │   └── template                               
        ├── router                        
        ├── service                          
        ├── source                             
        └── utils                            
            ├── timer                              
            └── upload                         
    
            web
        ├── babel.config.js
        ├── Dockerfile
        ├── favicon.ico
        ├── index.html               
        ├── limit.js                  
        ├── package.json            
        ├── src                       
        │   ├── api                    
        │   ├── App.vue              
        │   ├── assets               
        │   ├── components            
        │   ├── core                   
        │   │   ├── config.js        
        │   │   ├── gin-vue-admin.js  
        │   │   └── global.js         
        │   ├── directive              
        │   ├── main.js              
        │   ├── permission.js          
        │   ├── pinia                 
        │   │   ├── index.js           
        │   │   └── modules            
        │   │       ├── dictionary.js
        │   │       ├── router.js
        │   │       └── user.js
        │   ├── router                
        │   │   └── index.js
        │   ├── style                 
        │   │   ├── base.scss
        │   │   ├── basics.scss
        │   │   ├── element_visiable.scss  
        │   │   ├── iconfont.css           
        │   │   ├── main.scss
        │   │   ├── mobile.scss
        │   │   └── newLogin.scss
        │   ├── utils                  
        │   │   ├── asyncRouter.js     
        │   │   ├── btnAuth.js         
        │   │   ├── bus.js            
        │   │   ├── date.js            
        │   │   ├── dictionary.js      
        │   │   ├── downloadImg.js     
        │   │   ├── format.js         
        │   │   ├── image.js         
        │   │   ├── page.js            
        │   │   ├── request.js         
        │   │   └── stringFun.js       
        |   ├── view 
        |   |   ├── about 
        |   |   ├── dashboard 
        |   |   ├── error 
        |   |   ├── example 
        |   |   ├── iconList 
        |   |   ├── init  
        |   |   |   ├── index 
        |   |   |   ├── init 
        |   |   ├── layout  
        |   |   |   ├── aside 
        |   |   |   ├── bottomInfo     
        |   |   |   ├── screenfull   
        |   |   |   ├── setting        
        |   |   |   └── index.vue      
        |   |   ├── login            
        |   |   ├── person             
        |   |   ├── superAdmin         
        |   |   ├── system             
        |   |   ├── systemTools       
        |   |   └── routerHolder.vue   
        ├── vite.config.js            
        └── yarn.lock
```
