package main

import (
	"fmt"
	"golang.org/x/tools/go/ssa/interp/testdata/src/runtime"
	"math"
	"time"
)

func main() {
	fmt.Println("-------------------------- for --------------------------")
	/**
	1、for循环
	 */
	sum := 0
	for i:=0;i<10;i++{
		sum += i
	}
	fmt.Println(sum)

	/**
	for循环可省略 init和post statements
	 */
	sum = 1
	for ;sum<1000;{
		sum += sum
	}
	fmt.Println(sum)

	/**
	把for当while用
	 */
	sum = 1
	for sum < 1000 {
		sum+=sum
	}
	fmt.Println(sum)


	fmt.Println("-------------------------- if --------------------------")
	/**
	2、if条件判断
	 */
	fmt.Println(sqrt(2),sqrt(-4))
	fmt.Println(pow(3,2,10),pow(3,3,20))

	fmt.Println("-------------------------- switch --------------------------")
	/**
	3、switch在执行完匹配项后会停止。
	不像java中遇到break才停止。
	 */
	fmt.Print("Go runs on ")
	switch  os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("%s.\n",os)
	}

	/**
	没有条件语句的switch相当于switch true
	 */
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 18:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	fmt.Println("-------------------------- defer --------------------------")
	/**
	4、defer推迟一个函数的执行，直到surrounding function（外层函数）返回
	 */
	defer fmt.Println("world(deferred)")
	fmt.Println("hello")

	/**
	defer将函数入栈，当外层函数返回时，defer函数按先进后出执行。
	 */
	for i:=0;i<10;i++{
		defer fmt.Println(i)
	}
}


/**
if可省略() 但不可省略 {}
 */
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
/**
if中可以由一个short statement开头，该statement定义的变量在if结束之前有效
 */
func pow(x,n,lim float64) float64{
	if v:=math.Pow(x,n);v<lim{
		return v
	}else {
		fmt.Printf("%g >= %g\n",v,lim)
	}
	return lim
}
