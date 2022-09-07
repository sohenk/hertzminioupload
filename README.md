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

#### 直接上传文件

http://127.0.0.1:8765/upload

```参数 file```

返回参数

```
{
    "code": 200,
    "message": "success",
    "reason": "success",
    "data": {
        "FileName": "微信图片_20211220093130.jpg",
        "FileSize": 127248,
        "FileType": "image/jpeg",
        "FileUrl": "文件链接"
    }
}
```

### 浏览器直传

http://127.0.0.1:8765/getuploadurl

提交参数
```
filename : 上传文件名
```

返回参数
```
{
    "code": 200,
    "message": "success",
    "reason": "success",
    "data": {
        "UploadUrl": "上传文件url", //获取了上传文件url后，直接put文件到这个url那里。主要要用“put”方式
        "Expires": "上传过期时间",
        "FileUrl": "文件url"
    }
}
```


· 想要获取永久链接需要先配置bucket的访问权限


###