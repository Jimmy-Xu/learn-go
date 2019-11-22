客户端接收gamil
-------------------

# create credential

## 方法1
https://console.developers.google.com/apis/credentials

- 选择一个gmail账号
- 选择一个project
- 创建Credentials
  - 创建 OAuth 2.0 client IDs， 下载json到credentials.json
    
# 方法2

https://developers.google.com/gmail/api/quickstart/go
点击 "Enable the Gmail API", 下载json到credentials.json

# 编译

```
$ go build
```

# 运行

第一次执行时，token不存在，需要打开url，获取token，并粘贴。 token会被自动保存到token.json。 之后运行就直接读取token。
```
$ ./gmail-client
Go to the following link in your browser then type the authorization code: 
https://accounts.google.com/o/oauth2/auth?access_type=offline&client_id=407313515723-ckjjjd1nhkma4eo24i1tcfo27775nodm.apps.googleusercontent.com&redirect_uri=urn%3Aietf%3Awg%3Aoauth%3A2.0%3Aoob&response_type=code&scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fgmail.readonly&state=state-token
```
- 程序hang住
- 浏览器打开上述链接
- 选择gmail账号
- 点"允许"后得到token，复制并粘贴到上述hang处

```
4/sAF-KC2tHPxxxxxxxxxxxxxxxxxxxxxxxxxxVh0lVwuH35V--S_M  (粘贴)
Saving credential file to: token.json
Labels:
- site24x7
- Devops Monitor
- To Me
...

```

