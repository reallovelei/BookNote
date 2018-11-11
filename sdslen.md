 各位朋友大家好，欢迎来到解码俱乐部，我是晓磊哥，我是大路  
上一期留下了一个问题，SDS最大可以容纳多少字节。  
关于这个问题 最常见的是两个答案：1 根据len的数据类型来判断 len存储的最大值. 2的31次方-1.  


创建一个新的SDS的时候，代码里并没有调用这个 checkstringlen 函数，  
我用php  golang 分别set 一个超过512m大小的value.  
都报错,但这2个语言链接redis都用第三方模块链接的。  
我于是使用redis-cli 直接set 一个 超过512M大小字符串。  

用如下方式就可以直接用cli 导入命令。  文件如下：  


还是报错：  

但是这次报错信息有所改变。  
protocol type err
这就把我的思路转向到协议层面。  

我们可以看到不论是PHP 还是 golang，还是redis-cli 的报错信息都和checkstringlen 是不一致的。  
也就说直接生成一个新的sds的时候  大小是否超过512m  的判断并不是这个函数在起作用。  



我在继续去代码里寻找答案，发现loadconfig的时候有
而是根据redis.conf里的proto-max-bulk-len  

redis 默认设置了512M的大小, 但是我们可以根据配置文件里的参数进行调整。我觉得那个setrange 和append判断的地方其实也应用这个配置变量来判断。  

