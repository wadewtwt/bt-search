## 安装
1.导入storage/ai.sql到mysql

2.拉取依赖
> go mod tidy

3.创建.env文件以及配置
> cat .env.example > .env


## 启动
```shell
go run main.go

# or

go build main.go && ./main
```

## nginx配置
我这边修改了nginx默认端口号为8888，前端请求到127.0.0.1:8888然后转发到9501上
start nginx
nginx -s stop

```shell
location ~ (/api/|/admin/) {
    add_header Access-Control-Allow-Origin *;
    add_header Access-Control-Allow-Methods 'GET, POST, OPTIONS, DELETE';
    add_header Access-Control-Allow-Headers 'DNT,Keep-Alive,User-Agent,Cache-Control,Content-Type,Authorization,X-Token';
    if ($request_method = 'OPTIONS') {
        return 204;
    }
    # 将客户端的 Host 和 IP 信息一并转发到对应节点
    proxy_set_header Host $http_host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    
    # 执行代理访问真实服务器
    proxy_pass http://127.0.0.1:9501;
}
```
服务器配置
镜像拉不下来可能是没有登录阿里云的仓库，复制命令时候记得改下账户名
然后按照教程下载amd版本
https://www.cnblogs.com/warrenwt/p/18109103

## 参考
> 去水印copy该项目部分: https://github.com/IHuan123/watermark-server