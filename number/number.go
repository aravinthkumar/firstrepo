package number

func sumOfNumberRef(a int) {
	//fmt.Println("During Function Call:", *a, a)
	a = 2
	//	fmt.Println("During Function Call Value is changed", *a, a)
}

func sumOfNumber(a int) {
	//	fmt.Println("During Function Call:", a)
	a = 2
	//	fmt.Println("During Function Call Value is changed", a)
}

func sumOfNumberMore(a int) map[int]string {
	//	fmt.Println("During Function Call:", a)
	a = 2
	b := 10
	var c = map[int]string{}
	c[1] = "Abcd"
	a = b
	return c
	//	fmt.Println("During Function Call Value is changed", a)
}
