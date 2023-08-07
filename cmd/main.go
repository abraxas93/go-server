package main

import (
	"go-server/pkg/config"
	"go-server/pkg/logger"
)

func main() {
	cfg := config.GetConfig(".env")
	logger.InitLogger(make(map[string]string))
	logger, _ := logger.GetLogger()
	logger.Info("%+v\n", cfg)
	// postgres.Connect()
	// filename := "config/local.yaml"
	// data, err := os.ReadFile(filename)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// c := &Cfg{}
	// err = yaml.Unmarshal([]byte(data), c)

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// logger.Info("%+v\n", c)

}
