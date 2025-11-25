需要实现一个 make release
make release 编译完成，会得到一个 bin/version.txt 文件，这个文件记录版本号，比如：

VERSION=0.1.5

FASTPVE_SHA256=5974b91a2634dcdb4dd03eda41f80cddd6898a854c11650f2229528f810576d0

并且最终编译带 HAS_REMOTE_URL 的结果 ./bin/FastPVE-0.1.5

这样实现方便我们提供一个 fastpve-install.sh 的时候，从默认地址下载 version.txt 不缓存，而 ./bin/FastPVE-0.1.5 可以进行缓存下载

我们的URL 地址是：
https://fw.kspeeder.com/binary/fastpve/version.txt

https://fw.kspeeder.com/binary/fastpve/fastpve-install.sh

https://fw.kspeeder.com/binary/fastpve/FastPVE-0.1.5

.txt .sh 结尾的都是不缓存的。
