# 跳动字节Hertz框架上传minio

## 使用了hertz框架上传文件到minio，并且返回连接

项目启动方式用了nacos配置中心和本地

启动方式拉去项目 ，然后

 本地启动方式:
```
go run . -ctype local -chost ./config/config
```

nacos配置

```
go run . -ctype nacos -chost 127.0.0.1:8848
```

### 接口

http://127.0.0.1:8765/upload

参数 file

想要获取永久链接需要先配置bucket的访问权限