### 上期问题

上一期的最后，大路提出了一个问题，那就是sds一共可以存储多大的字节呢？

> 在目前版本的 Redis 中， SDS_MAX_PREALLOC 的值为 1024 * 1024 ， 也就是说， 当大小小于 1MB 的字符串执行追加操作时， sdsMakeRoomFor 就为它们分配多于所需大小一倍的空间； 当字符串的大小大于 1MB ， 那么 sdsMakeRoomFor 就为它们额外多分配 1MB 的空间。

从这里可以看到，sds的默认存储最大值是1MB，但它实际上是可以通过不断追加长度，多分配新的内存空间的。因此，理论上，内存中有多少可用空间，redis就可以存储多大的字符。

你看SDS里有一个len字段 可以表示buf已经使用的字节数。那么这个len的类型是int的。 所以看一下len的最大值可以是多少。是不是答案就浮出水面了呢！~

BUT,  在我翻阅redis源码的时候发现 src/t_string.c 文件里有一个检查字符串长度的函数。
```C
static int checkStringLength(client *c, long long size) {
    if (size > 512*1024*1024) {
        addReplyError(c,"string exceeds maximum allowed size (512MB)");
        return C_ERR;
    }
    return C_OK;
}
```
O Ho，redis作责居然限制了512M的大小。[官网上也是这么说的](https://redis.io/topics/data-types)
但是好奇心驱使着我继续发问，哪里用到这个函数呢 发现只在2个命令里用到 一个是append  一个是setrange.
原本我以为是应该是 sdsnew 创建一个新的sds里会检测长度，但是却让我失望了。
```C
/* Create a new sds string starting from a null terminated C string. */
sds sdsnew(const char *init) {
    size_t initlen = (init == NULL) ? 0 : strlen(init);
    return sdsnewlen(init, initlen);
}
```
于是乎我就一脸懵逼的继续去看代码。我承认我暂时找不到答案了。
我去做实验，首先我先验证了一下 append  和 setrange2个命令。
我生成了一个512M的文本，
然后直接set发现可以set. 没毛病。
开始测试：
```C
127.0.0.1:6379> append key a
(error) ERR string exceeds maximum allowed size (512MB)
127.0.0.1:6379> setrange key 536870912 a
(error) ERR string exceeds maximum allowed size (512MB)
```
可以看到 这2个命令和代码里的错误提示是一致的。
那么set 一个大于512m的字符串呢。我用了PHP 和 golang2个语言发现都不行。
php报这个错：
```
Redis::set(): send of 8192 bytes failed with errno=32 Broken pipe in xxxx
```
golang报这个错:
```
redis set failed: write tcp 127.0.0.1:55045->127.0.0.1:6379: write: broken pipe.
```
这两个错都是管道相关的错误。我在Redis源码没有找到相关错误。我以为是client的错误。
于是我想直接用命令行 redis-cli 还是一样报错。 
```
/redis-cli -h 127.0.0.1 < /Users/Ben/work/test/redisexp/512 
Error: Protocol wrong type for socket
```
**至此我的问题还是没有得到有效解决。望各位大神指点一二。**


<hr />

### 双端链表

今天我们来一起学习redis的另一种基本的内置数据结构：双端链表。提到链表，它的用处可多了，redis的list列表结构，其底层实现之一就是双端链表。此外，事务、订阅/发送等功能，都离不开它。

> 数组和链表的区别？
我们我们知道遍历数组和链表的时间复杂度是O(n)。
但是实际使用来说，数组的速度是要优于链表的，
1.数组元素的类型统一
2.数组的内存空间是一块连续空间，而链表的内存空间是散列的。
插入和删除的操作比较：
1.在数组中间插入（或删除）一个元素，那么这个元素后的所有元素的内存地址都要往后（前）移动（数组的内存地址是连续的），操作最后一个元素时才比较快， 而链表不需要改变内存的地址，只需要修改节点的指针即可（包括指针指向，节点值）。
2.因为定义数组时所占用的空间大小都是固定的，如果存储满了，无法扩展，只能新建一个更大空间的数组。链表的扩展性较好，

> 那另一种底层实现是什么呢？
> 也是链表，叫做压缩列表。其实，因为压缩列表占用的内存更小，在创建新列表的时候，会优先使用压缩列表。在之后的章节中，我们还会再进一步介绍它。

双端链表由listNode 和 List 两个数据结构组成。

- listNode代表节点。包括三个指针，prev，前一个节点，next，下一个节点，value，它存储的值。因此，可以双向遍历,其实这个结构体本身就完成双向链表了。但是redis为了操作方便又为我们封装了list结构体。
- list结构体为链表提供了表头/表尾指针，因此，对表头和表尾的插入的复杂度都是O(1)。lpush, rpush 都很便捷。还定义了链表的长度，len属性，所以求长度的复杂度也是O(1)。此外还定义了一些操作函数。
```c
typedef struct listNode {
    struct listNode *prev;
    struct listNode *next;
    void *value;
} listNode;

typedef struct list {
    listNode *head; // 表头指针
    listNode *tail; // 表尾指针

    unsigned long len;  // 节点数量

    // 复制函数
    void *(*dup)(void *ptr);
    // 释放函数
    void (*free)(void *ptr);
    // 比对函数
    int (*match)(void *ptr, void *key);

} list;
```
**注意， listNode 的 value 属性的类型是 void * ，说明这个双端链表对节点所保存的值的类型不做限制。(这也是一种多态的体现)**

> 一共有多少个节点？
> 你仔细读一下代码啊。unsigned long 代表 32位，4个字节，数字范围在0 -- 2^32-1

> 为什么除了头尾指针和length，还有函数指针？
> 这是一个很好的问题。对于不同类型的值，有时候需要不同的函数来处理这些值，因此， list 类型保留了三个函数指针 —— dup 、 free 和 match ，分别用于处理值的复制、释放和对比匹配。在对节点的值进行处理时，如果有给定这些函数，就会调用这些函数。举个例子：当删除一个 listNode 时，如果包含这个节点的 list 的 list->free 函数不为空，就会先调用删除函数 list->free(listNode->value) 来清空节点的值，再执行余下的删除操作（比如说，释放节点）

？？？这块没懂。你再解释下。我大概的理解是怕有listnode本身有副作用，会占用内存啊，或者什么的 ，不能直接简单地释放。

### 迭代器

链表的缺陷，或者说是特点，就在于删除和插入给定的节点，复杂度都是O(1)，但根据Key查找一个节点，就需要遍历操作，复杂度是O(n)级别的。

为了方便遍历操作，Redis 为双端链表实现了一个迭代器，可以从两个方向对双端链表进行迭代。

```c
typedef struct listIter {

    // 下一节点
    listNode *next;

    // 迭代方向
    int direction;

} listIter;
```
它包括next指针，指向下一个节点。还有一个Int类型的值direction，记录迭代应该从那里开始。
如果值为 adlist.h/AL_START_HEAD ，那么迭代器沿着节点的 next 指针前进，执行从表头到表尾的迭代；
如果值为 adlist.h/AL_START_TAIL ，那么迭代器沿着节点的 prev 指针前进，执行从表尾到表头的迭代

*感觉这里说迭代不如说遍历来得顺。迭代和遍历有区别吗？？？*

*是不是每一期来个小节比较好？？？*

总结一下：链表特性:1.双端，2.无环，3.带表头和表尾指针，4.带链表长度计数器，5.多态
