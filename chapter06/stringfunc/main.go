package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	  str6 := "hello北"						  //假如这里我们有 北 这个中文
	  fmt.Println("str len=",len(str6))		  //则输出 hello 的 5 个字节和 北 的 3 个字节

	   //想把这个 str2 中的每个字符串打印出来，我们可以通过遍历的方法
	  str7 := "hello北京"

	   //处理有中文的我们需要将他转换问 rune 的切片
	 r := []rune(str7)
	 for i :=0 ; i < len(r) ; i++ {
	  	fmt.Printf("字符=%c\n",r[i])
	 }
	   //这里有一个 n,err 赋值了一个 stsconv.Atoi 的函数，并且转一个 123 的整数
	 n, err := strconv.Atoi("hello")
	
	//    这里的 nil 表示空指针，也就是判断 err 这个变量有没有值，如果 err 没有值就为真，就执行 if 判断里面的语句块
	//    如果 err 有值就为假，执行其他的语句块
	 if err != nil {
	
	  	// 如果 err 不等于空就输出下面这个 println 语句
	  	fmt.Println("转换错误\n",err)
	 } else {
	 	fmt.Println("转换成功结果=\n",n)
	 }

	 str := strconv.Itoa(12345)
	 fmt.Printf("str=%v,str=%T\n",str,str)

	//  这样输出的时候就是将 hello 这个单词对应的 ASCII 码进行输 出
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

	 index3 := strings.LastIndex("go golang","go")
	 fmt.Printf("go 最后一次出现的 index 值为%v",index3)
	 index4 := strings.Replace("go go hello","go","go 语言",1)
	 fmt.Printf("替换后的go %v\n",index4)

	 index5 := "go go hello"
	 index6 := strings.Replace(index5,"go","北京",-1)
	 fmt.Printf("替换后%v\nindex5 本身没有发生变化还等于 %v\n",index6,index5)
	
	 strArr := strings.Split("hello,wrold,ok",",")
	 for i := 0 ; i < len(strArr) ; i++ {
	 	fmt.Printf("strArr[%v]=%v\n",i,strArr[i])
	 }
	fmt.Printf("strArr=%q\n",strArr)

	 //该字符串里面有大小写的字母
	strx := "goLang Hello"
	
	 //我们想将他统一的换成小写的字母
	 strxx := strings.ToLower(strx)
	 fmt.Printf("大写转为小写 %v\n",strxx)

	//  小写转为大写
	strdx := strings.ToUpper(strx)
	fmt.Printf("小写转为大写 %v\n",strdx)

	// 该函数将字符串左右两边的空格去掉

	strkg := strings.TrimSpace("  去掉左右两边的空格       ")
	fmt.Printf("去掉之后 = %q\n",strkg)

	str = "! hello! !"
	strzy := strings.Trim(str," !")
	fmt.Printf("将左右两边的字符去掉之后 %v\n",strzy)

	str = "! hello!"
	strz := strings.TrimLeft(str," !")
	fmt.Printf("去掉左边之后的 str= %v\n",strz)

	str = "! hello!"
	stry := strings.TrimRight(str,"!")
	fmt.Printf("右边去掉后 str = %v\n",stry)

	b = strings.HasPrefix("ftp://192.168.10.1","www")
	fmt.Printf("判断是否以 www 开头是就返回 = %v\n",b)

	d = strings.HasSuffix("NLT_abc.jpg","abc")
	fmt.Printf("该字符串是否以 abc 结束 %v",d)
}
