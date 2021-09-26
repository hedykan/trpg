# 一个基于go实现的trpg框架

## 一个web-ui
可以现在本机运行trpg.exe/trpg应用程序，然后打开这个[页面](https://trpg.juhuan.store)进行自己故事的编辑和试玩。  
生成的故事脚本以及其他文件均保存在与trpg.exe相同目录的file文件夹下。  

## 是什么
这是一个基于go实现的跑团框架，可以提供跑团所需要的相应功能

## 为什么
想做一个可以像真正面团一样的跑团系统，可以有既定的故事模组团，也可以是想个开头然后慢慢补充的脑洞团。  
所有做了一个以图为结构基础的故事系统，kp可以预先写好故事直接跑，也可以在跑的过程中调用故事新增/链接接口为故事新增故事内容。

## 系统详解
这是一个基于go实现的跑团api，由几个模块组成：
1. 故事系统 (DOWN)
2. 跑团操作系统 (DOWN)
3. 掷骰系统
4. 属性系统(角色卡)
5. 情报系统(可获取的情报记录)
6. 角色系统(kp/pl分类，token分类，可换kp/pl，也可多kp)
7. 多故事系统(同时运行多个故事，可一个系统跑多个团)
总的来说，这个程序的操作员/管理员应该是kp。  
但是跑团作为一个大家都能参与的游戏来看，其实整个故事的发生与结束都是由所有的参与者决定的。  
为了保证故事的流畅性以及体验感，pl的使用权限将在下面给出。  

## 操作说明
在根目录`go build`后会生成目标可执行程序。  
运行后将监听`:12345`端口。  

## 目前支持操作
### 故事相关api(kp权限)
故事可以预先设计好，也可以现编添加，使用添加/链接接口就可以了
1. 故事展示`/story/list` GET
2. 故事节点获取`/story/get?id` GET
3. 故事节点新增`/story/node_add` POST  
    json必传参数：
    1. val：故事内容
    2. input：输入节点组，传数组
    3. output：输出节点组，传数组
4. 故事节点链接`/story/node_link` POST  
    json必传参数：
    1. val：故事内容
    2. input：输入节点，传节点号
    3. output：输出节点，传节点号
5. 故事节点删除`/story/node_delete?id` GET
6. 故事节点链接添加`/story/selecter_add` POST
    json必传参数：
    1. nodeId：当前节点号
    2. linkId：目标链接节点号
    3. val：选项内容
7. 故事节点链接删除`/story/selecter_delete` POST
    json必传参数：
    1. nodeId：当前节点号
    2. linkId：目标链接节点号

### 跑团相关api(pl/kp权限)
1. 故事运行状态重置`/run/status_reset` GET
2. 故事运行状态获取`/run/status_list` GET
3. 故事运行当前故事节点获取`/run/now_node_get` GET
4. 故事运行已经过节点列表获取`/run/now_record_lsit` GET
5. 故事运行执行`/run/step?id` GET
6. 故事运行回退`/run/return?id` GET
```bash
# 示例
$ curl http://127.0.0.1:12345/store/list
{
    "code": 200,
    "data": [
        {
            "Id": 0,
            "Val": "start",
            "Input": [],
            "Output": [
                1,
                2
            ]
        },
        {
            "Id": 1,
            "Val": "end",
            "Input": [
                0,
                3
            ],
            "Output": []
        },
        {
            "Id": 2,
            "Val": "新增节点1",
            "Input": [
                0
            ],
            "Output": [
                3
            ]
        },
        {
            "Id": 3,
            "Val": "新增节点1",
            "Input": [
                2
            ],
            "Output": [
                1
            ]
        }
    ],
    "msg": "ok"
}
```

## TODO
日志
故事模板