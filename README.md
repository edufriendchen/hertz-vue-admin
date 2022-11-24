
[English](./README-en.md) | 简体中文

<div align=center>
<img src="https://user-images.githubusercontent.com/78396698/203819698-18a1b65d-02e1-4524-bd7c-a6306a00f570.png" width="400" height="380" />
</div>
<div align=center>
<img src="https://img.shields.io/badge/golang-1.16-blue"/>
<img src="https://img.shields.io/badge/gin-1.7.0-lightBlue"/>
<img src="https://img.shields.io/badge/vue-3.2.25-brightgreen"/>
<img src="https://img.shields.io/badge/element--plus-2.0.1-green"/>
<img src="https://img.shields.io/badge/gorm-1.22.5-red"/>
</div>




Hertz-Vue-Admin环境说明

```
- node版本 > v16.8.3
- golang版本 >= v1.16
- IDE推荐：Goland
```

### 1、 server项目

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



### 2 、web项目

```
# 进入web文件夹
cd web

# 安装依赖
npm install

# 启动web项目
npm run serve
```



测试用户名：admin

测试密码：123456
