# ss5

### 安装

```bash
go get https://github.com/du5/ss5
```

### 命令行参数 


```bash
Usage of ss5:
  -u string
        授权账户 (default "")
  -s string
        授权密码 (default "")
  -p int
        监听端口 (default 8080)
  -b string
        监听地址 (default "0.0.0.0")
```

#### 示例

```bash
# 在 8080 端口开启代理，无身份验证
ss5
```

```bash
# 在 4399 端口开启代理，无身份验证
ss5 -p 4399
```

```bash
# 在 4399 端口开启代理，使用用户名 "myss" 密码 "mypass" 验证身份
ss5 -u myss -s mypass -p 4399
```


```bash
# 在 IP "127.0.0.1" 开启监听，使用用户名 "myss" 密码 "mypass" 验证身份
ss5 -u myss -s mypass -b 127.0.0.1
```

### Docker 

```bash
# 在 145 端口开启代理，使用用户名 "myss" 密码 "mypass" 验证身份
docker run --name=ss5 -d --restart=always -p 145:145 gtary/ss5 \
    --p 145 \
    --u myss \
    --s mypass
```