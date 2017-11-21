# README

selpg 允许用户指定从输入文本抽取的页的范围，这些输入文本可以来自文件或另一个进程，本次实验用go语言实现该功能。功能详情点击 [开发 Linux 命令行实用程序](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html) 函数主要分为Parser，processArgs，processInput三个函数，基本沿袭了源代码的结构。

##参数解析##

本次实验的参数解析运用了Go语言中的flag包来进行解析，举个例子

> flag.IntVar(&p.start_page, "s", -1, "首页号")

当输入的参数没有-s的选项，start_page的值则默认为-1；倘若有-s选项，如-s=1或-s 1，则会自动识别-s之后的数字并将start_page赋值为1。

命令行flag的语法有如下三种形式：
> -flag // 只支持bool类型
> -flag=x
> -flag x // 只支持非bool类型

关于更多flag的使用请点击链接。[命令行参数解析](http://blog.studygolang.com/2013/02/%E6%A0%87%E5%87%86%E5%BA%93-%E5%91%BD%E4%BB%A4%E8%A1%8C%E5%8F%82%E6%95%B0%E8%A7%A3%E6%9E%90flag/)



---

##测试##

为了测试方便，page_len的默认值改为5，输出的每页行数为5

测试文件（在第3、8、12和15行的尾部各有一个换页符/f）：

![1](https://raw.githubusercontent.com/Leungchiho/ServiceComputing-CLI/master/%E6%88%AA%E5%9B%BE/in.png)

1.


> ./selpg -s=1 -e=1 in 

![1](https://raw.githubusercontent.com/Leungchiho/ServiceComputing-CLI/master/%E6%88%AA%E5%9B%BE/1.png)

该命令将把“input_file”的第1页写至标准输出（也就是屏幕），因为这里没有重定向或管道。page_len为5，所以第一页打印了5行line，中间空行是因为换页符的存在。

2.
> ./selpg -s=1 -e=1 <in 

![2](https://raw.githubusercontent.com/Leungchiho/ServiceComputing-CLI/master/%E6%88%AA%E5%9B%BE/2.png)

该命令与示例 1 所做的工作相同，但在本例中，selpg 读取标准输入，而标准输入已被 shell／内核重定向为来自“input_file”而不是显式命名的文件名参数。输入的第 1 页被写至屏幕，输出效果同上。

3.
> cat in | ./selpg -s=1 -e=1 

![3](https://raw.githubusercontent.com/Leungchiho/ServiceComputing-CLI/master/%E6%88%AA%E5%9B%BE/3.png)

“cat”的标准输出被 shell／内核重定向至 selpg 的标准输入。将第 1 页写至 selpg 的标准输出（屏幕）。

4.
> ./selpg -s=1 -e=1 in > out

![4](https://raw.githubusercontent.com/Leungchiho/ServiceComputing-CLI/master/%E6%88%AA%E5%9B%BE/4.1.png)

![4](https://raw.githubusercontent.com/Leungchiho/ServiceComputing-CLI/master/%E6%88%AA%E5%9B%BE/4.2.png)

selpg 将第 1页写至标准输出,标准输出被shell／内核重定向至“out”。输出没有在命令行显示，而是显示在输出文件中。

5.
> ./selpg -s=1 -e=-1 in 2> error

![4](https://raw.githubusercontent.com/Leungchiho/ServiceComputing-CLI/master/%E6%88%AA%E5%9B%BE/5.1.png)

![4](https://raw.githubusercontent.com/Leungchiho/ServiceComputing-CLI/master/%E6%88%AA%E5%9B%BE/5.2.png)

把end_page设为-1，发生错误，所有的错误消息被 shell／内核重定向至“error”。

6.
> ./selpg -s=1 -e=1 in > out 2> error

![4](https://raw.githubusercontent.com/Leungchiho/ServiceComputing-CLI/master/%E6%88%AA%E5%9B%BE/6.1.png)

![4](https://raw.githubusercontent.com/Leungchiho/ServiceComputing-CLI/master/%E6%88%AA%E5%9B%BE/6.2.png)

![4](https://raw.githubusercontent.com/Leungchiho/ServiceComputing-CLI/master/%E6%88%AA%E5%9B%BE/6.3.png)

selpg 将第1页写至标准输出，标准输出被重定向至“out”, 发生的“错误”写至error（实际上没有错误，“done！”只是为了演示）。

7.
> ./selpg -s=1 -e=1 in > out 2>/dev/null

![4](https://raw.githubusercontent.com/Leungchiho/ServiceComputing-CLI/master/%E6%88%AA%E5%9B%BE/7.1.png)

![4](https://raw.githubusercontent.com/Leungchiho/ServiceComputing-CLI/master/%E6%88%AA%E5%9B%BE/7.2.png)

selpg 将第 1页写至标准输出，标准输出被重定向至“out”，丢弃错误信息。可见成功输出到out文件，而error文件为空。

8.
> ./selpg -s=1 -e=1 in > /dev/null

![4](https://raw.githubusercontent.com/Leungchiho/ServiceComputing-CLI/master/%E6%88%AA%E5%9B%BE/8.1.png)

标准输出被丢弃

9.

> ./selpg -s=1 -e=1 in | wc

![4](https://raw.githubusercontent.com/Leungchiho/ServiceComputing-CLI/master/%E6%88%AA%E5%9B%BE/9.png)

selpg 的标准输出透明地被 shell／内核重定向，成为“wc”的标准输入，显示第一页中包含的行数、字数和字符数。

10.

> ./selpg -s=1 -e=1 in 2> error | wc

![4](https://raw.githubusercontent.com/Leungchiho/ServiceComputing-CLI/master/%E6%88%AA%E5%9B%BE/10.png)

![4](https://raw.githubusercontent.com/Leungchiho/ServiceComputing-CLI/master/%E6%88%AA%E5%9B%BE/6.2.png)

同上，唯一不同就是错误消息被写至“error_file”。

11.

> ./selpg -s=1 -e=1 -l=2 in

![4](https://raw.githubusercontent.com/Leungchiho/ServiceComputing-CLI/master/%E6%88%AA%E5%9B%BE/11.png)

输出的页长设置为2行，可见输出了一页为2行的信息。

12.

> ./selpg -s=1 -e=2 -f in

![4](https://raw.githubusercontent.com/Leungchiho/ServiceComputing-CLI/master/%E6%88%AA%E5%9B%BE/12.png)

页面由换页符定界。第1页到第2页被写至 selpg 的标准输出。

13.

> ./selpg -s=1 -e=2 -d=lp1 in

![4](https://raw.githubusercontent.com/Leungchiho/ServiceComputing-CLI/master/%E6%88%AA%E5%9B%BE/13.png)

由于缺少打印机，用cat命令代替lp命令，结果正确。

14.

> ./selpg -s=1 -e=2 in > out 2> error &

![4](https://raw.githubusercontent.com/Leungchiho/ServiceComputing-CLI/master/%E6%88%AA%E5%9B%BE/14.png)

selpg后台运行正常。










