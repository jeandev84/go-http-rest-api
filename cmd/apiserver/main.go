package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/gopherschool/http-rest-api/internal/app/apiserver"
)


var (
	configPath string
)


func init()  {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}


/**
 * @ Точка входа приложения
*/
func main () {
	
	// Чтение файла
	flag.Parse()



	// Конфиг
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)


	if err != nil {
		log.Fatal(err)
	}

	// Запускать сервер
	s := apiserver.New(config)
	 
	
	// Пытаемся запустить сервер
	if err := s.Start(); err != nil {
		 log.Fatal(err)
	}
}