一、关于 module、package和repository

1、go由packages构成，package是同一个目录中的一组需要被同时编译的源文件。
在同一个package不同源文件中定义的functions、types、variables、constants是相互可见的。

2、repository包含一个或多个modules。module是一起发布的一组相关的packages。
典型的go repository只包含一个module，位于repository的root。
一个名为go.mod的文件用于定义module path：该module中所有packages的the import path prefix
这个module包含go.mod文件所在的目录及其子目录，直到某个目录中存在另一个go.mod文件。

3、每个module path不仅表明了其中packages的import path prefix，还表明了go命令应该去哪里下载它。
例如对于golang.org/x/tools module，go命令需要去https://golang.org/x/tools下载。

4、import path是import packages时使用的字符串。
一个package的import path是module path加上该包在module中的子目录。
例如：github.com/google/go-cmp包含一个package在其子目录cmp中，
那么该包的import path就是github.com/google/go-cmp/cmp。
另外：标准库中的包不需要module path prefix。


二、关于命令行
1、go build将不会输出结果文件到当前目录，而是输出到local build cache。
2、go install的作用范围只包括当前工作目录所含有的module。
3、module dependencies会自动下载到$GOPATH/pkg/mod。使用go clean -modcache来清除所有downloaded modules。

