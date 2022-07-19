package global

import (
	"hzminioupload/biz/pkg/global/systeminit"

	"github.com/spf13/viper"
)

var (
	S_CONFIG      *viper.Viper
	S_MinioClient *systeminit.MinioClient
)
