# 组合模式

![图 6](../../image/4c80ad9e918e3408bc126edca72648be961f24ab6e1065b1bde897f9e89aef70.png)  

![图 23](../../image/cfae944587fd440ff729892a21714f8ea6d1a72cd6afabe0beefcff27e53da9e.png)  


这里讲的“组合模式”，**主要是用来处理树形结构数据。这里的“数据”，你可以简单理解为一组对象集合**。

>Compose objects into tree structure to represent part-whole hierarchies.Composite lets client treat individual objects and compositions of objects uniformly.

翻译成中文就是：将一组对象组织（Compose）成树形结构，以表示一种“部分 - 整体”的层次结构。组合让客户端（在很多设计模式书籍中，“客户端”代指代码的使用者。）可以统一单个对象和组合对象的处理逻辑。

假设我们有这样一个需求：设计一个类来表示文件系统中的目录，能方便地实现下面这些功能：
1. 动态地添加、删除某个目录下的子目录或文件；
2. 统计指定目录下的文件个数；
3. 统计指定目录下的文件总大小。

🌰：代码中，每个部门都会有员工，部门还会有子部门，对于这种嵌套结构，可以使用能多方式解决，树的前中后序遍历，但代码使用接口与实现的方式解决，解耦复杂函数的同时，提高代码的扩展性和可读性。
