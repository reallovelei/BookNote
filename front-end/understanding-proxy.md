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

参考文献：
1. [understanding es6: proxies and reflections](https://github.com/nzakas/understandinges6/blob/master/manuscript/12-Proxies-and-Reflection.md)