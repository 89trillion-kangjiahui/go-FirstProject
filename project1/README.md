# 需求1：app.ini配置文件解析
### 通过第三方go包github.com/robfig/config读取.int配置文件，使用的函数如下：
```java
	c, _ := config.ReadDefault("./config/app.ini").  //读取配置文件
	httpPort, _ := c.String("server", "HttpPort")  //获取server标签下的HttpPort字段值
```

# 需求2：读取命令行参数，获得json文件的路径：
### 通过第三方go包github.com/spf13/pflag读取命令行参数json-path，使用的函数如下：
```java
var CliJsonPath = flag.StringP("json-path", "p", "./config/config.army.model.json", "Input Json Path")
func GetJsonPath() string {
	// 设置标准化参数名称的函数
	flag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc)
	flag.Parse()
	return *CliJsonPath
}
```
# 需求3：输入稀有度，当前解锁阶段，获取该稀有度合法且已解锁的所有士兵 
### http请求方法：
http  GET
### 接口地址：
localhost:8000/soldier/getAll
### 请求参数：
参数名|参数类型|注释
------------ | ------------- | -------------
rarity｜string|稀有度
unlockArena｜string|解锁阶段
### 请求响应
参数名|参数类型|注释
------------ | ------------- | -------------
code|json|状态码
returnData|json|所有士兵

### 返回json数据格式如下：
```java
[
    {
        "id":"10102",        //士兵id
        "Name":"Swordsman",  //士兵名字
        "UnlockArena":"0",   //解锁阶段
        "Rarity":"1",        //稀有度
        "Atk":"140"          //战斗力
    }
]
```

# 需求4：输入士兵id获取稀有度
### http请求方法：
http  GET
### 接口地址：
localhost:8000/soldier/getRarity 
### 请求参数：
参数名|参数类型|注释
------------ | ------------- | -------------
id｜string|士兵id
### 请求响应
参数名|参数类型|注释
------------ | ------------- | -------------
code|string|状态码
 rarity|string|士兵稀有度

### 返回数据格式如下：
```java
    "士兵的稀有度: 1"
```


# 需求5：输入士兵id获取战斗力
### http请求方法：
http  GET
### 接口地址：
localhost:8000/soldier/atc 
### 请求参数：
参数名|参数类型|注释
------------ | ------------- | -------------
id｜string|士兵id
### 请求响应
参数名|参数类型|注释
------------ | ------------- | -------------
code|string|状态码
atc|string|士兵战斗力

### 返回数据格式如下：
```java
    "士兵的战斗力: 1"
```

# 需求5：获取每个阶段解锁相应士兵的json数据
### http请求方法：
http  GET
### 接口地址：
localhost:8000/soldier/getAll/unlockArena
### 请求参数：无
### 请求响应
参数名|参数类型|注释
------------ | ------------- | -------------
code|string|状态码
ret|string|按阶段分类的所有士兵信息

### 返回数据格式如下：
```java
    [
        0:{
           {
               "id":"10102",        //士兵id
               "Name":"Swordsman",  //士兵名字
               "UnlockArena":"0",   //解锁阶段
               "Rarity":"1",        //稀有度
               "Atk":"140"          //战斗力
            }
     ]
```
