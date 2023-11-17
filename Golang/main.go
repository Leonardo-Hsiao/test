package main

import "fmt"

func main() {
	fmt.Println("Hello World! 2019-10-12 13:25:40.000000000 +0800")

	// 1. var
	fmt.Println("step 1 -----------------------")
	var var_a int = 1
	var var_b = 2
	var_c := 3

	var var_d uint8 = 255
	var var_e uint16 = 65535
	var var_f uint32 = 4294967295
	var var_g uint64 = 18446744073709551615

	var var_h int8 = -128
	var var_i int8 = 127

	var var_j, var_k = 4, 5
	var var_l, var_m bool

	fmt.Println(var_a, var_b, var_c, var_d, var_e, var_f, var_g, var_h, var_i, var_j, var_k, var_l, var_m)

	// 2. const iota
	fmt.Println("step 2 -----------------------")
	const var_const_a int = 1
	const var_const_b = 2
	const var_const_c = "my name is xrf"

	const (
		var_cc_a = 0
		var_cc_b = 1
		var_cc_c = 2
	)

	fmt.Println(var_const_a, var_const_b, var_const_c, var_cc_a, var_cc_b, var_cc_c)

	// 3.iota 从const开始为0, 每加一个变量iota加1
	fmt.Println("step 3 -----------------------")
	const (
		var_ci_a0 = "?"
		var_ci_a  = iota
		var_ci_b  = iota
		var_ci_c
		var_ci_d
		// 自动跟随上一个
		var_ci_e = "xrf"
		var_ci_f
		var_ci_g = iota
		var_ci_h
	)
	fmt.Println(var_ci_a0, var_ci_a, var_ci_b, var_ci_c, var_ci_d, var_ci_e, var_ci_f, var_ci_g, var_ci_h)

	const (
		var_iota_a = 1 << iota
		var_iota_b = 3 << iota
		var_iota_c
		var_iota_d
	)
	fmt.Println(var_iota_a, var_iota_b, var_iota_c, var_iota_d)

	// 4. 运算
	fmt.Println("step 4 -----------------------")
	fmt.Println(var_a+var_b, var_a <= var_b, var_a >= var_b, var_a != var_b)
	fmt.Println(3/2, 3%2, 2.3/2)
	fmt.Println(true && false, true && true, false || true, false || false, !false, !true)
	fmt.Println(0<<1, 1<<2, 2>>1)
	fmt.Println(1&3, 1|3, 0|3, 3^3, 3^1)

	fmt.Println("step 5 -----------------------")
	// 5. 指针
	var var_ptr *int
	fmt.Println(var_ptr)
	var_ptr = &var_a
	fmt.Println(var_ptr)

	fmt.Println("step 6 -----------------------")
	// 6.if
	var_a = 10
	if var_a > 10 {
		fmt.Println("step 6.1 in if")
		fmt.Println("step 6.2 in if")
	} else if var_a > 9 {
		fmt.Println("step 6.3 in else if")
	} else if var_a > 8 {
		fmt.Println("step 6.4 in else if")
	} else {
		fmt.Println("step 6.4 in else")
	}

	// 7. switch
	fmt.Println("step 7 -----------------------")
	var_b = 1
	switch var_b {
	case 1:
		{
			fmt.Println("step 7.1 in case 1")
			fmt.Println("step 7.1 in case 1")
		}
	// case 1:
	// 	fmt.Println("step 7.2 in case 1")
	case 2, 3, 4:
		fmt.Println("step 7.2 in case 2")
	default:
		fmt.Println("step 7.3 in default")
	}

	//  使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。
	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("4、case 条件语句为 true")
	case false:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}

	switch {

	}

	fmt.Println("step 8 -----------------------")
	fmt.Println("step 9 -----------------------")
	fmt.Println("step 10 -----------------------")
	fmt.Println("step 11 -----------------------")
	fmt.Println("step 12 -----------------------")
	fmt.Println("step 13 -----------------------")

}
