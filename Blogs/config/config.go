package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}
type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	QiniuBucket     string
	QiniuDomain     string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}
type TomlConfig struct {
	Viewer Viewer
	System SystemConfig
}

var Cfg *TomlConfig

func init() {
	Cfg = new(TomlConfig)
	Cfg.System.AppName = "Blog"
	Cfg.System.Version = 1.0
	Cfg.System.CurrentDir, _ = os.Getwd()

	//读问配置文件并赋值给Cfg
	_, err := toml.DecodeFile("config/config.toml", &Cfg)
	if err != nil {
		log.Println("配置文件解析错误：", err)
		panic(err)
	}

}
