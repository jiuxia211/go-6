# go-6

gin+grpc+etcd

西二在线golang第六轮考核~~~~

接口文档 https://www.showdoc.com.cn/2275177058881415?page_id=10206616509552948

#### 运行

1.在config.ini 的mysql部分修改对应的数据库信息，并先手动创建etcd

2.启动etcd

启动用户模块 make build-user

启动事务模块 make build-task

启动网关 make build-gateway

#### 项目结构

总体

├─conf #配置文件
├─gateway #网关层
│  ├─cmd
│  ├─internal
│  │  ├─handler
│  │  └─service
│  │      └─pb
│  ├─middlerware
│  └─routes
├─pkg 
│  ├─discovery
│  ├─e
│  ├─res
│  └─util
├─task #事务层
│  ├─cmd
│  └─internal
│      ├─handler
│      ├─respository
│      └─service
│          └─pb
└─user #用户层
    ├─cmd
    └─internal
        ├─handler
        ├─respository
        └─service
            └─pb

##### 用户模块

├─cmd
│      main.exe
│      main.go
│
└─internal 
    ├─handler
    │      user.go
    │
    ├─respository  #数据库操作
    │      init.go
    │      user.go
    │
    └─service
        │  userModels.pb.go
        │  userService.pb.go
        │  userService_grpc.pb.go
        │
        └─pb
                userModels.proto
                userService.proto

##### 事务模块

├─cmd
│      main.exe
│      main.go
│
└─internal
    ├─handler
    │      task.go
    │
    ├─respository  #数据库操作
    │      init.go
    │      task.go
    │      user.go
    │
    └─service
        │  taskModels.pb.go
        │  taskService.pb.go
        │  taskService_grpc.pb.go
        │
        └─pb
                taskModels.proto
                taskService.proto

##### 网关层

├─cmd
│      main.exe
│      main.go
│
├─internal
│  ├─handler  service接口的调用
│  │      task.go
│  │      user.go
│  │
│  └─service #proto文件和对应service服务
│      │  taskModels.pb.go
│      │  taskService.pb.go
│      │  taskService_grpc.pb.go
│      │  userModels.pb.go
│      │  userService.pb.go
│      │  userService_grpc.pb.go
│      │
│      └─pb
│              taskModels.proto
│              taskService.proto
│              userModels.proto
│              userService.proto
│
├─middlerware
│      cors.go  #跨域
│      init.go   #利用中间件 获取grpc返回的conn
│      jwt.go  
│
└─routes  gin路由

​        routes.go

##### pkg包

├─discovery #etcd服务发现的实现
│      instance.go
│      register.go
│      resolver.go
│
├─e
│      code.go
│      msg.go
│
├─res
│      response.go
│
└─util
        token.go