package main

import "fmt"

func main() {
	a := 6

	testTypeConvert(a)
}

func testTypeConvert(a interface{}) {
	//i2:= int(a)//Cannot convert an expression of the type 'interface{}' to the type 'int'
	i := a.(int)
	fmt.Printf("%T\n", i)

	i64, err := a.(int64)
	fmt.Printf("%#v\n", err)

	if err {
		fmt.Printf("%#v\n", err)
	}
	fmt.Printf("%T\n", i64)
	fmt.Printf("%#v", i64)


	//j := a.(int64)
	//fmt.Printf("%T", j)

	//f := a.(float32)
	//fmt.Printf("%T", f)

}
