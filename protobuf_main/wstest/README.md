## 使用方法
1.

  1. 调用　EnterGameReq 输入openId,serverId.

  1. 调用　CreateUserReq 创建角色

  1. 添加物品 DebugAddGoodsReq




## 如果需要启动该项目到本地　
需要安装：
nodejs
protocolbuf.js
npm

个性化配置:
config.json  根据config.json.sample各自配置
public/js/proto.js 通过genproto生成


## 启动　
1. 服务器端配置sandbox = true. 并启动．

1. npm install 安装所需要包

1. npm run start 启动客户端

1. 浏览器中打开　　eg : localhost:3501

1. add deploy by pipeline.
