package main


import (

	"github.com/gin-gonic/gin"
)

type s struct{
	a float64
}

func main(){
	r := gin.Default()

	r.Run(":8080")
	
}
