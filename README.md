# go-file-type
golang获取文件类型
参考资料：https://www.cnblogs.com/enjong/articles/10741244.html

#### 示例1
~~~
f, err := os.Open("./testfile/0010.gif")
if err != nil {
    t.Logf("open error: %v", err)
}

fSrc, err := ioutil.ReadAll(f)
t.Log(GetFileType(fSrc[:10]))
~~~

#### 示例2
~~~
GetFileTypeByPath("./testfile/0010.gif")
~~~

#### 示例3
~~~
GetFileTypeByUrl("http://img.gif.cn/temp_makegif/20210806/1628230128817044.gif")
~~~