## sync.Mutex

Go中使用sync.Mutex类型实现mutex(排他锁、互斥锁)。在源代码的sync/mutex.go文件中，有如下定义：

```
// A Mutex is a mutual exclusion lock.
// The zero value for a Mutex is an unlocked mutex.
//
// A Mutex must not be copied after first use.
type Mutex struct {
	state int32
	sema uint32
}
```

这没有任何非凡的地方。**和mutex相关的所有事情都是通过sync.Mutex类型的两个方法sync.Lock()和sync.Unlock()函数来完成的，前者用于获取sync.Mutex锁，后者用于释放sync.Mutex锁**。sync.Mutex一旦被锁住，其它的Lock()操作就无法再获取它的锁，只有通过Unlock()释放锁之后才能通过Lock()继续获取锁。

也就是说，**已有的锁会导致其它申请Lock()操作的goroutine被阻塞，且只有在Unlock()的时候才会解除阻塞**。

另外需要注意，**sync.Mutex不区分读写锁，只有Lock()与Lock()之间才会导致阻塞的情况**，如果在一个地方Lock()，在另一个地方不Lock()而是直接修改或访问共享数据，这对于sync.Mutex类型来说是允许的，因为mutex不会和goroutine进行关联。如果想要区分读、写锁，可以使用sync.RWMutex类型，见后文。

**在Lock()和Unlock()之间的代码段称为资源的临界区(critical section)，在这一区间内的代码是严格被Lock()保护的，是线程安全的，任何一个时间点都只能有一个goroutine执行这段区间的代码**。

以下是使用sync.Mutex的一个示例，稍后是非常详细的分析过程。

```
package main

import (
	"fmt"
	"sync"
	"time"
)

// 共享变量
var (
	m  sync.Mutex
	v1 int
)

// 修改共享变量
// 在Lock()和Unlock()之间的代码部分是临界区
func change(i int) {
	m.Lock()
	time.Sleep(time.Second)
	v1 = v1 + 1
	if v1%10 == 0 {
		v1 = v1 - 10*i
	}
	m.Unlock()
}

// 访问共享变量
// 在Lock()和Unlock()之间的代码部分是是临界区
func read() int {
	m.Lock()
	a := v1
	m.Unlock()
	return a
}

func main() {
	var numGR = 21
	var wg sync.WaitGroup

	fmt.Printf("%d", read())

	// 循环创建numGR个goroutine
	// 每个goroutine都执行change()、read()
	// 每个change()和read()都会持有锁
	for i := 0; i < numGR; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			change(i)
			fmt.Printf(" -> %d", read())
		}(i)
	}

	wg.Wait()
}
```

第一次执行结果：

```
0 -> 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 -> -100 -> -99
-> -98 -> -97 -> -96 -> -95 -> -94 -> -93 -> -92 -> -91 -> -260 -> -259
```

第二次执行结果：注意其中的-74和-72之间跨了一个数

```
0 -> 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 -> -80 -> -79 
-> -78 -> -77 -> -76 -> -75 -> -74 -> -72 -> -71 -> -230 -> -229 -> -229
```

上面的示例中，change()、read()都会申请锁，并在准备执行完函数时释放锁，它们如何修改数据、访问数据本文不多做解释。需要详细解释的是main()中的for循环部分。

在for循环中，会不断激活新的goroutine(共21个)执行匿名函数，在每个匿名函数中都会执行change()和read()，意味着每个goroutine都会申请两次锁、释放两次锁，且for循环中没有任何Sleep延迟，这21个goroutine几乎是一瞬间同时激活的。

但由于change()和read()中都申请锁，对于这21个goroutine将要分别执行的42个critical  section，Lock()保证了在某一时间点只有其中一个goroutine能访问其中一个critical  section。当释放了一个critical section，**其它的Lock()将争夺互斥锁，也就是所谓的竞争现象(race condition)**。因为竞争的存在，**这42个critical section被访问的顺序是随机的**，完全无法保证哪个critical section先被访问。

对于前9个被调度到的goroutine，无论是哪个goroutine取得这9个change(i)中的critical section，都只是对共享变量v1做加1运算，但当第10个goroutine被调度时，由于v1加1之后得到10，它满足if条件，会执行`v1 = v1 - i*10`，但这个i可能是任意0到numGR之间的值(因为无法保证并发的goroutine的调度顺序)，这使得v1的值从第10个goroutine开始出现随机性。但从第10到第19个goroutine被调度的过程中，也只是对共享变量v1做加1运算，这些值是可以根据第10个数推断出来的，到第20个goroutine，又再次随机。依此类推。

