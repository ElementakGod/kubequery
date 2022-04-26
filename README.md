# kubequery
    参照osquery的思想，将k8s的所有资源作为一个关系型数据，使用
    sql的方式进行k8s资源查询 
***
## 架构
### 模式
* kubequeryd 后台进程模式
* kubequeryi 交互模式
### 模块
* k8s-watch  k8s资源watch
* kubequery-cli 交互client
* watchdog 监听慢查询，待后续优化
##选型
* k8s-watch: client-go
* kubequery-cli: 复用sqlite的shell交互
* watchdog: 协程
## 例子
    ==> select name, ip from pod
