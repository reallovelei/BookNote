通过昨天的学习，我们对Rust的类型系统有了新的认识。
还学习了泛型数据结构和泛型函数来处理参数多态。
接下来，我们将会学习特设多态和子类型多态。

定义:
**特设多态**：包括运算符重载，是指同一种行为有很多不同的实现；
**子类型多态**：把子类型当成父类型使用，比如把Cat当成Animal用。

不过这两种多态都用到了trait。我们今天先来了解一下trait和特设多态。

### Trait
Rust里的Trait可以理解为接口，就是我们常说的面向接口编程的那个interface,它**定义了类型使用这个接口的行为**。
但是吧，看下面这个例子，还允许方法提供了默认实现。从这个角度看又觉得有点像抽象类。
```rust
pub trait Write {
    fn write(&mut self, buf: &[u8]) -> Result<usize>;
    fn flush(&mut self) -> Result<()>;
    fn write_vectored(&mut self, bufs: &[IoSlice<'_>]) -> Result<usize> { ... }
    fn is_write_vectored(&self) -> bool { ... }
    fn write_all(&mut self, buf: &[u8]) -> Result<()> { ... }
    fn write_all_vectored(&mut self, bufs: &mut [IoSlice<'_>]) -> Result<()> { ... }
    fn write_fmt(&mut self, fmt: Arguments<'_>) -> Result<()> { ... }
    fn by_ref(&mut self) -> &mut Self where Self: Sized { ... }
}
```
对于这个trait，我们只需要实现那2个没有提供默认实现的方法 -- write,flush 两个方法就可以了。

这个trait里的Self和self两个关键字的含义:
* Self:当前类型。如：File类型实现了Write,实现过程中用到的Self是指File。
* self:当第一个参数用时，等价于 self:Self,所以（&self == self:&Self）,（&mut self == self:&mut Self）。

我们再来看一坨代码，帮助理解一下。
```rust
use std::fmt;
use std::io::Write;

struct BufBuilder {
    buf: Vec<u8>,
}

impl BufBuilder {
    pub fn new() -> Self {
        Self {
            buf: Vec::with_capacity(1024),
        }
    }
}

// 实现 Debug trait，打印字符串
impl fmt::Debug for BufBuilder {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{}", String::from_utf8_lossy(&self.buf))
    }
}

impl Write for BufBuilder {
    fn write(&mut self, buf: &[u8]) -> std::io::Result<usize> {
        // 把 buf 添加到 BufBuilder 的尾部
        self.buf.extend_from_slice(buf);
        Ok(buf.len())
    }

    fn flush(&mut self) -> std::io::Result<()> {
        // 由于是在内存中操作，所以不需要 flush
        Ok(())
    }
}

fn main() {
    let mut buf = BufBuilder::new();
    buf.write_all(b"Hello world!").unwrap();
    println!("{:?}", buf);
}
```
这里我们用BufBuilder实现了Write trait的write和flush2个方法。这时候write_trait里的所有方法都可以调用了。
比如write_all()

我们看write_all的签名：
```rust
fn write_all(&mut self, buf: &[u$$8]) -> Result<()>
```
有2个入参：
&mut self： 例子中传的buf的可变引用。
&[u8]：传的是b"Hello world!"。

### trait 练习
做个练习巩固一下trait。

需求：解析字符串的某部分成某个类型。
定义一个方法parse()待实现的trait 。
```rust
pub trait Parse {
    fn parse(s: &str) -> Self;
}
```

试着为u8类型来实现这个trait。
如果入参为"123abc" 则解析成整数123。
如果入参为"abc" 则解析成0。
```rust
use regex::Regex;
pub trait Parse {
    fn parse(s: &str) -> Self;
}

impl Parse for u8 {
    fn parse(s: &str) -> Self {
        let re: Regex = Regex::new(r"^[0-9]+").unwrap();
        if let Some(captures) = re.captures(s) {
            // 取第一个 match，将其捕获的 digits 换成 u8
            captures
                .get(0)
                .map_or(0, |s| s.as_str().parse().unwrap_or(0))
        } else {
            0
        }
    }
}

#[test]
fn parse_should_work() {
    assert_eq!(u8::parse("123abcd"), 123);
    assert_eq!(u8::parse("1234abcd"), 0);
    assert_eq!(u8::parse("abcd"), 0);
}

fn main() {
    println!("result: {}", u8::parse("255 hello world"));
}
```
这样我们就按需求实现了这个trait，那如果要再实现一个f64类型的呢？
除了正则匹配的过程，其余部分都差不多，重复就是坏味道，这时候我们前面接触的泛型参数就可以帮我们解决这类问题。

通过对带有约束的泛型参数实现 trait，一份代码就实现了 u32 / f64 等类型的 Parse trait，非常精简。

```rust

use std::str::FromStr;

use regex::Regex;
pub trait Parse {
    fn parse(s: &str) -> Self;
}

// 我们约束 T 必须同时实现了 FromStr 和 Default
// 这样在使用的时候我们就可以用这两个 trait 的方法了
impl<T> Parse for T
where
    T: FromStr + Default,
{
    fn parse(s: &str) -> Self {
        let re: Regex = Regex::new(r"^[0-9]+(\.[0-9]+)?").unwrap();
        // 生成一个创建缺省值的闭包，这里主要是为了简化后续代码
        // Default::default() 返回的类型根据上下文能推导出来，是 Self
        // 而我们约定了 Self，也就是 T 需要实现 Default trait
        let d = || Default::default();
        if let Some(captures) = re.captures(s) {
            captures
                .get(0)
                .map_or(d(), |s| s.as_str().parse().unwrap_or(d()))
        } else {
            d()
        }
    }
}

#[test]
fn parse_should_work() {
    assert_eq!(u32::parse("123abcd"), 123);
    assert_eq!(u32::parse("123.45abcd"), 0);
    assert_eq!(f64::parse("123.45abcd"), 123.45);
    assert_eq!(f64::parse("abcd"), 0f64);
}

fn main() {
    println!("result: {}", u8::parse("255 hello world"));
}
```

