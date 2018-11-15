[toc]

## 背景
在安卓，ios和较新的的windows平台上，操作系统对app有启动和运用的权限，这些平台会合理地为应用分配资源。但是由于历史原因，web上的app可以永远保持活跃状态。因此，如果有大量的页面在同时运行，关键的系统资源如内存，cpu，电池，和网络资源会被过度索取，造成很差的用户体验。
web平台很早就有了生命周期的概念，如`load`, `unload`, `visibilitychange`，但这些时间只能让开发者响应用户发起的生命周期变化。对那些性能很差的设备，浏览器需要提前知道这件事，以便更合理地回收和重新分配系统资源。实际上，现代浏览器已经在这么做了。但还有更多可以优化的，它们也还想做更多。问题在于，开发者并不清楚这些机制，所以浏览器还是得采取保守做法，或者冒着页面崩溃的危险。

页面生命周期API尝试通过以下方式去解决这个问题：
- 引入并标准化生命周期状态的概念
- 定义新的，系统启动的生命周期状态（new system-initialted states），允许浏览器限制非激活状态或者被隐藏的页面占用资源
- 建立新的api和事件，让web开发者可以响应生命周期状态的改变。

*chrome 68已经引入了这些页面的生命周期特性了。*

## 概览页面生命周期和状态

