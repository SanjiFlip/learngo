package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// go modules 添加依赖 删除未使用的依赖项
// go get , get mode ,go install
// go mod tidy 清除mod文件中未使用的依赖
// go get -u 升级到最新的次要版本或修订版本
// go get -u=patch 升级到最新的修订版本
// go get github.com/gin-gonic/gin@version go get 是会修改go.mod文件
// 一开始我写了一个A项目，仓库是project-A,但是我代码仓库的go.mod中设置的是 github.com/xmaven/A
// B项目由于依赖了A项目，import的github.com/xmaven/project-A，使用go get命令的时候，由于package和代码仓库的名称不一样，
// replace 使用的是 go mod edit -replace github.com/xmaven/A=github.com/xmaven/project-A@v1.0.0
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
