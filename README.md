# gohttpclient

个人认为封装了一些httpclient难以忍受的细节。当然只是做一些常用场景的封装。用的是装饰器模式，这种模式更适合多种情况的交叉组合。但是并非设计很完美。

因此该代码库定位为，轻量级，易用，解决常用场景，That‘s all，保持简单。

后面可能会根据具体情况做一些扩展，看情况。

大概就这么用：

```
code,body,err :=gohttpclient.Get("https://www.baidu.com").Exec().String()


if err != nil {
    //http请求过程中产生的错误，url不对？，网络异常？等等。
		t.Fatal("error:", err)
}

if code != http.StatusOK {
    //根据response code可以做一些处理，和其他错误分开比较好。
}



```
