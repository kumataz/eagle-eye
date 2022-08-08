## EagleEye
A mining machine management system, TBD ..
Include 2 modules:
- eagle: server
- collect: a tool for collecting miner data

### How To Use

1. Edit eagle-eye/configs/config.ini and choose database
```
mysql -u [user] -p [database name]< xxx.sql
```

2. edit all route / controller / services / models ...

3. build server and run
```
cd eagle-eye/cmd/eagle/
go mod tidy; go build
```

### EagleEye Server

1. 后台功能

Jump to:
```
http://localhost:9099/admin/login
http://localhost:9099/admin/home
```

业务开发

```
// 1.1 页面注册
eagle-eye/internal/menu/menu.go
// 1.2 路由注册
eagle-eye/internal/router/admin_router.go
// 1.3 资源路径
eagle-eye/internal/controllers/admin/miner/ironController.go
   // 数据结构定义
   eagle-eye/internal/models/Miner.go
   // 定义getMiners实现方法
   eagle-eye/internal/services/admin/minerService.go
   // db
   eagle-eye/internal/dao/minerDao.go
// 1.4 前端页面   
eagle-eye/web/views/template/miner/miner.html
```

1. 接收采集数据的api

```
http://localhost:9099/admin/miner/postjson

// 2.1 路由注册
eagle-eye/internal/router/admin_router.go  
// 2.2 定义数据结构、与mysql增删改查模块
eagle-eye/internal/models/Miner.go
eagle-eye/internal/database
// 2.2 控制器配置
eagle-eye/internal/controllers/admin/miner/ironController.go
```

### EagleEye Collector

Collect data from miner, quick start: 

```
// main
eagle-eye/cmd/collect/collect.go
// collect miner and system infos
eagle-eye/internal/collector
eagle-eye/pkg/tools
```