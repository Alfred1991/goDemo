package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("-------------------------- pointer --------------------------")

	/**
	1、pointers指针指向value的内存地址。
	*T 是一个执行类型T变量的pointer，其zero值是nil。
	 */
	var p *int

	/**
	&操作符 从其 operand(操作对象) 上获得一个pointer。
	 */
	i:=42
	p = &i

	/**
	*操作符 从其 operand 上获取其指向的value
	 */
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	fmt.Println("-------------------------- struct --------------------------")
	/**
	2、struct是fields的集合
	 */
	fmt.Println(Vertex{1,2})
	v := Vertex{1,2}
	v.X = 4
	fmt.Println(v.X)

	/**
	持有结构体的pointer时，可使用(*p).X访问field，也可简易地使用p.X访问field。
	 */
	pp := &v
	pp.X = 1.2e9
	fmt.Println(v)

	/**
	struct literal结构体字面量使用 列出fields的方式(未指定field默认赋zero) 创建结构体
	&返回指向结构体值的pointer
	 */
	pp = &Vertex{1,2}
	pp = &Vertex{X:1}
	pp = &Vertex{Y:1}
	pp = &Vertex{}

	fmt.Println("-------------------------- array --------------------------")
	/**
	3、Array：类型[n]T是由n个T类型值组成的array。array的大小是固定的，不能resize。
	 */
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0],a[1])
	fmt.Println(a)

	primes := [6]int{2,3,5,7,11,13}
	fmt.Printf("%T,%v,%T\n",primes,primes,primes[0])

	fmt.Println("-------------------------- slice --------------------------")
	/**
	4、Slices的大小是动态的，是一个flexible view into the elements of an array。
	通过指定某个array的low和high bound来创建indices，其中包含low，但不包含high。
	切片不存储数据，它只是指向array的引用。
	 */
	var s []int = primes[1:4]
	fmt.Printf("%T,%v\n",s,s)

	primes_p := &primes
	s  = primes_p[2:4]
	fmt.Printf("%T,%T,%T,%v,%v,%v\n",primes_p,s,s[0],primes_p,s,s[0])

	/**
	切片字面量会创建一个array，再创建一个指向该array的切片
	 */
	q := []int{2,3,5,7,11,13}
	fmt.Printf("%T,%v\n",q,q)

	/**
	切割array时low bound的默认值是0，high bound的默认值是array总长度
	 */
	ss := primes[1:4]
	fmt.Printf("%T,%v\n",ss,ss)
	ss = ss[:2]
	fmt.Printf("%T,%v\n",ss,ss)
	ss = ss[1:]
	fmt.Printf("%T,%v\n",ss,ss)

	/**
	slice有两个属性：
	  length是切片的长度
	  capacity是切片所指向的array的剩余长度(从切片的第一个元素算起array的长度)
	 */
	fmt.Printf("len=%d cap=%d %v \n",len(q),cap(q),q)
	fmt.Printf("len=%d cap=%d %v \n",len(ss),cap(ss),ss)

	/**
	切片的zero值是nil
	nil的切片的length和capacity都是0，且没有潜在的array
	 */
	q = nil

	/**
	切片可由内建的make函数创建
	make可创建slice、map和chan
	 */
	aa := make([]int,5)
	bb := make([]int,0,5)
	fmt.Printf("len=%d cap=%d %v \n",len(aa),cap(aa),aa)
	fmt.Printf("len=%d cap=%d %v \n",len(bb),cap(bb),bb)
	bb = bb[:5]
	fmt.Printf("len=%d cap=%d %v \n",len(bb),cap(bb),bb)

	/**
	append可以给切片追加元素，如果切片潜在的array长度不够则会创建一个新的array，并让切片指向这个新的array
	 */
	fmt.Printf("len=%d cap=%d %v\n", len(ss), cap(ss), ss)
	ss = append(ss,0)
	fmt.Printf("len=%d cap=%d %v\n", len(ss), cap(ss), ss)
	ss = append(ss,1,2,3,4,5)
	fmt.Printf("len=%d cap=%d %v\n", len(ss), cap(ss), ss)
	ss = append(ss,6,7)
	fmt.Printf("len=%d cap=%d %v\n", len(ss), cap(ss), ss)

	fmt.Println("-------------------------- range --------------------------")
	/**
	5、range用于遍历slice或map
	 */
	pow := []int{1,2,4,8,16,32,64,128}
	for i,v := range pow {
		fmt.Printf("2^%d = %d\n",i,v)
	}

	/**
	使用 _ 跳过index/key/value
	当不需要value时可直接省略它
	 */
	pow = make([]int,10)
	for i := range pow{
		pow[i] = 1 << uint(i)
	}
	fmt.Printf("%v\n",pow)
	for _,value := range pow{
		fmt.Printf("%v\n",value)
	}

	fmt.Println("-------------------------- map --------------------------")
	/**
	6、map是key到value的映射
	map的zero值是nil，nil的map没有keys，也不能添加keys
	 */
	var m map[string]FloatVertex
	m = make(map[string]FloatVertex)
	m["Bell Labs"] = FloatVertex{40.68433,-74.39967}
	fmt.Printf("%v\n",m["Bell Labs"])

	/**
	map字面量
	 */
	m = map[string]FloatVertex{
		"Bell Labs": FloatVertex{
			40.68433, -74.39967,
		},
		"Google": FloatVertex{
			37.42202, -122.08408,
		},
		"Xiaofeng": {
			1232.12,123123.41,
		},
	}
	fmt.Printf("%v\n",m)

	/**
	删除map中元素
	 */
	delete(m,"Bell Labs")

	/**
	测试元素是否存在
	 */
	elem, ok := m["Google"]
	if ok {
		fmt.Printf("Google:%v\n",elem)
	}

	fmt.Println("-------------------------- function variable --------------------------")
	/**
	7、函数也能是变量，也能作为函数参数/返回值
	 */
	hypot := func(x,y float64) float64{
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5,12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	fmt.Println("-------------------------- function closure --------------------------")
	/**
	8、函数闭包指 一个函数值 引用了函数体之外的变量。
	函数既可以访问该变量，也能给该变量赋值，可以认为函数绑定了这个变量
	*/
	pos,neg := adder(),adder()
	for i:=0;i<10;i++ {
		fmt.Println(pos(i),neg(-2*i))
	}

}

type Vertex struct {
	X int
	Y int
}
type VertexFloatVertex struct {
	Lat,Long float64
}
/**
function type的类型字面量 func(T ...) T
 */
func compute(fn func(float64,float64) float64) float64{
	return fn(3,4)
}
/**
一个闭包函数
 */
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		fmt.Println(sum)
		return x
	}
}