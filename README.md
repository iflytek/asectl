# asectl
A cli tool for users to develop and deploy the service quickly under the ASE framework

### Prerequisites
- golang 1.15+

### How to build?
```
$ git clone https://github.com/xfyun/asectl.git
$ cd asectl
$ go mod tidy
$ go mod vendor
$ make
```

### Components

1. 命令行配置模块 cfg
2. Wrapper模块
   1. python wrapper 目录初始化
   2. python wrapper run/debug
   3. python wrapper build to image
   4. python wrapper pack 打包成插件包(support ase)

3. 部署模块
   1. deploy by image
   2. deploy by code 
   
4. AI服务管理
   1. 服务创建
   2. 服务配置
   3. 服务 CRUD 
   4. 服务接口文档
   5. 服务测试
   6. 服务schema


### How to use asectl?
todo
