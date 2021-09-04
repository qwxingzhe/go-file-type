# go-file-type
[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/qwxingzhe/go-file-type)
[![Build Status](https://travis-ci.org/qwxingzhe/go-file-type.svg)](https://travis-ci.org/qwxingzhe/go-file-type)
[![GitHub release](http://img.shields.io/github/release/qwxingzhe/go-file-type.svg?style=flat-square)](release)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](license)
[![Commit activity](https://img.shields.io/github/commit-activity/m/qwxingzhe/go-file-type)](https://github.com/qwxingzhe/go-file-type/graphs/commit-activity)
[![Average time to resolve an issue](http://isitmaintained.com/badge/resolution/qwxingzhe/go-file-type.svg)](http://isitmaintained.com/project/qwxingzhe/go-file-type "Average time to resolve an issue")
[![Percentage of issues still open](http://isitmaintained.com/badge/open/qwxingzhe/go-file-type.svg)](http://isitmaintained.com/project/qwxingzhe/go-file-type "Percentage of issues still open")



<p align="center">go-file-type 是一个用go开发的获取文件类型的工具。 它的主要代码参考自 <a href="https://www.cnblogs.com/enjong/articles/10741244.html">Go语言 通过文件流判断文件头来识别文件类型</a>， 你可以很轻易的在任何 GO 项目中使用它。</p>



# 安装

```shell
go get -u github.com/qwxingzhe/go-file-type
```

# 使用指南

通过网络文件地址获取扩展名

~~~
package main

import (
	"fmt"
	gofiletype "github.com/qwxingzhe/go-file-type"
)

func main() {
	fileType := gofiletype.GetFileTypeByUrl("http://img.gif.cn/temp_makegif/20210806/1628230128817044.gif")
	fmt.Println(fileType)
}
~~~

通过文件前面几个字节来判断

~~~
f, _ := os.Open("./testfile/0010.gif")
fSrc, _ := ioutil.ReadAll(f)
fileType := gofiletype.GetFileTypeByByte(fSrc[:10])
fmt.Println(fileType)
~~~

通过本地文件路径获取扩展名

~~~ 
fileType := gofiletype.GetFileTypeByLocalPath("./testfile/0010.gif")
fmt.Println(fileType)
~~~


## 测试用例

- [x] gif
- [x] jpg
- [x] png
- [ ] webp

# Enjoy it! :heart:

# 参照

- [Go语言 通过文件流判断文件头来识别文件类型](https://www.cnblogs.com/enjong/articles/10741244.html)

# License

MIT

