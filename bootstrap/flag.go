package bootstrap

import "flag"

type CommandFlags struct {
	Conf       string
	Env        string
	ConfigHost string
	ConfigType string
}

func NewCommandFlags() *CommandFlags {
	return &CommandFlags{
		Conf:       "",
		Env:        "",
		ConfigHost: "",
		ConfigType: "",
	}
}

func (f *CommandFlags) Init() {
	flag.StringVar(&f.Conf, "conf", "./config", "config path, eg: -conf bootstrap.yaml")
	flag.StringVar(&f.Env, "env", "dev", "runtime environment, eg: -env dev dev是开发模式prod是生产模式")
	flag.StringVar(&f.ConfigHost, "chost", "127.0.0.1:8848", "config server host, eg: -chost 127.0.0.1:8848")
	flag.StringVar(&f.ConfigType, "ctype", "nacos", "config server host, eg: -ctype nacos")
}
