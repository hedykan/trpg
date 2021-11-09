# 接口文档
## 通用参数
### 必传header: Authorization
用户token，用以识别用户身份
### 回调结构
```json
{
    "code": 200,
    "data": {Object},
    "msg": "ok"
}
```

## 故事编辑模块接口 (kp)
### 故事展示api
请求方式: `get`  
接口uri: `/story/list`  
接口参数:  
```json
{
    // 房间id
    "roomId": int
}
```
回调参数:
```json
{
    // 节点Id
    "Id": int,
    // 节点内容
    "Val": string,
    // 通向节点的节点号记录
    "Input": {"Id": int, "Val": string},
    // 节点选择通向的节点号记录
    "Output": {"Id": int, "Val": string}
}
```

### 故事节点获取
请求方式: `get`  
接口uri: `/story/get`  
接口参数:  
```json
{
    // 房间id
    "roomId": int,
    // 故事节点id
    "nodeId": int
}
```
回调参数:
```json
{
    // 节点Id
    "Id": int,
    // 节点内容
    "Val": string,
    // 通向节点的节点号记录
    "Input": {"Id": int, "Val": string},
    // 节点选择通向的节点号记录
    "Output": {"Id": int, "Val": string}
}
```

### 故事节点新增
请求方式: `post`  
接口uri: `/story/node_add`  

### 故事节点编辑
请求方式: `post`  
接口uri: `/story/node_edit`  

### 故事节点删除
请求方式: `get`  
接口uri: `/story/node_delete`  

### 故事选择新增
请求方式: `post`  
接口uri: `/story/selecter_add`  

### 故事选择删除
请求方式: `post`  
接口uri: `/story/selecter_delete`  

## 跑团操作模块接口 (kp)
### 跑团故事选择
请求方式: `get`  
接口uri: `/run/get`  

### 跑团故事回退
请求方式: `get`  
接口uri: `/run/return`  

## 跑团操作模块接口 (kp/pl)
### 跑团状态展示
请求方式: `get`  
接口uri: `/run/status_list`  

### 跑团故事背景获取
请求方式: `get`  
接口uri: `/run/story_background_get`  

### 跑团当前节点获取
请求方式: `get`  
接口uri: `/run/now_node_get`  

### 跑团当前投票获取
请求方式: `get`  
接口uri: `/run/now_vote_get`  

### 跑团投票添加
请求方式: `get`  
接口uri: `/run/vote_add`  

### 跑团当前已经过节点展示
请求方式: `get`  
接口uri: `/run/now_record_list`  

## 用户身份模块接口 (kp/pl)
### 用户身份确认
请求方式: `get`  
接口uri: `/auth/check`  

### 用户状态请求
请求方式: `get`  
接口uri: `/auth/status`  
