package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	asset "demo/integrate"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/gin"
)

// loadTemplate 加载由 go-assets-builder 嵌入的模板
func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range asset.Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".html") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func main() {
	router := gin.Default()

	// 指定html的全局引用路径
	// router.LoadHTMLGlob("templates/*")
	// html文件包含层级（子目录）的话要这么写
	// router.LoadHTMLGlob("templates/**/*")

	// 然后注释下面这行
	// router.LoadHTMLGlob("templates/*")
	// 改成下面的部分
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	router.SetHTMLTemplate(t)
	// 这部分代码要写在router := gin.Default()这行的后面

	// 声明静态文件的路由以/static开头
	// 并指向resources文件夹下的文件
	// router.Static("/static", "resources")

	// 注释掉原先我们的静态资源引入方式的代码
	// 就是下面这行
	// router.Static("/static", "resources")
	// 然后从静态资源打包的文件中获取
	fs := assetfs.AssetFS{
		Asset:    asset.Asset,
		AssetDir: asset.AssetDir,
		Prefix:   "resources",
	}
	router.StaticFS("/static", &fs)

	// 测试页面
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
		// 注意前面打包html的参数:-s="/templates/"
		// 不加这个参数的话上面这行代码要改成下面这样:
		// c.HTML(http.StatusOK, "/templates/index.html", nil)
	})

	// 获取服务运行端口
	// 主要是为了后面测试静态文件打包后的读取
	// jsonContentBytes, err := ioutil.ReadFile("config/config.json")
	jsonContentBytes, err := asset.Asset("config/config.json") // 根据地址获取对应内容
	if err != nil {
		panic(err)
	}
	type Config struct {
		Port uint `json:"port"`
	}
	config := new(Config)
	err = json.Unmarshal(jsonContentBytes, config)
	if err != nil {
		panic(err)
	}
	runSetting := fmt.Sprintf(":%d", config.Port)

	router.Run(runSetting)
}
