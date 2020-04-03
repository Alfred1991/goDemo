package main

import (
	"fmt"
	"math/cmplx"
)

/**
1、go程序由packages组成。
2、程序从main包的main方法执行。
3、capital letter的names才可以exported，import packages的时候只能引用exported names
 */
func main() {

	fmt.Println("-------------------------- var --------------------------")
	/**
	4、var用于定义变量
	 */
	var c,python,java bool
	fmt.Println(c,python,java)

	/**
	5、初始化变量
	 */
	var i,j int = 1,2
	fmt.Println(i,j);

	/**
	6、:=用于简化定义变量，此时会有type inference
	在函数外，所有statement需要以keyword(var,func等)开头。因此函数外不需要:=
	 */
	k := 1
	fmt.Println(k)

	fmt.Println("-------------------------- basic types --------------------------")
	/**
	7、基本类型
	bool
	string
	int   int8  int16  int32  int64
	uint unit8 uint16 uint32 unit64
	byte
	rune
	float32 float64
	complex64 complex128
	 */

	/**
	8、factored 风格的var声明
	 */
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	fmt.Println("-------------------------- type converter --------------------------")
	/**
	9、类型转换
	 */
	var ii int = 42
	var ff float64 = float64(ii)
	var uu uint = uint(ff)
	fmt.Printf("Type: %T Value: %v\n",ii,ii)
	fmt.Printf("Type: %T Value: %v\n",ff,ff)
	fmt.Printf("Type: %T Value: %v\n",uu,uu)

	fmt.Println("-------------------------- constant --------------------------")
	/*
	10、const
	 */
	const comp_const = 1+1i
	const char_const = 'c'
	const string_const = "s\""

	fmt.Printf("Type: %T Value: %v\n",comp_const,comp_const)
	fmt.Printf("Type: %T Value: %v\n",char_const,char_const)
	fmt.Printf("Type: %T Value: %v\n",string_const,string_const)

	/**

	 */

}

/**
11、定义有返回值的函数
 */
func add(x int,y int) int{
	return x + y
}

/**
12、简化函数定义
 */
func add2(x,y int) int{
	return x + y
}

/**
13、多返回值函数
 */
func swap(x,y string) (string,string) {
	return y,x
}

/**
14、命名返回值
 */
func split(sum int) (x,y int){
	x = sum * 4 /9
	y = sum - x
	return
}

