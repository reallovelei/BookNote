@[toc]

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

```javascript
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

```javascript
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
```javascript
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
console.log(proxy.nme);             // 识别出了拼写错误，并throws error
```

### 用has陷阱来隐藏属性

- 使用in操作符会使has陷阱被调用。它接收两个参数trapTarget和key
- 内部的Refelct.has()接收同样的2个参数。可以修改其返回的默认值。

```javascript
let target = {
    name: "target",
    value: 42
};

let proxy = new Proxy(target, {
    has(trapTarget, key) {

        if (key === "value") {
            return false;
        } else {
            return Reflect.has(trapTarget, key);
        }
    }
});


console.log("value" in proxy);      // false
console.log("name" in proxy);       // true
console.log("toString" in proxy);   // true
```
### deleteProperty 来阻止属性被删除

- delete操作符移除一个对象上的属性，并返回一个boolean标识操作是否成功。
- 严格模式下，试图删除一个nonconfigurable属性（不可改）会抛出错误；非严格模式下则返回false。

```javascript
let target = {
    name: 'target',
    value: 42
}

Object.defineProperty(target, 'name', { configurable: false})

const res = delete target.name // 如果严格模式会抛出错误
console.log(res)   // false
console.log('name' in target)  //true
```
- delete 操作对应的是deleteProperty 陷阱。它接收2个参数， trapTarget 和 key。

```javascript
let proxy = new Proxy(target, {
    deleteProperty(trapTarget, key) {

        if (key === "value") {
            return false;
        } else {
            return Reflect.deleteProperty(trapTarget, key);
        }
    }
});
```
### getPrototypeOf 和 setPrototypeof

- setPrototypeOf陷阱接收两个参数， trapTarget 和 proto。如果操作不成功，必须返回false。
- getPrototypeOf陷阱接收一个参数，就是trapTarget。必须返回一个对象或者null。否则会跑错误。
- 对改写这两个函数的限制保证了js语言中Object方法的一致性。

```javascript
//通过一直返回null隐藏了target的原型，同时不允许修改其原型
let target = {}
let proxy = new Proxy(target, {
    getPrototypeOf(trapTarget) {
        return null
    }
    setPrototyoeof(trapTarget, proto) {
        return false
    }
})

let targetProto = Object.getPrototypeOf(target)
let proxyProto = Object.getPrototypeOf(proxy)

console.log(targetProto === Object.prototype)   //true
console.log(proxyProto === Object.prototype)    //false
console.log(proxyProto)                         //null

//succeeds
Object.setPrototypeOf(target, {})

// throw error
Object.setPrototypeOf(proxy, {})   
```
- 如果要用默认行为，直接调用Reflect.getPrototypeOf/setPrototypeOf, 而不是调用Object.getPrototypeOf/setPrototypeOf。这样做是有原因的，两者的区别在于：
1. Object上的方法是高层的，而Refect上的方法是语言底层的。
2. Refeclt.get/setPrototypeOf()方法其实是把内置的[[GetPrototypeOf]]操作包装了一层，做了输入校验。
3. Object上的这两个方法其实也是调用内置的[[GetPrototypeOf]]操作，但在调用之前还干了些别的，并且检查了返回值，来决定后续行为。
4. 比如说，当传入的target不是对象时，Refeclt.getPrototypeOf()会抛出错误，而Object.getPrototypeOf会强制转换参数到对象，再继续操作。有兴趣的童鞋不妨传个数字进去试试~
5. 如果操作不成功，Reflect.setPrototypeOf()会返回一个布尔值来标识操作是否成功，而如果Object.setPrototypeOf()失败，则会直接抛错。前者的返回false 其实就会导致后者的抛错。
6. 如果操作成功，Reflect.setPrototypeOf()返回true，而Object.setPrototypeOf()返回第一个参数，即target。

### 对象扩展(Obejct Extensibility)的陷阱

- 在preventExtensions()和isExtensible()方法上，Object 和 Reflect 表现基本一致。只有在传入一个非对象的变量作为参数时，Object.isExtensible()会返回false，而Reflect.isExtensible()会抛错。这是因为底层的操作会有更严格的错误检查。

### 属性描述(Property Descriptor)陷阱

### 结论
Reflect是跟Object同级的一个js数据类型。一个类。
Reflect上的方法返回的值通常是布尔值，标识着操作的成功与否。Object上的方法如preventExtensions等则返回参数对象。
底层操作的验证会更加严格，Object上的操作则有更高的容错。
介绍了基本的api。本质是重写方法。meta编程。

参考文献：
1. [understanding es6: proxies and reflections](https://github.com/nzakas/understandinges6/blob/master/manuscript/12-Proxies-and-Reflection.md)