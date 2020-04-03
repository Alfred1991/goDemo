package main

import (
	"fmt"
	"image"
	"io"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("-------------------------- mothod --------------------------")
	/**
	1、使用方法
	 */
	v := NewVertex{3,4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))

	/**
	使用非结构体的方法
	 */
	f:=MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	/**
	对于方法来说，无论方法定义中的receiver类型是不是pointer。都既可以用值调用，也可以用值的引用来调用该方法，但
	当receiver类型不是pointer时:
	  方法是值传递，方法参数是形式参数，方法无法改变原值。
	当receiver类型是pointer时：
	  方法是引用传递，方法参数是实参的引用，方法可以改变原值。此时避免的值复制，因此更高效。

	在创建某个类型的方法时，最好统一receiver类型。要么都是pointer，要么都不是。
	 */
	v.toString()

	v.Scale_value(2)
	v.toString()

	(&v).Scale_value(2)
	v.toString()

	(&v).Scale_point(2)
	v.toString()

	v.Scale_point(2)
	v.toString()

	/**
	对于函数来说，若函数参数类型是pointer则必须传递pointer，若函数参数类型不是pointer则不可传递pointer，且
	当函数参数类型不是pointer时:
	  函数是值传递，函数参数是形式参数，函数无法改变原值。
	当函数参数类型是pointer时：
	  函数是引用传递，函数参数是实参的引用，函数可以改变原值。
	 */
	Scale_point(&v,2)
	v.toString()
	Scale_value(v,2)
	v.toString()


	fmt.Println("-------------------------- interface --------------------------")
	/**
	2、interface类型作为方法签名被定义。
	interface类型的值可以 指向 任何实现该interface指定方法的 值

	当实现接口的所有方法的receiver中有pointer时，接口值 需 指向实现类型的值的引用(&)
	当实现接口的所有方法的receiver都不是pointer时，接口值 可 指向实现类型的值本身(不使用&)，也可 指向实现类型的值的引用(&)

	因此建议 在接口实现统一使用 引用，如下：
	 */
	var a Abser = &NewVertex{3,4}
	a.toString()
	fmt.Println(a.Abs())

	/**
	toString方法中追加了nil保护逻辑
	 */
	var b *NewVertex
	var c Abser = b
	c.toString()

	/**
	empty interface可以指向任何值
	 */
	var i interface{}
	fmt.Printf("(%v,%T)\n",i,i)
	i = 42
	fmt.Printf("(%v,%T)\n",i,i)
	i = "hello"
	fmt.Printf("(%v,%T)\n",i,i)
	i = NewVertex{1,2}
	fmt.Printf("(%v,%T)\n",i,i)
	i = &NewVertex{3,4}
	fmt.Printf("(%v,%T)\n",i,i)


	fmt.Println("-------------------------- type assertion & type switches --------------------------")
	/**
	3、类型断言提供 访问 interface潜在的 实现的值 的方式
	 */
	i = "hello"
	s := i.(string)
	fmt.Println(s)
	s,ok := i.(string)
	fmt.Println(s,ok)
	ff,ok := i.(float64)
	fmt.Println(ff,ok)
	//ff = i.(float64) //panic
	//fmt.Println(ff)

	/**
	type switches
	 */
	do(21)
	do("hello")
	do(true)


	fmt.Println("-------------------------- Stringers --------------------------")
	/**
	4、最普遍存在的接口:
	type Stringer interface {
	    String() string
	}
	fmt使用该接口打印values
	 */
	aa := Person{"Arthur Dent", 42}
	zz := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(aa,zz)


	fmt.Println("-------------------------- error --------------------------")
	/**
	5、error类型是一个类似fmt.Stringer的内建类型：
	type error interface {
	    Error() string
	}
	函数经常返回一个error值，调用函数的代码通过判断该error值是否为nil来处理异常

	使用fmt打印一个error值，也会调用其Error()方法。
	 */
	i,err := strconv.Atoi("42")
	if err != nil {
		fmt.Printf("couldn`t convert number:%v\n",err)
		return
	}
	fmt.Println("Converted integer:",i)

	if err := run();err != nil {
		fmt.Println(err)
	}


	fmt.Println("-------------------------- Readers --------------------------")
	/**
	6、io.Reader接口代表数据流的读入:
	  func (T) Read(b []byte) (n int, err error)
	*/
	r:=strings.NewReader("Hello,Reader!")
	bb := make([]byte,8)
	for {
		n,err := r.Read(bb)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, bb)
		fmt.Printf("b[:n] = %q\n", bb[:n])
		if err == io.EOF {
			break
		}
	}

	fmt.Println("-------------------------- Images --------------------------")
	/**
	7、package image

	type Image interface {
	    ColorModel() color.Model
	    Bounds() Rectangle
	    At(x, y int) color.Color
	}
	 */
	m := image.NewRGBA(image.Rect(0,0,100,100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0,0).RGBA())
	fmt.Println(m.At(1105,1105).RGBA())
	fmt.Printf("%T",m.ColorModel())



}

/**
go没有classes，但可以也能定义方法
方法是特殊的函数—— a function with a receiver argument
方法的receiver的类型必须也是在 定义方法的包 中定义的，因此 不能为内建的基本类型定义方法。
 */
func(v NewVertex) Abs() float64{
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
func Abs(v NewVertex) float64{
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

type NewVertex struct {
	X, Y float64
}

/**
非结构体也能定义方法
 */
type MyFloat float64
func(f MyFloat) Abs() float64{
	if f<0{
		return float64(-f)
	}
	return float64(f)
}

/**
receiver类型为pointer的方法
 */
func(v *NewVertex) Scale_point(f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}
func(v NewVertex) Scale_value(f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}
func(v *NewVertex) toString(){
	if v == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Printf("X:%v,Y:%v\n",v.X,v.Y)
}
//receiver类型的pointer的函数
func Scale_point(v *NewVertex,f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}
func Scale_value(v NewVertex,f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}
/**
interface的实现是隐式的，非声明的
这解耦了interfaces和它的实现
 */
type Abser interface{
	toString()
	Abs()float64
}

func do(i interface{}){
	switch v:=i.(type) {
	case int:
		fmt.Printf("Twice %v is %v \n",v,v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type Person struct{
	Name string
	Age int
}

func(p Person) String() string{
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type MyError struct {
	When time.Time
	What string
}
func(e *MyError) Error() string{
	return fmt.Sprintf("At %v, %s",e.When, e.What)
}
func run() error {
	return &MyError{
		time.Now(),"it didn`t work",
	}
}

