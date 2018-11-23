package main

import (
	_ "Guardian/routers"
	"github.com/astaxie/beego"
	"Guardian/models"
)

func main() {
	models.InitDB()

	//_, err := models.GetSpaceNum()
	//if err != nil {
	//	fmt.Println(err)
	//}
	beego.Run()
}