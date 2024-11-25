package main

import "github.com/renatocardoso243/gopportunities.git/config"

var (
	logger config.Logger
)


func main() {
	logger = *config.GetLogger("main")
	//Initialize config
	err := config.Init()
	if err != nil {
		logger.Errorf("Error to initialize config: %v", err)
		return
	}
	
}