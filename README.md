# go-6

gin+grpc+etcd

西二在线golang第六轮考核~~~~

临时入门了一下makefile，docker还没来得及学，先这样吧

#### 运行

准备 

1.修改conf/congif.ini中的mysql部分为你自己的配置 同时创建对应名称的数据库

2.打开etcd

##### 用户模块启动 make build-user

##### 事务模块启动 make build-task

##### 网关层启动  make build-gateway 

然后在 grpc-todoList\gateway\cmd && grpc-todoList\user\cmd && grpc-todoList\task\cmd

中的 main.exe分别打开 即可正常运行(有go环境直接运行三个main.go也行)

项目结构

##### 总体

├─conf
├─gateway
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
├─task
│  ├─cmd
│  └─internal
│      ├─handler
│      ├─respository
│      └─service
│          └─pb
└─user
    ├─cmd
    └─internal
        ├─handler
        ├─respository
        └─service
            └─pb

##### 网关层

├─cmd
│      main.go
│
├─internal
│  ├─handler   #微服务的调用
│  │      task.go  
│  │      user.go
│  │
│  └─service   #proto文件及其生成的service
│      │  taskModels.pb.go
│      │  taskService.pb.go
│      │  taskService_grpc.pb.go
│      │  userModels.pb.go
│      │  userService.pb.go
│      │  userService_grpc.pb.go
│      │
│      └─pb
├─middlerware #gin路由的中间件
│      cors.go  #跨域
│      init.go  #使用中间件获取grpc的 conn
│      jwt.go  #jwt 校验token
│
└─routes   #gin路由
        routes.go

##### 用户层

├─cmd
│      main.go  
│
└─internal
    ├─handler
    │      user.go
    │
    ├─respository  #数据库加载和操作
    │      init.go
    │      user.go
    │
    └─service   #proto文件及其生成的service
        │  userModels.pb.go
        │  userService.pb.go
        │  userService_grpc.pb.go
        │
        └─pb
                userModels.proto
                userService.proto

##### 事务层

├─cmd
│      main.go
│
└─internal
    ├─handler
    │      task.go
    │
    ├─respository #数据库加载和操作
    │      init.go
    │      task.go
    └─service  #proto文件及其生成的service
        │  taskModels.pb.go
        │  taskService.pb.go
        │  taskService_grpc.pb.go
        │
        └─pb
                taskModels.proto
                taskService.proto

##### pkg

├─discovery  #etcd实现服务发现
│      instance.go
│      register.go
│      resolver.go
│
├─e 用于返回的错误码和消息
│      code.go
│      msg.go
│
├─res  序列化器
│      response.go
│
└─util  
        token.go