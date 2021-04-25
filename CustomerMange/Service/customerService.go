package Service

import (
	"go_code/CustomerMange/Model"
	"fmt"
)

// 改结构体完成对 Model 包中 Customer 结构体的操作
// 包括增删改查
type CustomerService struct {
	customers []Model.Customer
	
	// 声明一个字段，表示当前切片含有多少个客户
	// 该字段后面还可以作为新客户的编号+1
	customerNum int
}

// 编写函数，可以返回 *CustomerService ，这个函数的功能是初始化 customer
func NewCustomerService() *CustomerService {
	// 为了能够看到有客户在切片中，我们初始化一个客户
	customerservice := &CustomerService{}
	customerservice.customerNum = 1

	// 调用 model 包，中的 NewCustomer 函数并传入初始值，赋值给 customers 变量
	customer := Model.NewCustomer(1,"tom","男",23,"17749962680",50038119961024777)

	// 通过 append 往 customerservice.customers 切片中追加 customers 中的元素
	customerservice.customers = append(customerservice.customers,customer)

	return customerservice
}

// 返回客户切片而且一定是一个指针类型，都是在原先得 customer 实例上添加
func (this *CustomerService) List() []Model.Customer {
	return this.customers
}

// 添加客户到 customer 切片中,添加成功返回 true
// *CustomerService 为指针类型，这样才能保证在这个系统里面一直用的就是这个 CustomerService
func (this *CustomerService) Add(customer Model.Customer) bool {

	// 我们确定一个分配 id 的规则，就是添加的顺序
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers,customer)
	return true
}

// 根据 Id 查找客户在切片中对应的下标，如果没有该客户返回 -1
func (this *CustomerService) FindById(id int) int {

	// 定义 index 初始为 -1
	index := -1
	
	// 遍历
	for i := 0 ; i < len(this.customers) ; i++ {
		if id == this.customers[i].Id{
			// 说明找到
			index = i
		} else {
			fmt.Printf("没有编号为 %v 的客户",id )
		}
	}

	return index
}

// 根据 id 删除客户，是从切片中删除
func (this *CustomerService) Delete(id int) bool {
	// 调用 this.FindById 并传入 id
	index := this.FindById(id)

	// 如果 index 为 -1 就说明要删除的客户根本不存在
	if index == -1 {
		return false 
	}

	// 如果从切片中删除一个元素
	// this.customers[:index] 表示从 0 的小标找到 index 的小标，但是这个地方是不包含 index 的
	// this.customers[index+1:]... 表示把从 index+1 的这个下标到最后的元素通过 append 追加到 this.customers
	this.customers = append(this.customers[:index],this.customers[index+1:]...)
	return true
}

// 根据 id 修改客户
func (this *CustomerService) Revise(id int,customer Model.Customer) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	} else {
		// 将传入的需要修改的 id 赋值给  customer.Id
		customer.Id = id
	}

	this.customers[index] = customer
	return true
}