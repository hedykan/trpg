# 跑团！！！

## 说明
这是一个基于go实现的跑团api，由几个模块组成：
1. 故事系统 (DOWN)
2. 跑团操作系统 (DOWN)
3. 掷骰系统
4. 属性系统
5. 角色系统

## 操作说明
在根目录`go build`后会生成目标可执行程序。  
运行后将监听`:12345`端口。  

## 目前支持操作
1. 故事展示`/story/list`
2. 故事节点获取`/story/get?id`
3. 故事运行状态获取`/run/status_list`
4. 故事运行执行`/run/step?id`
```bash
# 示例
$ curl http://127.0.0.1:12345/store/list
{"code":200,"data":[{"Id":1,"Val":"hello world","Input":[0,1],"Output":[2]},{"Id":2,"Val":"hello world","Input":[0,1],"Output":[2]}],"msg":"ok"}
```