此外，**每个goroutine中的read()也都会参与锁竞争，所以并不能保证每次change(i)之后会随之执行到read()**，可能goroutine 1的change()执行完后，会跳转到goroutine 3的change()上，这样一来，goroutine  1的read()就无法读取到goroutine  1所修改的v1值，而是访问到其它goroutine中修改后的值。所以，前面的第二次执行结果中出现了一次数据跨越。只不过执行完change()后立即执行read()的几率比较大，所以多数时候输出的数据都是连续的。

总而言之，**Mutex保证了每个critical section安全，某一时间点只有一个goroutine访问到这部分，但也因此而出现了随机性**。

如果Lock()后忘记了Unlock()，将会永久阻塞而出现死锁。如果



### 适合sync.Mutex的数据类型

其实，对于内置类型的共享变量来说，使用sync.Mutex和Lock()、Unlock()来保护也是不合理的，因为它们自身不包含Mutex属性。真正合理的共享变量是那些包含Mutex属性的struct类型。例如：

```
type mytype struct {
	m   sync.Mutex
	var int
}

x := new(mytype)
```

这时只要想保护var变量，就先x.m.Lock()，操作完var后，再x.m.Unlock()。这样就能保证x中的var字段变量一定是被保护的。



## sync.RWMutex

Go中使用sync.RWMutex类型实现读写互斥锁rwmutex。在源代码的sync/rwmutex.go文件中，有如下定义：

```
// A RWMutex is a reader/writer mutual exclusion lock.
// The lock can be held by an arbitrary number of readers or a single writer.
// The zero value for a RWMutex is an unlocked mutex.
//
// A RWMutex must not be copied after first use.
//
// If a goroutine holds a RWMutex for reading and another goroutine might
// call Lock, no goroutine should expect to be able to acquire a read lock
// until the initial read lock is released. In particular, this prohibits
// recursive read locking. This is to ensure that the lock eventually becomes
// available; a blocked Lock call excludes new readers from acquiring the
// lock.
type RWMutex struct {
	w           Mutex  // held if there are pending writers
	writerSem   uint32 // 写锁需要等待读锁释放的信号量
	readerSem   uint32 // 读锁需要等待写锁释放的信号量
	readerCount int32  // 读锁后面挂起了多少个写锁申请
	readerWait  int32  // 已释放了多少个读锁
}
```

上面的注释和源代码说明了几点：

1. **RWMutex是基于Mutex的，在Mutex的基础之上增加了读、写的信号量，并使用了类似引用计数的读锁数量**

2. 读锁与读锁兼容，读锁与写锁互斥，写锁与写锁互斥，只有在锁释放后才可以继续申请互斥的锁

   ：

   - **可以同时申请多个读锁**
   - **有读锁时申请写锁将阻塞，有写锁时申请读锁将阻塞**
   - **只要有写锁，后续申请读锁和写锁都将阻塞**

此类型有几个锁和解锁的方法：

```
func (rw *RWMutex) Lock()
func (rw *RWMutex) RLock()
func (rw *RWMutex) RLocker() Locker
func (rw *RWMutex) RUnlock()
func (rw *RWMutex) Unlock()
```

其中：

1. **Lock()和Unlock()用于申请和释放写锁**
2. RLock()和RUnlock()用于申请和释放读锁
   - **一次RUnlock()操作只是对读锁数量减1，即减少一次读锁的引用计数**
3. **如果不存在写锁，则Unlock()引发panic，如果不存在读锁，则RUnlock()引发panic**
4. **RLocker()用于返回一个实现了Lock()和Unlock()方法的Locker接口**

此外，无论是Mutex还是RWMutex都不会和goroutine进行关联，这意味着它们的锁申请行为可以在一个goroutine中操作，释放锁行为可以在另一个goroutine中操作。

由于RLock()和Lock()都能保证数据不被其它goroutine修改，所以在RLock()与RUnlock()之间的，以及Lock()与Unlock()之间的代码区都是critical section。

以下是一个示例，此示例中同时使用了Mutex和RWMutex，RWMutex用于读、写，Mutex只用于读。

