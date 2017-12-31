#最后一次作业

---
**动机**

假设一个旅行社的信息系统由多个基本服务组成。这些服务可以使用不同的技术(JMS、EJB、WS、…)构建。为简单起见，我们假定服务可以通过HTTP方法调用(例如使用jax - rs客户端)来使用REST接口。我们还认为我们需要与之合作的基本服务是：


- 客户服务-提供关于旅行社客户的信息。
- 目的地服务—为经过身份验证的客户提供访问和推荐目的地的列表。
- 天气服务-为指定目的地提供天气预报。
- 报价服务-为客户提供价格计算，以到达推荐的目的地。

任务是为经过身份验证的用户创建一个公共可用特性，显示10个最近访问过的地方的列表，并显示10个推荐目的地的列表，其中包括用户的天气预报和价格计算。注意，有些请求(检索数据)依赖于先前请求的结果。例如，获取推荐目的地依赖于首先获取关于身份验证用户的信息。获取天气预报取决于目的地信息等等。一些请求之间的这种关系是问题的一个重要部分，也是一个您可以真正利用反应编程模型的领域。

![1](https://jersey.github.io/documentation/latest/images/rx-client-problem.png)

**一种天真的做法（同步实现）**
![2](https://jersey.github.io/documentation/latest/images/rx-client-sync-approach.png)

**更好的做法（异步实现）**
![3](https://jersey.github.io/documentation/latest/images/rx-client-async-approach.png)

**运行样例**

可以看到异步实现的时间比同步实现的时间短得多。客户端使用异步机制的主要目的是使请求之后的流程，不受服务器端请求处理过程的阻塞，以便节省执行时间。
![4](https://raw.githubusercontent.com/LeungChiHo/ServiceComputing/master/goroutine/1.png)