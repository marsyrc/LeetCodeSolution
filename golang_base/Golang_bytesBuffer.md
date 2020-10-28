## 1 bytes.Buffer定义

bytes.Buffer提供可扩容的字节缓冲区，实质是对切片的封装；结构中包含一个64字节的小切片，避免小内存分配：

```go
// A Buffer is a variable-sized buffer of bytes with Read and Write methods.
// The zero value for Buffer is an empty buffer ready to use.
type Buffer struct {
	buf       []byte   // contents are the bytes buf[off : len(buf)]
	off       int      // read at &buf[off], write at &buf[len(buf)]--->指示读指针
	bootstrap [64]byte // memory to hold first slice; helps small buffers avoid allocation.
	lastRead  readOp   // last read operation, so that Unread* can work correctly.
}
```

## 2 初始化bytes.Buffer的方法

```go
1） var buf bytes.Buffer ->定义一个空的字节缓冲区

2） func NewBuffer(buf []byte) *Buffer { return &Buffer{buf: buf} } -->将字节切片初始化为缓冲区

3） func NewBufferString(s string) *Buffer {return &Buffer{buf: []byte(s)}} -->将字符串初始化为缓冲区
```

## 3 提供的主要API函数

1）写字节流数据到缓冲区

```go
// Write appends the contents of p to the buffer, growing the buffer as
// needed. The return value n is the length of p; err is always nil. If the
// buffer becomes too large, Write will panic with ErrTooLarge.
func (b *Buffer) Write(p []byte) (n int, err error) {
	b.lastRead = opInvalid
	m := b.grow(len(p))
	return copy(b.buf[m:], p), nil
}
```

2）写字符串到缓冲区

```go
// WriteString appends the contents of s to the buffer, growing the buffer as
// needed. The return value n is the length of s; err is always nil. If the
// buffer becomes too large, WriteString will panic with ErrTooLarge.
func (b *Buffer) WriteString(s string) (n int, err error) {
	b.lastRead = opInvalid
	//返回写入的index
	m := b.grow(len(s))
	return copy(b.buf[m:], s), nil
}
```

3）从缓冲区中读取数据

```go
// Read reads the next len(p) bytes from the buffer or until the buffer
// is drained. The return value n is the number of bytes read. If the
// buffer has no data to return, err is io.EOF (unless len(p) is zero);
// otherwise it is nil.
func (b *Buffer) Read(p []byte) (n int, err error) {
	b.lastRead = opInvalid
	if b.off >= len(b.buf) {
		// Buffer is empty, reset to recover space.
		b.Truncate(0)
		if len(p) == 0 {
			return
		}
		return 0, io.EOF
	}
	n = copy(p, b.buf[b.off:])
	b.off += n
	if n > 0 {
		b.lastRead = opRead
	}
	return
}
```

4）从缓冲区中读取字符串，直到分隔符**delim** 位置

```go
// ReadString reads until the first occurrence of delim in the input,
// returning a string containing the data up to and including the delimiter.
// If ReadString encounters an error before finding a delimiter,
// it returns the data read before the error and the error itself (often io.EOF).
// ReadString returns err != nil if and only if the returned data does not end
// in delim.
func (b *Buffer) ReadString(delim byte) (line string, err error) {
	slice, err := b.readSlice(delim)
	return string(slice), err
}
```



5）将未被读取的字节数据返回

```sql
// Bytes returns a slice of length b.Len() holding the unread portion of the buffer.
// The slice is valid for use only until the next buffer modification (that is,
// only until the next call to a method like Read, Write, Reset, or Truncate).
// The slice aliases the buffer content at least until the next buffer modification,
// so immediate changes to the slice will affect the result of future reads.
func (b *Buffer) Bytes() []byte { return b.buf[b.off:] }
```

6）将未被读取的字节数据以字符串形式返回

```go
// String returns the contents of the unread portion of the buffer
// as a string. If the Buffer is a nil pointer, it returns "<nil>".
func (b *Buffer) String() string {
	if b == nil {
		// Special case, useful in debugging.
		return "<nil>"
	}
	return string(b.buf[b.off:])
}
```

7）返回缓冲区当前容量

```go
// Cap returns the capacity of the buffer's underlying byte slice, that is, the
// total space allocated for the buffer's data.
func (b *Buffer) Cap() int { return cap(b.buf) }
```

8）返回未被读取的字节数据大小

```go
// Len returns the number of bytes of the unread portion of the buffer;
// b.Len() == len(b.Bytes()).
func (b *Buffer) Len() int { return len(b.buf) - b.off }
```

## 4 bytes.Buffer自动扩容机制

 当向缓冲区写入数据时，首先会检查当前容量是否满足需求，如果不满足分三种情况处理：

1）当前内置缓冲区切片buf为空，且写入数据量小于bootstrap的大小（64字节），则bootstrap作为buf

2）当前未读数据长度+新写入数据长度**小于等于**缓冲区容量的1/2，则挪动数据（将未读的数据放到已读数据位置）

3）以上条件不满足，只能重新分配切片，容量设定为**2\*cap(b.buf) + n，**即两倍原来的缓冲区容量+写入数据量大小

```go
// grow grows the buffer to guarantee space for n more bytes.
// It returns the index where bytes should be written.
// If the buffer can't grow it will panic with ErrTooLarge.
func (b *Buffer) grow(n int) int {
	m := b.Len()
	// If buffer is empty, reset to recover space.
	if m == 0 && b.off != 0 {
		b.Truncate(0)
	}
	//如果需要的容量大于现在的容量--->
	if len(b.buf)+n > cap(b.buf) {
		var buf []byte
		//现有的预备64byte可以满足
		if b.buf == nil && n <= len(b.bootstrap) {
			buf = b.bootstrap[0:]
			//实际需要的小于本身切片容量
		} else if m+n <= cap(b.buf)/2 {
			// We can slide things down instead of allocating a new
			// slice. We only need m+n <= cap(b.buf) to slide, but
			// we instead let capacity get twice as large so we
			// don't spend all our time copying.
			copy(b.buf[:], b.buf[b.off:])
			buf = b.buf[:m]
		} else {
			// not enough space anywhere
			//不够，那就分配2倍加n的容量
			buf = makeSlice(2*cap(b.buf) + n)
			copy(buf, b.buf[b.off:])
		}
		b.buf = buf
		b.off = 0
	}
	b.buf = b.buf[0 : b.off+m+n]
	return b.off + m
}
```

## 5 bytes.Buffer的局限

  bytes.Buffer提供了对切片的初步封装，但也没做太多的事；对于已读的数据无法操作。  