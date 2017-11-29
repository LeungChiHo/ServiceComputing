xorm
---
xorm是一个简单而强大的Go语言ORM库.通过它可以使数据库操作非常简便。xorm的目标并不是让你完全不去学习SQL，我们认为SQL并不会为ORM所替代，但是ORM将可以解决绝大部分的简单SQL需求。xorm支持两种风格的混用。

**1. 启动 mysql 作为主机服务**

现在默认已经安装好mysql，然后输入命令启动mysql
![1](https://raw.githubusercontent.com/LeungChiHo/ServiceComputing/master/CloudGo-database/screenshot/1.png)

**2. 启动 mysql client 访问服务器**
![2](https://raw.githubusercontent.com/LeungChiHo/ServiceComputing/master/CloudGo-database/screenshot/2.png)

**3. 创建数据库**
创建数据库test，用户表userinfo，关联用户信息表userdetail。
![3](https://raw.githubusercontent.com/LeungChiHo/ServiceComputing/master/CloudGo-database/screenshot/3.png)

**4. 测试 web 服务**

**添加数据**

用 curl POST 一些数据到网站。
![4](https://raw.githubusercontent.com/LeungChiHo/ServiceComputing/master/CloudGo-database/screenshot/4.png)

**查询数据**
查询刚刚上传的数据
![5](https://raw.githubusercontent.com/LeungChiHo/ServiceComputing/master/CloudGo-database/screenshot/5.png)

**错误输入**
![6](https://raw.githubusercontent.com/LeungChiHo/ServiceComputing/master/CloudGo-database/screenshot/6.png)

**5. 性能测试**

![7](https://raw.githubusercontent.com/LeungChiHo/ServiceComputing/master/CloudGo-database/screenshot/7.png)

下面是database/sql的实现方法
![8](https://raw.githubusercontent.com/LeungChiHo/ServiceComputing/master/CloudGo-database/screenshot/%E5%8E%9F%E6%9D%A5.png)

经过对比可以看到，xorm的性能比database/sql的慢了。

**6. 对比**


1. **编程效率**
xorm它可以使数据库操作非常简便，通过连写操作，可以通过很少的语句完成数据库操作。程序员不再关注数据库细节，专心在业务逻辑，程序员可以不懂数据库就可以开发系统。所以相对于database/sql，xorm编程效率高于database/sql。


2. **程序结构**
xorm更加简洁。

3. **服务性能**
如上一点的性能测试所示，xorm虽然程序上简化了，但内部使用了大量反射以及数据库检测，所以其性能却没有database/sql好。

