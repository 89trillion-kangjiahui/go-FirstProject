# 整体框架
```text
功能：用fyne开发聊天客户端，连接第5题的聊天服务。
1）连接管理：连接、断开、展示连接状态（连接、断开）
2）展示用户列表
3）发送消息
4）展示收到的消息
整体流程：
项目启动，用户输入用户名，点击连接按钮，启动websocket连接服务器请求。
启动用户的读，写协程。
用户的读协程：不断的读取服务端发送来的消息放入读取通道。
用户的写协程：不断的读取客户端发送的消息，发送给服务端。
fyne中开启刷新页面的协程，不断的从读取通道中读取数据：
	1.当消息类型为“login”，将用户ip，广播内容重新渲染到页面。
	2.当消息类型为“user_list”,将用户列表重新渲染到页面。
	3.当消息类型为“talk”，将用户输入的内容重新渲染到页面上。
	4.当消息类型为“exit”，将用户ip重新设为空，用户列表清空。
```
![流程图](./doc/img.jpg)

# 目录结构
```text
.
├── general.proto            //消息内容的proto文件
├── global
│   └── globalVar.go         //全局变量文件
├── go.mod
├── main
│   └── main.go              //项目启动文件
├── model
│   └── client.go            //客户端管理
├── response
│   └── general.pb.go        //消息内容的proto的go文件
├── service
│   └── clientService.go。   //客户端服务层
├── test
│   └── project6_test.go     //单元测试
└── view
    └── myApp.go             //设计页面，渲染页面
```
# 代码逻辑分层
层|文件夹|主要职责
------------ | ------------- | ------------- 
服务层|service|负责处理客户端发送消息的逻辑
实体层|model|封装客户端的读写协程代码
单元测试层|test|测试是否能连接到服务端
启动层|main|启动项目，调用应用层处理http请求
fyne页面渲染层|view|将用户输入的内容，重新渲染到页面
全局变量层|global|项目中的全局变量

# 存储设计
### 用户输入的数据封装为proto结构。格式为：
内容|field|类型
------------ | ------------- | ------------- 
用户ip地址|Ip|string
用户消息类型|Type|string
消息用户的来源|From|string
用户输入的内容|Content|string
用户名字|User|string
用户在线列表|Userlist|[]string


# 接口设计
### 1.用户发起连接请求，建立连接
### websocket接口地址
ws://localhost:8080/ws

### 发送登录消息
field|注释|类型|内容
------------ | ------------- | ------------- | -------------
Type|用户消息的类型|string|"login"
Content|用户输入的内容|string|用户的名字

### 发送获取用户列表消息
field|注释|类型|内容
------------ | ------------- | ------------- | -------------
Type|用户消息的类型|string|"user_list"

### 返回消息
field|注释|类型
------------ | ------------- | -------------
Ip|用户Ip地址|string
Content|用户输入的内容|string
Type|用户消息的类型|string
From|哪个用户说的话|string
Userlist|用户列表|[]string


### 2.用户发起断开连接请求，断开连接
### websocket接口地址
ws://localhost:8080/ws

### 发送登录消息
field|注释|类型|内容
------------ | ------------- | ------------- | -------------
Type|用户消息的类型|string|"exit"

### 发送获取用户列表消息
field|注释|类型|内容
------------ | ------------- | ------------- | -------------
Type|用户消息的类型|string|"user_list"

### 返回消息
field|注释|类型
------------ | ------------- | -------------
Ip|用户Ip地址|string
Content|用户输入的内容|string
Type|用户消息的类型|string
From|哪个用户说的话|string
Userlist|用户列表|[]string

### 3.用户发送talk类型消息，发送用户用户输入的内容
### websocket接口地址
ws://localhost:8080/ws

### 发送登录消息
field|注释|类型|内容
------------ | ------------- | ------------- | -------------
Type|用户消息的类型|string|"talk"
Content|用户输入的内容|string|用户输入的内容

### 返回消息
field|注释|类型
------------ | ------------- | -------------
Ip|用户Ip地址|string
Content|用户输入的内容|string
Type|用户消息的类型|string
From|哪个用户说的话|string
Userlist|用户列表|[]string

# 第三方库
### websocket
```text
github.com/gorilla/websocket
使用websocket框架处理用户聊天信息
```

### proto
```text
github.com/golang/protobuf
使用此包处理protobuf格式消息的解析
```

### fyne
```text
fyne.io/fyne
使用此包制作客户端页面
```
