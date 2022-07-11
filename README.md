# Eagle Eye

A miner monitor, TBD ..


## usage

```
1. mysql config
edit eagle-eye/configs/config.ini and choose database

2. edit all route / controller / services / models

3. build and run
cd eagle-eye/cmd/eagle/
go mod tidy; go build
./eagle
```


## diretory
```
// 页面注册
ginadmin/internal/menu/menu.go
// 路由注册
ginadmin/internal/router/admin_router.go
   // 资源路径
   ginadmin/internal/controllers/admin/miner/ironController.go
      // 前端页面   
      ginadmin/web/views/template/miner/miner.html
   // 定义getMiners实现方法
   ginadmin/internal/services/admin/minerService.go
   // db
   ginadmin/internal/dao/minerDao.go
      // 数据结构定义
      ginadmin/internal/models/Miner.go
```


