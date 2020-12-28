package main

func main() {
	
}

/*
https://github.com/golang/go/wiki/Performance

1.Combine objects into larger objects. For example, replace *bytes.Buffer struct member with bytes.
Buffer (you can preallocate buffer for writing by calling bytes.Buffer.Grow later). This will reduce number
of memory allocations (faster) and also reduce pressure on garbage collector (faster garbage collections).

4.If possible use smaller data types. For example, use int8 instead of int.

5.Objects that do not contain any pointers (note that strings, slices, maps and chans contain implicit pointers),
are not scanned by garbage collector. For example, a 1GB byte slice virtually does not affect garbage collection time.
So if you remove pointers from actively used objects, it can positively impact garbage collection time.
Some possibilities are: replace pointers with indices, split object into two parts one of which does not contain pointers.
不包含任何指针的对象（请注意，字符串，切片，映射和通道包含隐式指针）不会被垃圾收集器扫描。
例如，一个1GB的字节片实际上不会影响垃圾回收时间。 因此，如果从活动使用的对象中删除指针，则会对垃圾回收时间产生积极影响。
一些可能性是：用索引替换指针，将对象分为两部分，其中之一不包含指针。

string是一个只读的切片类型，其中的stringHeader包含一个指针和len(但没有cap)。

6.Use freelists to reuse transient objects and reduce number of allocations. Standard library contains
sync.Pool type that allows to reuse the same object several times in between garbage collections.
However, be aware that, as any manual memory management scheme, incorrect use of sync.Pool can lead to use-after-free bugs.
使用空闲列表重用临时对象并减少分配数量。 标准库包含sync.Pool类型，该类型允许在垃圾回收之间多次重用同一对象。
但是，请注意，与任何手动内存管理方案一样，对sync.Pool的错误使用可能会导致使用后释放错误。

*/