![Page Lifecycle states](https://developers.google.com/web/updates/images/2018/07/page-lifecycle-api-state-event-flow.png)

状态包括：
- active
- passive
- hidden
- frozen  浏览器停止执行可冻结的事件，比如js计时器和fetch的回调，都不会再进行了。这是一种节约资源的手段。
- terminated 页面一旦开始unload，并从内存中被浏览器清掉，就是被terminated(终结)了。
- discarded


事件包括（斜体为新出的api）：
- focus
- blur
- visibilitychange
- *freeze* 任务不会再执行
- *resume* 浏览器重新启动了一个冻结的页面
- pageshow
- pagehide
- beforeunload 仅仅用于提醒用户别忘了保存，不可滥用！
- unload 永远不要使用这个事件！

> frozen和discarded都是系统发起的状态，而不是用户发起的。如前所述，当今的浏览器可能会偶尔冻结或者丢弃了隐藏的tab，但开发者对此一无所知。所以在chrome68中，新引入了document上的freeze, resume这两个事件，以让开发者监听。

```js
document.addEventListener('freeze', (event) => {
  // The page is now frozen.
});

document.addEventListener('resume', (event) => {
  // The page has been unfrozen.
});

if (document.wasDiscarded) {
  // Page was previously discarded by the browser while in a hidden tab.
}
```

### 检测生命周期

在active, passive, hidden状态下，可以通过js代码来判断当前的生命周期状态。
```js
const getState = () => {
  if (document.visibilityState === 'hidden' ) {
    return 'hidden'
  } 
  else if (document.hasFocus()){
    return 'active'
  }
  return 'passive'
}
```
*frozen和terminated只能观测他们相对的freeze/pagehide事件*

```js
// Stores the initial state using the `getState()` function (defined above).
let state = getState();

// Accepts a next state and, if there's been a state change, logs the
// change to the console. It also updates the `state` value defined above.
const logStateChange = (nextState) => {
  const prevState = state;
  if (nextState !== prevState) {
    console.log(`State change: ${prevState} >>> ${nextState}`);
    state = nextState;
  }
};

// These lifecycle events can all use the same listener to observe state
// changes (they call the `getState()` function to determine the next state).
['pageshow', 'focus', 'blur', 'visibilitychange', 'resume'].forEach((type) => {
  window.addEventListener(type, () => logStateChange(getState()), {capture: true});
});

// The next two listeners, on the other hand, can determine the next
// state from the event itself.
window.addEventListener('freeze', () => {
  // In the freeze event, the next state is always frozen.
  logStateChange('frozen');
}, {capture: true});

window.addEventListener('pagehide', (event) => {
  if (event.persisted) {
    // If the event's persisted property is `true` the page is about
    // to enter the page navigation cache, which is also in the frozen state.
    logStateChange('frozen');
  } else {
    // If the event's persisted property is not `true` the page is
    // about to be unloaded.
    logStateChange('terminated');
  }
}, {capture: true});
```

这段代码已经很清晰了，注意这里是在捕获阶段监听的。为什么要这么做呢？

- 没有共同的触发对象。这些事件中, pagehide/pageshow 在window上触发，visibilitychange，freeze, resume在document上触发，focus和blur在对应的dom元素上触发
- 大部分事件都不会冒泡。
- 捕获阶段在target/冒泡阶段之前，所以在这里加入监听保证了他们会在其他可能取消这一事件的代码前执行。

### 跨浏览器差异

浏览器对上述API的实现还存在差异，例如：

- 一些浏览器在切换标签页的时候不会触发`blur`事件。这意味着一个页面可能直接由`active`状态变为了`hidden`状态。而跳过了`passive`状态。
- `freeze`和`resume`事件没有被完全支持。
- IE10- 不支持`visibilitychange`事件。
- 以前的浏览器，`visibilitychange`在`pagehide`之后触发，而chrome无视了`document`在`unload`的可见状态，先触发`visibilitychange`事件，再触发`pagehide`事件。

这一切都可以通过一个js库来解决：[PageLifecycle.js](https://github.com/GoogleChromeLabs/page-lifecycle)

### 开发者应该在什么`state`做什么事

- active: 响应用户输入行为的最重要时机。任何会阻碍主线程的非UI行为应该放到这之后来做。
- passive: 在`passive`状态用户没有跟页面交互，但页面仍然可见。这意味着UI的更新和动画仍然应该流畅进行，但更新的时机就没那么重要了。*页面从`active`变到`passive`也是去保存应用状态的最佳时机*。
- hidden: 这可能是开发者能可靠地检测到的最后一次状态改变了，因为用户可能直接关闭了浏览器或应用。诸如`beforeunload`, `pagehide`, `unload`事件，在这种情况下都不会被触发了。因此应该把`hidden state`当做用户`session`的结束点。换句话说，持久化那些未被保存的应用状态，并发送数据调查数据。停止UI更新和任何用户不希望在后台运行的任务。
- frozen: 可以被冻结的任务都会被暂停，直到页面解冻（可能永远都不会解冻了，嘤嘤嘤）。应该阻止任何的计时器，切断可能会影响其他开启的同源Tab的连接。具体来说，需要：
  - 关闭所有开启的`IndexedDB`的连接
  - 关闭所有开启的`BroadcasrChannel`的连接
  - 关闭所有激活态的`webRTC`连接
  - 关闭所有的`web Socket`连接
  - 释放所有可能拿着的`Web Locks`
  - 持久化动态的视图状态（如滚动高度）到`sessionStorage`或`IndexedDB`

  当页面从冻结态返回到`hidden`状态时，重连上述连接。
- terminated: 不做任何事，不做任何事，不做任何事。`beforeunload, pagehide, unload`都不能被可靠地监听到。
- discarded: 对开发者不可见。可以在一个被丢弃的页面重新加载的时候检测`document.wasDiscarded`。

### 避免使用废弃的生命周期API

- unload: 宜用`visibilitychange`事件取代来判断何时`session`终止，用`hidden`状态作为最后保存应用和用户数据的可靠之机。
- beforeunload: 和`unload`事件有同样的问题，会组织浏览器在`page navigation cache`中缓存页面。*仅当提示用户还有未保存的变化时调用，并且在保存后立即移除*。

正确操作：
```
const beforeUnloadListener = (event) => {
  event.preventDefault();
  return event.returnValue = 'Are you sure you want to exit?';
};

// A function that invokes a callback when the page has unsaved changes.
onPageHasUnsavedChanges(() => {
  addEventListener('beforeunload', beforeUnloadListener, {capture: true});
});

// A function that invokes a callback when the page's unsaved changes are resolved.
onAllChangesSaved(() => {
  removeEventListener('beforeunload', beforeUnloadListener, {capture: true});
});
```

*`pagelifecycle.js`库已经提供了`addUnsavedChanges()`和`removeUnsavedChanges()`方法*

### FAQs
1. 我的页面要在`hidden`时仍然工作，怎么阻止它被`frozen`或者`discarded`呢（比如音乐类APP）？

> chrome只会在确保安全时冻结或丢弃它。在有以下资源使用时则不会：
> - 播放音视频
> - 使用`WebRTC`
> - 更新表头或favicon
> - 弹`alert`
> - 发送`push notificatoins`

2. 什么是`page navigation cache`（页面导航缓存）？

这是一个通用名词，用来描述浏览器对页面导航的优化，让前进后退按钮更加快捷。`webkit`把它叫做`page cache`，火狐则成为`Back-Forwards Cache`。当导航离开时这些浏览器会冻结当前页面以节约cpu和电量，因此在前进后退再进入这个页面的时候，可以重新`resume`。添加`beforeunload`和`unload`事件监听器都会阻止浏览器所做的优化。

3. 为什么没有提到`laod`/`DOMContentLoaded`事件呢？

页面生命周期API要求状态是离散而且独立的。页面可能以`active`，`passvie`，`hidden`状态载入(load)，因此一个单独的`loading`状态毫无意义。并且二者都不能指示着页面生命周期的变化，所以与这些API无关。

4. 如果我不能在冻结态和终止态去运行异步的api，那我怎么把数据存到`IndexedDB`呢？

这确实是个问题。在`frozen`和`terminated`状态，可冻结的任务会被暂停，所以异步的回调都不能保证可靠。
未来会在`IDBTransaction`加入`commit()`方法，保证开发者可以执行不需要回调的**只写型**事务。也就是说，如果不需要读，`commit`方法可以在任务队列被暂停前完成。
目前，开发者还有这两种选择：
- 使用`session storage`，这是同步的，页面被丢弃也会持久化。
- 用`service worker`写入`IndexedDB`。可以在`freeze`/`pagehide`事件监听器上通过`postMessage()`给`service worker`发送数据，让后者来完成。但当存在内存压力的时候，不建议使用后者。

### 测试你的app的`frozen`和`discarded`状态

打开[chrome://discards/](chrome://discards/)来真正尝试一下冻结和丢弃打开的标签页是怎么回事儿吧~

同时还可以看看`document.wasDiscarded`的值是否跟预期一致。

### 总结

为了更合理地使用系统资源，开发者应善用页面周期状态。

另外，对浏览器而言，越多的开发者开始应用新的页面周期API，冻结和丢弃页面也会变得更安全可靠，从而节约内存，cpu，电量和网络资源。

最后，如果不想记住和手写这么多API，可以尝试`pagelifecycle.js`这个库。








