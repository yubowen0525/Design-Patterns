# 观察者模式 / 发布订阅模式

今天，我们学习第一个行为型设计模式，也是在实际的开发中用得比较多的一种模式：观察者模式。**根据应用场景的不同，观察者模式会对应不同的代码实现方式：有同步阻塞的实现方式，也有异步非阻塞的实现方式；有进程内的实现方式，也有跨进程的实现方式**。

> Define a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically.

翻译成中文就是：在对象之间定义一个一对多的依赖，当一个对象状态改变的时候，所有依赖的对象都会自动收到通知。


![图 1](../../image/bfc0310b53d6c47fdf3c6b8e9340c0a4d72c2a5e30955ea128c9da6ecf49f912.png)  

---
UML: 同步模型
![图 1](../../image/36e03f7a05606be39ac6a1c8a54b54c9b294231ad575f7e5ed805eecd6581567.png)  

UML: 异步模型
![图 2](../../image/686851286f183a84c6830a7dc7f20c1a5be1e1254af3f73eaf94ff34c8876b29.png)  
