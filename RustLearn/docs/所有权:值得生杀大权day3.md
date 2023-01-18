# 《陈天 · Rust 编程第一课》学习笔记Day 3
今天我们一起来啃一下rust学习里的硬骨头--所有权、生命周期。

**这是Rust和其他编程语言的主要区别。也是Rust其他知识点的。**

从变量使用堆栈的行为开始，探究Rust设计所有权和生命周期的用意。

### 先看下变量在函数调用时发生了什么？

这段代码，在main函数里 初始化了一个动态数组data和 一个整形值 v.

然后把他们2个传递给find_pos函数，

在data中查找v 是否存在，如果存在返回下标，不存返回None。

find_pos最后一个语句 None 直接返回，不用return。

```rust
fn main() {
    let data = vec![10, 42, 9, 8];
    let v = 42;
    if let Some(pos) = find_pos(data, v) {
        println!("Found {} at {}", v, pos);
    }
}

fn find_pos(data: Vec<u32>, v: u32) -> Option<usize> {
    for (pos, item) in data.iter().enumerate() {
        if *item == v {
            return Some(pos);
        }
    }
    
    None
}
```
可以看到 data是动态数组，在编译期无法确定大小，所以实际上是在堆上申请的内存，在栈上有一个包含长度和容量的指针指向堆上的内存。（类似golang的slice，会内存逃逸到堆上）

调用find_pos的时候 data 和 v 作为参数传递，放在 find_pos 的参数区。

<img decoding="async" high="350" src="https://github.com/reallovelei/BookNote/blob/master/RustLearn/docs/day3%E6%89%80%E6%9C%89%E6%9D%831.png?raw=true"></img>

按大多数编程语言的做法，现在 data 就有2个引用了。且每当把data作为参数传递一次，就会多一个引用。

那么问题来了：堆上的内存什么时候释放？

C/C++：手动处理释放，占用码农心智。

Java:追踪式GC，定期扫描堆上数据。

Golang: 三色标记法+内存屏障。
但Java和Golang的GC都会带来**STW**的问题。

那么Rust是怎么解决的？
之前我们开发的时候，引用是一种随意、可隐式产生的行为。
比如C语言里到处乱飞的指针。
比如Java里随处可见的按引用传参。
而Rust决定**限制开发者随意引用**的行为。

先来看一个问题：谁真正拥有数据，值的生杀大权？这种权利是共享的还是独占的？

### 所有权和Move的语义
如何保持独占？ 要考虑的情况还是比较多的。可能造成这个变量的拥有者不唯一。
比如以下情况：
*   变量A被赋给 变量B。
*   变量A作为参数被传递给函数C。
*   作为返回值从函数D返回。

对于上面这些情况，Rust制定了一些规则：
1.  一个值只能被一个变量所拥有，这个变量被称为所有者。（一夫一妻制？）
2.  一个值同一时刻 只能有一个所有者。 不能有两个变量拥有相同的值。 函数返回 旧的所有者会把值的所有权 转交给新的所有者。（类似 离婚后 再与其他人结婚？）
3.  当所有者离开作用域，其拥有的值被丢弃。

这三条规则的核心就是保证单一所有权。

规则2里 讲的是所有权转移是Move语义，这个概念Rust是从C++那里借鉴的。

规则3里 提到的作用域（scope）是一个新概念，在Rust里主要是指{} 里的代码区，区分与其他语言一般是函数。

举个例子：
 在{}里声明的变量r1，离开这个{}后，作用域就结束了。
```rust
let s = String::from("Hello");
{
    let r1 = &s;
    println!("r1: {}", r1);
}
```
在这三条所有权规则的约束下，我们再来看之前data引用的问题是如何解决的。

<img decoding="async" high="350" src="https://github.com/reallovelei/BookNote/blob/master/RustLearn/docs/day3%E6%89%80%E6%9C%89%E6%9D%832.png?raw=true"></img>

原先main函数中的data，在调用find_pos()后，就失效了，编译器会保证main函数后的代码无法访问data这个变量，这样就确保了堆上的内存有且只有一个引用。

回到最开始的那段代码，调用find_pos的时候，在main里data 的所有权被转移到 find_pos里的data。

再来看一段代码，对所有权的理解

```rust

fn main() {
    let data = vec![1, 2, 3, 4];
    let data1 = data;
    println!("sum of data1: {}", sum(data1));
    println!("data1: {:?}", data1); // error1
    println!("sum of data: {}", sum(data)); // error2
}

fn sum(data: Vec<u32>) -> u32 {
    data.iter().fold(0, |acc, x| acc + x)
}
```

这时候我们运行 cargo run 
编译器会报错
<img decoding="async" high="350" src="https://github.com/reallovelei/BookNote/blob/master/RustLearn/docs/day3%E6%89%80%E6%9C%89%E6%9D%833.png?raw=true"></img>

从错误信息可以看出，不能使用已经移动过的变量。
这是因为我们在打印data1的时候，data1实际上已经在上面调用sum(data1)的时候。所有权已经转移(move)到sum()里了。所以后面就不能使用了。
那怎么才能 在sum()后面使用data1呢？
我现在知道的方式有2种：
1. 可以在传到sum()时候的 clone()一下。
```rust
println!("sum of data1: {}", sum(data1.clone()));
```
<img decoding="async" high="350" src="https://github.com/reallovelei/BookNote/blob/master/RustLearn/docs/day3%E6%89%80%E6%9C%89%E6%9D%834.png?raw=true"></img>

2. 传引用,但是sum函数的参数类型也需要修改。
```rust
println!("sum of data1: {}", sum(&data1));

fn sum(data: &Vec<u32>) -> u32 {
    data.iter().fold(0, |acc, x| acc + x)
}
```

可以看到，**所有权规则，解决了谁真正拥有数据的生杀大权问题，让堆上数据的多重引用不复存在，这是它最大的优势。**

