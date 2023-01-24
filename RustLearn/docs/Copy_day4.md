昨天我们发现一个问题：想要避免所有权转移后，变量不能访问的情况。我们需要手动clone(),这样操作会比较麻烦。

Rust提供了两种方案：
1. 如果不想转移变量的所有权，在Move语义外，Rust提供了**Copy语义**。如果一个数据结构实现了 Copy trait,那么他就会使用Copy语义。这样在赋值，传参的时候，值会自动按位拷贝(浅拷贝)。  

2. 如果你不希望值的所有权被转移，又无法使用Copy语义，那你可以”借用“数据。借用可以理解为引用，后面再详细看。

先来学习一下Copy语义
### Copy 语义和 Copy trait
当要移动一个值的时候，如果值的类型实现了Copy trait，就会自动使用Copy语义，进行拷贝，否则使用Move语义进行移动。

那么在Rust中，哪些类型实现了Copy trait呢？
可以跑一下这段代码。验证是否实现了Copy trait。
types_impl_copy_trait里的类型都是实现了Copy trait的。

```Rust
fn is_copy<T: Copy>() {}

fn types_impl_copy_trait() {
    is_copy::<bool>();
    is_copy::<char>();

    // all iXX and uXX, usize/isize, fXX implement Copy trait
    is_copy::<i8>();
    is_copy::<u64>();
    is_copy::<i64>();
    is_copy::<usize>();

    // function (actually a pointer) is Copy
    is_copy::<fn()>();

    // raw pointer is Copy
    is_copy::<*const String>();
    is_copy::<*mut String>();

    // immutable reference is Copy
    is_copy::<&[Vec<u8>]>();
    is_copy::<&String>();

    // array/tuple with values which is Copy is Copy
    is_copy::<[u8; 4]>();
    is_copy::<(&str, &str)>();
}

fn types_not_impl_copy_trait() {
    // unsized or dynamic sized type is not Copy
    is_copy::<str>();
    is_copy::<[u8]>();
    is_copy::<Vec<u8>>();
    is_copy::<String>();

    // mutable reference is not Copy
    is_copy::<&mut String>();

    // array / tuple with values that not Copy is not Copy
    is_copy::<[Vec<u8>; 4]>();
    is_copy::<(String, u32)>();
}

fn main() {
    types_impl_copy_trait();
    types_not_impl_copy_trait();
}
```
总结：
* 原生类型，包括函数，不可变引用和罗指针实现了Copy;
* 数组和元组，如果里面的元素实现了Copy，那么它们也就实现了Copy。
* 可变引用没有实现Copy。（<&mut String>）
* 非固定大小的结构，没有实现Copy。如：vec， hash。

核心点：** Rust 通过单一所有权来限制任意引用的行为**，就不难理解这些新概念背后的设计意义。

[官方文档](https://doc.rust-lang.org/std/marker/trait.Copy.html)也介绍实现了Copy trait的数据结构

![day4_copy.png](https://s2.loli.net/2023/01/19/SrfMQ5GzwuAkvX2.webp)

## 小结
回顾一下，这两天我们学习了Rust里的几个重要概念：
* 单一所有权模式：一个值只能被一个一个变量拥有，同一时刻只能有一个所有者，当所有者离开作用域，其拥有的值被丢弃，内存得到释放。
* Move语义：赋值，传参会导致值Move，所有权被转移，一旦所有权转移，之前的变量就不能访问了。
* Copy语义：今天刚学习的，讲人话，当移动时，支持Copy的就Copy，不支持的就Move了。

![day4_copy2.png](https://s2.loli.net/2023/01/19/X9eQ762aS5liwH3.png)


明天我们继续学习开始提到的第二个解决方案“借用”也叫引用。

