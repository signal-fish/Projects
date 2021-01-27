package service

import (
	"projects/go_projects/customer-management-system/model"
)

// CustomerService, 完成对Customer的操作,包括增删改查
type CustomerService struct {
	customers []model.Customer
	// 声明一个字段,表示当前切片含有多少个客户
	customerNum int
}

// 编写一个方法, 可以返回 *CustomerService
func NewCustomerService() *CustomerService {
	// 为了能够看到有客户在切片中,我们初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "Signal", "M", 23, "110", "s@gmail.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

// 返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

func (this *CustomerService) Add(customer model.Customer) bool {
	// 我们确定一个分配id的规则,就是添加的顺序
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	// 如果index == -1, 说明没有这个客户
	if index == -1 {
		return false
	}
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}

// 根据id查找客户在切片中对应下标,如果没有该客户,返回-1
func (this *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			index = i
		}
	}
	return index
}

func (this *CustomerService) GetInfoById(id int) model.Customer {
	i := id - 1
	return this.customers[i]
}

// 根据id修改客户信息
func (this *CustomerService) Modify(id int, customer model.Customer) bool {
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			this.customers[i].Name = customer.Name
			this.customers[i].Gender = customer.Gender
			this.customers[i].Age = customer.Age
			this.customers[i].Phone = customer.Phone
			this.customers[i].Email = customer.Email
		}
	}
	return true
}
