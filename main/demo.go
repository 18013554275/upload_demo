package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func main(){
	verifyToken("");
}

func uploadImg(){
	router := gin.Default()
	router.POST("/upload", func(c *gin.Context){
		name := c.PostForm("name");
		fmt.Println(name)
		file, err := c.FormFile("upload")
		if err != nil {
			c.String(http.StatusBadRequest, "a Bad request")
			return
		}
		filename := file.Filename
		fmt.Println("========", filename)
		if err := c.SaveUploadedFile(file, file.Filename); err != nil{
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
		c.String(http.StatusBadRequest, "upload successful")
	})
	router.Run(":3333")
}

func verifyToken(token string ) bool{
	//发送get请求
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	//响应状态
	fmt.Println(resp.Status)
	fmt.Println(resp.Header)
	//获取响应内容
	c := make([]byte, 2048)
	var result string
	for {
		n, err := resp.Body.Read(c)
		if err != nil && io.EOF == err {
			break
		}
		result += string(c[:n])
	}
	fmt.Println("result=" + result)

	return false;
}

