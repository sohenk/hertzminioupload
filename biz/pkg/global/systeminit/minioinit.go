package systeminit

import (
	"context"
	"io"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

type MinioClient struct {
	Client         *minio.Client
	BucketName     string
	TargetFilePath string
	UseSSL         bool
	Location       string
	ExposeUrl      string
}

func MinioClientInit(config *viper.Viper) (*MinioClient, error) {
	var endpoint = config.GetString("minio.endpoint")               //minio地址
	var accessKeyID = config.GetString("minio.accessKeyID")         //账号
	var secretAccessKey = config.GetString("minio.secretAccessKey") //密码
	var useSSL = config.GetBool("minio.useSSL")                     //使用http或https
	target, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Println("minio client init error:", err)
		return nil, err
	}
	bucketName := config.GetString("minio.bucket")
	exists, err := target.BucketExists(context.Background(), bucketName)
	if err != nil {
		return nil, err
	}
	if !exists {
		err = target.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: ""})
		if err != nil {
			// Check to see if we already own this bucket (which happens if you run this twice)
			exists, errBucketExists := target.BucketExists(context.Background(), bucketName)
			if errBucketExists == nil && exists {
				log.Printf("We already own %s\n", bucketName)
			} else {
				log.Fatalln(err)
			}
		} else {
			log.Printf("Successfully created %s\n", bucketName)
		}
	}

	return &MinioClient{
		Client:     target,
		BucketName: config.GetString("minio.bucket"), //目标bucket
		// Location:       "cn-north-1",                             //放在哪个地方，这里为中国
		TargetFilePath: config.GetString("filedriver.storepath"), //资源文件夹的路径
		ExposeUrl:      config.GetString("minio.exposeurl"),
	}, nil
}

func (c *MinioClient) UpLoadFile(ctx context.Context, objectName, contentType string, file io.Reader) (string, error) {
	objectName = c.TargetFilePath + "/" + objectName

	log.Println("upload file:", objectName)
	_, err := c.Client.PutObject(ctx, c.BucketName, objectName, file, -1, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Println("upload file error:", err)
		return "", err
	}
	// urls, err := c.Client.PresignedPutObject(ctx, c.BucketName, objectName, time.Second*24*60*60*7)
	// if err != nil {
	// 	log.Println("upload file error:", err)
	// 	return "", err
	// }
	// fmt.Println("upload file success:", urls)
	fullurl := c.ExposeUrl + "/" + objectName
	return fullurl, nil
}
