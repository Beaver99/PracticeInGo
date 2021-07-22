[toc]

# LC 100-200/2021 Jun-Sep 

* difficulty: medium

## stack

* 0071-simplify-path
  * :star:
  * 循环的索引如果要在**循环体后面用到**的话，那么它会多1次++/--
  * GO
    * `strings.Join`: `func Join(elems []string, seperator string) string` ，把字符串切片合并为一个字符串，其他语言也有
* 0114-flatten-binary-tree-to-linked-list
  * :star::star::star:
  * - [ ] 怎么用栈模拟递归--先/中/后序遍历
    - [ ] 尾递归？
    - [x] :dagger: O(1)空间写法；什么是前/后驱节点；值得重新一做
      - [x] * 妙处在于修改完树之后，先序遍历结果不变。
        * 但每回都要找前驱结点，浪费了时间
        * 注意考虑全部情况
      - [ ] 中序呢？
* 0143-reorder-list
  * :star::star::star:
  * 大问题化解为小问题，不要嫌麻烦
  * 非常考验链表基本功
* 0150-evaluate-reverse-polish-notation
  * :star:
  * 计算器/parser雏形

## queue

## tree

## binary search

## divide and conquer

## greedy

## union set

## dual pointer

## bfs/dfs

## heap

## tree-like array

## linked list

## dynamic programming

## array
