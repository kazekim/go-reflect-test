/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"fmt"
	"reflect"
)

type student struct {
	name      string
	yearOfGraduation int
	nickName  *string
	c company
	grade map[string]float64
	specialist [4]string
}

type company struct {
	name string
	city string
	postCode int64
	tel *string
}

func main() {

	n := "Kim"
	tel := "66800000000"
	s := student{
		name:      "Jirawat",
		yearOfGraduation: 58,
		nickName: &n,
		c: company{
			name: "Macademia Inc.",
			city: "Bangkok",
			postCode:10310,
			tel: &tel,
		},
		grade: map[string]float64{
			"1": 3.83,
			"2": 3.78,
		},
		specialist: [4]string{
			"GoLang",
			"Swift",
			"Rust",
			"Kotlin",
		},
	}

	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	k := t.Kind()
	fmt.Println("Type ", t)
	fmt.Println("Value ", v)
	fmt.Println("Kind ", k)

	nm := t.Name()
	fmt.Println("Name ", nm)

	pkg := t.PkgPath()
	fmt.Println("Package ", pkg)

	size := t.Size()
	fmt.Println("Size ", size)

	if f1, ok := t.FieldByName("nickName"); ok {
		fmt.Println("Type of ",f1.Name,": ", f1.Type)
	}

	f := t.Field(0)
	fmt.Println("Type of ",f.Name,": ", f.Type)

	f2 := t.FieldByIndex([]int{3,2})
	fmt.Println("Type of ",f2.Name,": ", f2.Type)

	num := t.NumField()
	fmt.Println("Num Field ", num)

	if f4, ok := t.FieldByName("grade"); ok {
		fmt.Println("Type of ",f4.Name,": ", f4.Type.Key())
	}

	if f5, ok := t.FieldByName("specialist"); ok {
		fmt.Println("Type of ",f5.Name,": ", f5.Type.Len())
	}

	f3 := t.Field(2)
	e := f3.Type.Elem()
	fmt.Println("Elem of ", f3.Name, " is ", e)

	k = v.Field(0).Type().Kind()
	fmt.Println("Kind1 ", k)

	k = v.Field(0).Kind()
	fmt.Println("Kind2 ", k)

	a := 1
	t = reflect.TypeOf(a)
	fmt.Println("Type ", t)

	v1 := reflect.Indirect(reflect.ValueOf(&s))
	addr := v1.Addr()
	fmt.Println("Addr ", addr)

	tv := v1.Type()
	fmt.Println("Type ", tv)

	kv := v1.Kind()
	fmt.Println("Kind ", kv)
}