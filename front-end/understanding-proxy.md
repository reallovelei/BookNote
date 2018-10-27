### 什么是proxy和refection
- 通过 new Proxy()可 生成一个proxy来代替目标对象（target object）来使用。它等于目标对象的虚拟化，对于使用了该proxy的方法而言，二者看起来是一样的。通过proxy可以一窥原来只能由js引擎完成的底层操作。
- reflection API 是以Reflect对象为代表的的一组方法，为同级的底层操作提供proxy可以重写的默认行为。每个proxy trap都有Reflect方法。

Proxy Trap | Overrides the Behavior Of | Default Behavior | 
|--------------------------|---------------------------|------------------|
|get | Reading a property value | Reflect.get() |
|set | Writing to a property | Reflect.set() |
|has | The in operator | Reflect.has() |
|deleteProperty | The delete operator | Reflect.deleteProperty() |
|getPrototypeOf | Object.getPrototypeOf() | Reflect.getPrototypeOf() |
|setPrototypeOf | Object.setPrototypeOf() | Reflect.setPrototypeOf() |
|isExtensible | Object.isExtensible() | Reflect.isExtensible() |
|preventExtensions | Object.preventExtensions(）|Reflect.preventExtensions() |
|getOwnPropertyDescriptor| Object.getOwnPropertyDescriptor() | Reflect.getOwnPropertyDescriptor() |
|defineProperty | Object.defineProperty() | Reflect.defineProperty | 
|ownKeys | Object.keys, Object.getOwnPropertyNames(), Object.getOwnPropertySymbols() | Reflect.ownKey() |
|apply | Calling a function | Reflect.apply() |
|construct | Calling a function with new | Reflect.construct() |

### 生成一个新的简单proxy

- 传入两个参数，目标对象和句柄 (target and handler)。
- 句柄（handler），就是一个定义了一个或多个“陷阱”（trap）的对象。
- proxy在做没有定义陷阱的其他操作时，使用默认的行为。此时二者的表现是相同的，操作proxy就等于操作target。

```
let target = {};

let proxy = new Proxy(target, {});

proxy.name = "proxy";
console.log(proxy.name);        // "proxy"
console.log(target.name);       // "proxy"

target.name = "target";
console.log(proxy.name);        // "target"
console.log(target.name);       // "target"
```
### 用set陷阱来验证属性

set 陷阱接受4个参数：
1. trapTarget
2. key 对象的key，字符或者symbol，重写的就是这个属性啦
3. value 赋予这个属性的值
4. receiver 操作在哪个对象上发生，receiver就是哪个对象，通常就是proxy。

```
let target = {
    name: "target"
};

let proxy = new Proxy(target, {
    set(trapTarget, key, value, receiver) {
        // ignore existing properties so as not to affect them
        if (!trapTarget.hasOwnProperty(key)) {
            if (isNaN(value)) {
                throw new TypeError("Property must be a number.");
            }
        }
        console.table(Reflect)
        // add the property
        return Reflect.set(trapTarget, key, value, receiver);
    }
});

// adding a new property
proxy.count = 1;   // 这时trapTarget就是target，key等于count，value 等于1， receiver 是 proxy本身
console.log(proxy.count);       // 1
console.log(target.count);      // 1

// you can assign to name because it exists on target already
proxy.name = "proxy";
console.log(proxy.name);        // "proxy"
console.log(target.name);       // "proxy"

// throws an error
proxy.anotherName = "proxy";
```
- new Proxy 里，传入的第二个参数即为handler，这里定义了set方法，对应的内部操作是Reflect.set()，同时也是默认操作。
- set proxy trap 和 Reflect.set() 接收同样的四个参数
- Reflect.set() 返回一个boolean值标识set操作是否成功，因此，如果set了属性，trap会返回true，否则返回false。

### 用get陷阱来验证对象结构（object shape）
 
 > object shape: 一个对象上可用的属性和方法的集合。

 与很多其他语言不通，js奇葩的一点在于，获取某个不存在的属性时，不会报错，而是会返回undefined。在大型项目中，经常由于拼写错误等原因造成这种情况。那么，如何用Proxy的get方法来避免这一点呢？

使用Object.preventExtensions(), Object.seal(), Object.freeze() 等方法，可以强迫一个对象保持它原有的属性和方法。现在要使每次试图获取对象上不存在的属性时抛出错误。在读取属性时，会走proxy。.get()接收3个参数。

1. trapTarget
2. key: 属性的键。一个字符串或者symbol。
3. receiver

比起上面的set，少了一个value。Reflect.get()方法同样接收这3个参数，并返回属性的默认值。
```
let proxy = new Proxy({}, {
    get(trapTartet, key, receiver) {
        if(!(key in receiver)) {
            throw new TypeError(`property ${key} doesn't exist`)
        }
        return Reflect.get(trapTarget, key, recevier)
    }
})

// adding a property still works
proxy.name = "proxy";
console.log(proxy.name);            // "proxy"

// nonexistent properties throw an error
console.log(proxy.nme);             // throws error
```


参考文献：
1. [understanding es6: proxies and reflections](https://github.com/nzakas/understandinges6/blob/master/manuscript/12-Proxies-and-Reflection.md)