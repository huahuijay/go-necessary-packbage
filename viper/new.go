package main

import (
	"fmt"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

func main() {
	viper.AddSecureRemoteProvider("etcd","http://127.0.0.1:4001","/config/config.toml",".secring.gpg")
	viper.SetConfigType("toml")
	go func(){viper.WatchRemoteConfig()}()
	err := viper.ReadRemoteConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(viper.GetString("Mysql.port"))
	var w string
	fmt.Scanf("%s", &w)
	fmt.Println(w)
	fmt.Println(viper.GetString("Mysql.port"))
}