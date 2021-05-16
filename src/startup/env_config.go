package startup

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/wallet-tracky/Golang-backend/src/util"
)

type Config struct {
	config *viper.Viper
}

func InitializeConfig() *Config{
	c:=new(Config)
	c.config=readConfig()
	return c
}

func (c *Config) GetConfig() *viper.Viper{
	if c.config!=nil{
		return c.config
	}
	return nil
}


func readConfig()  *viper.Viper{

	v:=viper.New()
	v.AutomaticEnv()

	env:=v.GetString("ENVIRONMENT")

	if env==""{
		env="dev"
	}

	fmt.Printf("App Environment :%s\n",env)
	v.Set("env",env)
	v.SetConfigName(env)
	v.SetConfigType("yaml")
	v.AddConfigPath("config")
	err := v.ReadInConfig()

	if util.IsError(err){
		fmt.Printf("Error while reading config :%s",err.Error())
	}

	return v
}
