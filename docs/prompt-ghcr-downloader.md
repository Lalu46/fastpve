我们实现了项目：
/projects/workspace-linkease-ubuntu/linkease-github/blobDownload

用法可以参考：
/projects/workspace-linkease-ubuntu/linkease-github/BlobTran

就是说呢，如果要下载 win10 win11 我们目前可以直接从 ghcr.io 进行下载，具体为：

ghcr.io/kspeeder/win10x64:cn_simplified

ghcr.io/kspeeder/win11x64:cn_simplified

上面相当于 --url windows 10 "Chinese (Simplified)"

--url windows 11 "Chinese (Simplified)"

有时候用 quickget 无法获取 windows URL，或者用 remotecacheurl 也无法获取 url，那么我们最后的办法就是，从 ghcr ，调用 blobDownload 进行下载。并且这个还可以指定具体的文件名字。这样就 保证了，任何时候都可以下载到 windows 10 windows 11 的 chinese 的文件了。至于 en 文件暂时没有，以后再来加入到 ghcr。
而且 ghcr 可能是有 mirror 的，比如 默认用 mirror 地址为：ghcr.1ms.run
用 mirror 下载更快。