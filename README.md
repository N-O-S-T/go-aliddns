#     

[![Security Status](https://www.murphysec.com/platform3/v31/badge/1721471504210288640.svg)](https://www.murphysec.com/console/report/1721468007020584960/1721471504210288640)

## 介绍

用golang实现aliddns，同时对certbot进行txt信息更新提交

## 配置文件

- 设置系统变量 `ALIDDNSCONFIG` 指向配置文件名
  ```shell
    export ALIDDNSCONFIG=/xxxx/xxxx/xxxx.yaml
  ```
- 配置文件格式如下
  ```yaml
  AccessKeyId: xxxxxxxxxx #阿里云id
  AccessSecret: xxxxxxxxxxxxxxxxxxxx #阿里云token
  MainDomain: xxx.xxx #顶级域名
  SubDomains:
  - SubDomain: xx
    ipv4: false
    ipv6: true
    cert: false
  - SubDomain: cc
    ipv4: false
    ipv6: true
    cert: false
  ```

## 更新dns解析

```shell
aliddns # 如果采用默认配置，且目标为更新dns，则只需要此条命令
```

## 证书相关

此功能改版后未经过完整测试

### 申请证书

注意，aliddns通过读取config.yaml来决定更新内容

```shell
certbot certonly  -d *.example.com --manual --preferred-challenges dns --dry-run  --manual-auth-hook "aliddns -rt cert -f /xxxx/xxx.yaml"
```

### 更新证书(同时重启gitlab nginx服务)

```shell
certbot renew  --manual --preferred-challenges dns --manual-auth-hook "aliddns -rt cert" --deploy-hook "gitlab-ctl restart nginx"
```