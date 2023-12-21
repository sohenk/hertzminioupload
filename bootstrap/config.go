package bootstrap

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	remote "github.com/sohenk/nacos-viper-remote"
	"github.com/spf13/viper"
)

func InitConfig(serviceName, configtype, configHost, username, password string) (*viper.Viper, error) {
	runtime_viper := viper.New()
	runtime_viper.SetConfigType("yaml")
	switch configtype {
	case "local":
		runtime_viper.SetConfigName(configHost)
		runtime_viper.AddConfigPath(".")
		err := runtime_viper.ReadInConfig()
		if err != nil {
			return nil, err
		}
		return runtime_viper, nil
	case "nacos":
		h := strings.Split(configHost, ":")

		if username != "" {
			username = "nacos"
		}

		addr := h[0]
		port := 8848
		if len(h) > 1 {
			uport, _ := strconv.Atoi(h[1])
			port = uport
		}
		fmt.Println("addr:", addr, "port:", port, "configtype:", configtype, "configHost:", configHost)
		remote.SetOptions(&remote.Option{
			Url:         addr,                               // nacos server 多地址需要地址用;号隔开，如 Url: "loc1;loc2;loc3"
			Port:        uint64(port),                       // nacos server端口号
			NamespaceId: "public",                           // nacos namespace
			GroupName:   "DEFAULT_GROUP",                    // nacos group
			Config:      remote.Config{DataId: serviceName}, // nacos DataID
			Auth: &remote.Auth{
				Enable:   true,
				User:     username,
				Password: password,
			}, // 如果需要验证登录,需要此参数

		})
		err := runtime_viper.AddRemoteProvider("nacos", configHost, serviceName)
		if err != nil {
			return nil, err
		}
		err = runtime_viper.ReadRemoteConfig()
		if err != nil {
			return nil, err
		}
		return runtime_viper, nil
	default:
		return nil, errors.New("config type error")
	}
}
