package main

import (
	"fmt"
	"raytheon/datamodels"
	"raytheon/web/routers"
	"runtime"
	// _ "github.com/go-sql-driver/mysql"
	"raytheon/utils"
)

func main() {
	var err error

	runtime.GOMAXPROCS(runtime.NumCPU())
	// utils.InitCasbin()
	utils.DBConn.AutoMigrate(&datamodels.User{}, &datamodels.Tenants{}, &datamodels.UserTenant{})

	defer func() {
		if fErr := utils.DBConn.Close(); fErr != nil {
			err = fErr
		}
	}()
	if err != nil {
		panic(err)
	}

	router := routers.InitRouter()

	err = router.Run(fmt.Sprintf(":%d", utils.APIConfig.MainConfig.Port))
	if err != nil {
		panic(err)
	}
}