### 让trait支持泛型
比如定义一个字符串拼接的接口。
让它可以和String进行拼接，也可以和&str进行拼接。

这时候trait就需要支持泛型了。
我们先来看一下标准库里的操作符是怎么做重载的？``
std::ops::Add 是用于做加法运算的trait。
```rust
pub trait Add<Rhs = Self> { // 这里就表示支持泛型了？
    type Output;
    #[must_use]
    fn add(self, rhs: Rhs) -> Self::Output;
}
```
这个 trait 有一个泛型参数 Rhs，代表加号右边的值，它被用在 add 方法的第二个参数位。这里 Rhs 默认是 Self，也就是说你用 Add trait ，如果不提供泛型参数，那么加号右值和左值都要是相同的类型。

我们来复数类型实现这个Add。
```rust

use std::ops::Add;

#[derive(Debug)]
struct Complex {
    real: f64,
    imagine: f64,
}

impl Complex {
    pub fn new(real: f64, imagine: f64) -> Self {
        Self { real, imagine }
    }
}

// 对 Complex 类型的实现
impl Add for Complex {
    type Output = Self;

    // 注意 add 第一个参数是 self，会移动所有权
    fn add(self, rhs: Self) -> Self::Output {
        let real = self.real + rhs.real;
        let imagine = self.imagine + rhs.imagine;
        Self::new(real, imagine)
    }
}

fn main() {
    let c1 = Complex::new(1.0, 1f64);
    let c2 = Complex::new(2 as f64, 3.0);
    println!("{:?}", c1 + c2);
    // c1、c2 已经被移动，所以下面这句无法编译
    // println!("{:?}", c1 + c2);
}
```
复数类型有实部和虚部，两个复数的实部相加，虚部相加，得到一个新的复数。注意 add 的第一个参数是 self，它会移动所有权，所以调用完两个复数 c1 + c2 后，根据所有权规则，它们就无法使用了。所以最后那行被注释了。

如果想解决我们可以对Complex的引用实现Add。
其他部分不变。
```rust
// ...

// 如果不想移动所有权，可以为 &Complex 实现 add，这样可以做 &c1 + &c2
impl Add for &Complex {
    // 注意返回值不应该是 Self 了，因为此时 Self 是 &Complex
    type Output = Complex;

    fn add(self, rhs: Self) -> Self::Output {
        let real = self.real + rhs.real;
        let imagine = self.imagine + rhs.imagine;
        Complex::new(real, imagine)
    }
}

fn main() {
    let c1 = Complex::new(1.0, 1f64);
    let c2 = Complex::new(2 as f64, 3.0);
    println!("{:?}", &c1 + &c2);
    println!("{:?}", c1 + c2);
}
```
这样最后一行就可以放出来了。因为这时候所有权不会转移了。

需求再变化一下，加深巩固。
现在设计一个复数和一个实数直接相加，相加的结果是实部和实数相加，虚部不变。
现在的完整实现应该是这样了。
```rust
use std::ops::Add;

#[derive(Debug)]
struct Complex {
    real: f64,
    imagine: f64,
}

impl Complex {
    pub fn new(real: f64, imagine: f64) -> Self {
        Self { real, imagine }
    }
}

// 对 Complex 类型的实现
impl Add for Complex {
    type Output = Self;

    // 注意 add 第一个参数是 self，会移动所有权
    fn add(self, rhs: Self) -> Self::Output {
        let real = self.real + rhs.real;
        let imagine = self.imagine + rhs.imagine;
        Self::new(real, imagine)
    }
}

// 如果不想移动所有权，可以为 &Complex 实现 add，这样可以做 &c1 + &c2
impl Add for &Complex {
    // 注意返回值不应该是 Self 了，因为此时 Self 是 &Complex
    type Output = Complex;

    fn add(self, rhs: Self) -> Self::Output {
        let real = self.real + rhs.real;
        let imagine = self.imagine + rhs.imagine;
        Complex::new(real, imagine)
    }
}

// 因为 Add<Rhs = Self> 是个泛型 trait，我们可以为 Complex 实现 Add<f64>
impl Add<f64> for &Complex {
    type Output = Complex;

    // rhs 现在是 f64 了
    fn add(self, rhs: f64) -> Self::Output {
        let real = self.real + rhs;
        Complex::new(real, self.imagine)
    }
}

fn main() {
    let c1 = Complex::new(1.0, 1f64);
    let c2 = Complex::new(2 as f64, 3.0);
    println!("{:?}", &c1 + &c2);
    println!("{:?}", &c1 + 5.0);
    println!("{:?}", c1 + c2);
    // c1, c2 已经被移动，所以下面这句无法编译
    // println!("{:?}", c1 + c2);
}
```

通过使用 Add，为 Complex 实现了和 f64 相加的方法。所以泛型 trait 可以让我们在需要的时候，对同一种类型的同一个 trait，有多个实现。

## 小结
今天我们一起认识了trait，以及如何让trait支持泛型。
我们明天接着$$




