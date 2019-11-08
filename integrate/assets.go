package demo

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets9f0ae851af9adb08e3242917d7d52dc270c7bae2 = "<!doctype html>\r\n\r\n<head>\r\n    <meta charset=\"UTF-8\">\r\n    <title>demo</title>\r\n    <meta name=\"renderer\" content=\"webkit|ie-comp|ie-stand\">\r\n    <meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge,chrome=1\">\r\n    <meta name=\"viewport\" content=\"width=device-width,user-scalable=yes, minimum-scale=0.4, initial-scale=0.8\" />\r\n    <meta http-equiv=\"Cache-Control\" content=\"no-siteapp\" />\r\n</head>\r\n\r\n<body>\r\n\r\n    <p>hello <span>world</span></p>\r\n    \r\n    <script src=\"/static/js/jquery.min.js\" charset=\"utf-8\"></script>\r\n    <script>\r\n        (function () {\r\n            $(\"span\").text(\"golang\");\r\n        })()\r\n    </script>\r\n</body>\r\n\r\n</html>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{}, map[string]*assets.File{
	"index.html": &assets.File{
		Path:     "index.html",
		FileMode: 0x1ff,
		Mtime:    time.Unix(1573195761, 1573195761179625600),
		Data:     []byte(_Assets9f0ae851af9adb08e3242917d7d52dc270c7bae2),
	}}, "")
