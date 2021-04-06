package model

// 定义一个结构体
type student struct {
	Name string
	score float64
}

// 因为 student 结构体首字母小写，只能在 model 包使用
// 通过工厂模式解决这个问题

// 定义一个 NewStudent 函数,形参为 n 和 s 对应 student 结构体的 Name 和 score
// 返回值为 *student 结构体的取值运算
// 然后 return 返回值中通过地址传递 n、s 给 student 结构体的 Name 和 score 字段
func NewStudent(n string,s float64) *student {
 	return &student {
 		Name : n,
 		score : s,
	}
}

// 如果 score 字段首字母小写则在其他包不可以直接访问，这时候我们可以提供一个方法来解决
// 因为 GetScore() 首字母大写就说明是对外公开的 
func (stu *student) GetScore() (float64) {
	return stu.score
}