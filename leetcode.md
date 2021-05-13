

# leetcode

[toc]

## star rank

:new_moon_with_face: wasted my time

:star: learned something, not much

:star::star: worth reviewing and rewriting with other language

:star::star::star: ​classic and worth memorizing 



## Binary Tree

* 110
  * :star::star::star:
  * 递归-减枝法
  * - [ ] 自底向上
* 101
  * :star::star:
* 617
  * :star::star:
  * - [ ] 高效做法？
    - [ ] 

* 814
  * :star:
  * 谋定而后动

## Link List ##

* 链表操作类型题目 最主要有三点：
  * 想清楚**普通**情况下的操作，比如反转链表要假定前n个已经反转好了，要反转下一个，而不是从第一个开始想，因为太特殊了。
  * 想清楚每一个指针的**语义**，别瞎起名
  * 想清楚初始情况/极端情况



* 链表 如果需要 **双向移动/头对齐/尾对齐**， 可以考虑把**链表放到 数组/队列/栈** 里面



* nc-066
  * :star:
* 



## Array ##



* nc-037 合并区间
  * :star::star:
  * 预先排序可以得到很好的性质
  * 需要遍历数组时，当数组本身长度变化时，可以把指针的变化放进loop里面，外面什么也不放
* nc-022 合并有序数组
  * :star::star::star:
  * if if 不等于 if elseif  ，两者有逻辑上的不同
  * 数组题时刻记住**下标含义**
  * 用索引之前务必保证**索引有效**，for内改index尤其注意
* nc-   括号表达式求值
  * :star::star::star:
  * - [ ] for内改index很丑陋，有更优解吗？
  * 当if if可能多次匹配时，它们的**顺序**很重要
  * **栈的精髓**是有用信息在顶层的一个或几个







## 递归 ##



* nc-008 二叉树指定和
  * :star::star:
  * 用指针之前务必保证**指针有效非空**
  * - [ ] 用标准bfs做？
* nc-005 二叉树路径和
  * :star::star::star:
  * 递归的**返回结果**一定要设计的方便利用递归特性，加上**递归方程**，加上**奠基** == 秒杀



## DFS ##

* 