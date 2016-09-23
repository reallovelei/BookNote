## 结合walu 扩展函数

## 第九节
```
/**
 * ld  资源销毁时的回调函数
 * pld 用在一个类似于长链接类型的资源上回调函数,在MSHUTDOWN阶段被内核调用。
 * type_name 资源类别的名称
 *
 */
ZEND_API int zend_register_list_destructors_ex(rsrc_dtor_func_t ld, rsrc_dtor_func_t pld, char *type_name, int module_number);

   //每一个资源都是通过它来实现的。
    typedef struct _zend_rsrc_list_entry
    {
        void *ptr;   // 指向资源最终实现的指针, 如文件句柄/数据库连接/
        int type;    // 类型标记
        int refcount;  // 引用计数
    }zend_rsrc_list_entry;
```

## 方法
emalloc pemalloc
ZEND_MM_CHECK(chunk->heap == heap, "zend_mm_heap corrupted");   如果args1条件不成立 则报错信息为args2
