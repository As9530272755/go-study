package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	 str6 := "hello北"						 // 假如这里我们有 北 这个中文
	 fmt.Println("str len=",len(str6))		 // 则输出 hello 的 5 个字节和 北 的 3 个字节

	 // 想把这个 str2 中的每个字符串打印出来，我们可以通过遍历的方法
	 str7 := "hello北京"

	 // 处理有中文的我们需要将他转换问 rune 的切片
	r := []rune(str7)
	for i :=0 ; i < len(r) ; i++ {
	 	fmt.Printf("字符=%c\n",r[i])
	}
	 // 这里有一个 n,err 赋值了一个 stsconv.Atoi 的函数，并且转一个 123 的整数
	n, err := strconv.Atoi("hello")
	
	 // 这里的 nil 表示空指针，也就是判断 err 这个变量有没有值，如果 err 没有值就为真，就执行 if 判断里面的语句块
	 // 如果 err 有值就为假，执行其他的语句块
	if err != nil {
	
	 //	 如果 err 不等于空就输出下面这个 println 语句
	 	fmt.Println("转换错误\n",err)
	} else {
		fmt.Println("转换成功结果=\n",n)
	}

	str := strconv.Itoa(12345)
	fmt.Printf("str=%v,str=%T\n",str,str)

	// 这样输出的时候就是将 hello 这个单词对应的 ASCII 码进行输 出
	var bytes = []byte("hello go")
	fmt.Printf("hello 字符串转换后%v\n",bytes)

	str1 := string([]byte{97,98,99})
	fmt.Printf("str 通过字符串转换后为%v\n",str1)

	str2 := strconv.FormatInt(123,2)
	fmt.Printf("转为 2 进制之后的 str 为%v\n",str2)

	str3 := strconv.FormatInt(123,8)
	fmt.Printf("转为 8 进制之后的 str 为%v\n",str3)

	str4 := strconv.FormatInt(123,16)
	fmt.Printf("转为 16 进制之后的 str 为%v\n",str4)

	str5 := strconv.FormatInt(123,32)
	fmt.Printf("转为 32 进制之后的 str 为%v\n",str5)

	b := strings.Contains("seafood","zhang")  //返回值是 bool 类型
	fmt.Printf("b=%v\n",b)

	num := strings.Count("ceheese","ae")
	fmt.Printf("有出现ae=%v次\n",num)

	d := strings.EqualFold("asd","AsD")
	fmt.Printf("d = %v\n",d)

	fmt.Println("结果","ab" == "AB")

	index := strings.Index("NLT_abc","abc")
	fmt.Printf("abc 的 index 值为%v\n",index)

	index1 := strings.Index("NLT_abc","asd")
	fmt.Printf("abc 的 index 值为%v\n",index1)

	index2 := strings.Index("NLT_abcabcabc","abc")
	fmt.Printf("多次 abc 的 index 值为%v\n",index2)
}