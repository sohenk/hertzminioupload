package bootstrap

import "flag"

type CommandFlags struct {
	ServiceName   string
	Conf          string
	Env           string
	ConfigHost    string
	ConfigType    string
	NacosUserName string
	NacosPassword string
}

func NewCommandFlags() *CommandFlags {
	return &CommandFlags{
		ServiceName:   "",
		Conf:          "",
		Env:           "",
		ConfigHost:    "",
		ConfigType:    "",
		NacosUserName: "",
		NacosPassword: "",
	}
}

func (f *CommandFlags) Init() {
	flag.StringVar(&f.ServiceName, "sname", "allmedia.uploadservice.v1", "service name , eg: -sname allmedia.uploadservice.v1")
	flag.StringVar(&f.Conf, "conf", "./config", "config path, eg: -conf bootstrap.yaml")
	flag.StringVar(&f.Env, "env", "dev", "runtime environment, eg: -env dev dev是开发模式prod是生产模式")
	flag.StringVar(&f.ConfigHost, "chost", "127.0.0.1:8848", "config server host, eg: -chost 127.0.0.1:8848")
	flag.StringVar(&f.ConfigType, "ctype", "nacos", "config server host, eg: -ctype nacos")
	flag.StringVar(&f.NacosUserName, "nacosusername", "nacos", "config server host, eg: -nacosusername nacos")
	flag.StringVar(&f.NacosPassword, "nacospassword", "nacos", "config server host, eg: -nacospassword nacos")
}