```go
package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var Password = secret{password: "myPassword"}

type secret struct {
	RWM      sync.RWMutex
	M        sync.Mutex
	password string
}

// 通过rwmutex写
func Change(c *secret, pass string) {
	c.RWM.Lock()
	fmt.Println("Change with rwmutex lock")
	time.Sleep(3 * time.Second)
	c.password = pass
	c.RWM.Unlock()
}

// 通过rwmutex读
func rwMutexShow(c *secret) string {
	c.RWM.RLock()
	fmt.Println("show with rwmutex",time.Now().Second())
	time.Sleep(1 * time.Second)
	defer c.RWM.RUnlock()
	return c.password
}

// 通过mutex读，和rwMutexShow的唯一区别在于锁的方式不同
func mutexShow(c *secret) string {
	c.M.Lock()
	fmt.Println("show with mutex:",time.Now().Second())
	time.Sleep(1 * time.Second)
	defer c.M.Unlock()
	return c.password
}

func main() {
	// 定义一个稍后用于覆盖(重写)的函数
	var show = func(c *secret) string { return "" }

	// 通过变量赋值的方式，选择并重写showFunc函数
	if len(os.Args) != 2 {
		fmt.Println("Using sync.RWMutex!",time.Now().Second())
		show = rwMutexShow
	} else {
		fmt.Println("Using sync.Mutex!",time.Now().Second())
		show = mutexShow
	}
	
	var wg sync.WaitGroup

	// 激活5个goroutine，每个goroutine都查看
	// 根据选择的函数不同，showFunc()加锁的方式不同
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Go Pass:", show(&Password),time.Now().Second())
		}()
	}
	
	// 激活一个申请写锁的goroutine
	go func() {
		wg.Add(1)
		defer wg.Done()
		Change(&Password, "123456")
	}()
	// 阻塞，直到所有wg.Done
	wg.Wait()
}
```

Change()函数申请写锁，并睡眠3秒后修改数据，然后释放写锁。

rwMutexShow()函数申请读锁，并睡眠一秒后取得数据，并释放读锁。注意，rwMutexShow()中的print和return是相隔一秒钟的。

mutexShow()函数申请Mutex锁，和RWMutex互不相干。和rwMutexShow()唯一不同之处在于申请的锁不同。

main()中，先根据命令行参数数量决定运行哪一个show()。之所以能根据函数变量来赋值，是因为先定义了一个show()函数，它的函数签名和rwMutexShow()、mutexShow()的签名相同，所以可以相互赋值。

for循环中激活了5个goroutine并发运行，for瞬间激活5个goroutine后，继续执行main()代码会激活另一个用于申请写锁的goroutine。这6个goroutine的执行顺序是随机的。

如果show选中的函数是rwMutexShow()，则5个goroutine要申请的RLock()锁和写锁是冲突的，但5个RLock()是兼容的。所以，只要某个时间点调度到了写锁的goroutine，剩下的读锁goroutine都会从那时开始阻塞3秒。

除此之外，还有一个不严格准确，但在时间持续长短的理论上来说能保证的一个规律：当修改数据结束后，各个剩下的goroutine都申请读锁，因为申请后立即print输出，然后睡眠1秒，但1秒时间足够所有剩下的goroutine申请完读锁，使得`show with rwmutex`输出是连在一起，输出的`Go Pass: 123456`又是连在一起的。

某次结果如下：

```
Using sync.RWMutex! 58
show with rwmutex 58
Change with rwmutex lock
Go Pass: myPassword 59
show with rwmutex 2
show with rwmutex 2
show with rwmutex 2
show with rwmutex 2
Go Pass: 123456 3
Go Pass: 123456 3
Go Pass: 123456 3
Go Pass: 123456 3
```

如果show选中的函数是mutexShow()，则读数据和写数据互不冲突，但读和读是冲突的(因为Mutex的Lock()是互斥的)。

某次结果如下：

```
Using sync.Mutex! 30
Change with rwmutex lock
show with mutex: 30
Go Pass: myPassword 31
show with mutex: 31
Go Pass: myPassword 32
show with mutex: 32
Go Pass: 123456 33
show with mutex: 33
show with mutex: 34
Go Pass: 123456 34
Go Pass: 123456 35
```



## 用Mutex还是用RWMutex

Mutex和RWMutex都不关联goroutine，但RWMutex显然更适用于读多写少的场景。仅针对读的性能来说，RWMutex要高于Mutex，因为rwmutex的多个读可以并存。