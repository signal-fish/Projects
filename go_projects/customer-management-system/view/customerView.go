package main

import (
	"fmt"
	"projects/go_projects/customer-management-system/model"
	"projects/go_projects/customer-management-system/service"
)

type customerView struct {
	// 定义必要字段
	key  string // 接受用户输入
	loop bool   // 表示是否循环的显示主菜单
	// 增加一个字段customerService
	customerService *service.CustomerService
}

// 显示所有的客户信息
func (this *customerView) list() {
	// 首先,获取到当前所有的客户信息(在切片中)
	customers := this.customerService.List()
	// 显示
	fmt.Println("----------------客户列表----------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("\n--------------客户列表完成--------------\n")
}

// 显示主菜单
func (this *customerView) mainMenu() {
	for {
		fmt.Println("----------------客户信息管理软件----------------")
		fmt.Println("                   1 添加客户")
		fmt.Println("                   2 修改客户")
		fmt.Println("                   3 删除客户")
		fmt.Println("                   4 客户列表")
		fmt.Println("                   5 退    出")
		fmt.Println("请选择(1-5): ")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			this.modify()
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("输入有误,请重新输入!")
		}

		if !this.loop {
			break
		}
	}
	fmt.Println("退出客户管理系统")
}

func main() {
	// 在main函数中,创建一个customerView,并运行显示主菜单
	customerView := customerView{
		key:  "",
		loop: true,
	}
	// 完成对customerView结构体的customerService字段的初始化
	customerView.customerService = service.NewCustomerService()
	// 显示主菜单
	customerView.mainMenu()
}

func (this *customerView) add() {
	fmt.Println("----------------添加客户----------------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮:")
	email := ""
	fmt.Scanln(&email)

	// 构建一个新的Customer实例
	customer := model.NewCustomer2(name, gender, age, phone, email)

	// 调用
	if this.customerService.Add(customer) {
		fmt.Println("----------------添加完成----------------")
	} else {
		fmt.Println("----------------添加失败----------------")
	}
}

func (this *customerView) modify() {
	fmt.Print("请输入要修改的id:")
	id := 0
	fmt.Scanln(&id)
	if this.customerService.FindById(id) != -1 {
		customer := this.customerService.GetInfoById(id)
		fmt.Printf("姓名(%v) :", customer.Name)
		name := ""
		fmt.Scanln(&name)
		fmt.Printf("性别（%v）：", customer.Gender)
		gender := ""
		fmt.Scanln(&gender)
		fmt.Printf("年龄（%v）：", customer.Age)
		age := 0
		fmt.Scanln(&age)
		fmt.Printf("电话（%v）：", customer.Phone)
		phone := ""
		fmt.Scanln(&phone)
		fmt.Printf("邮箱（%v）：", customer.Email)
		email := ""
		fmt.Scanln(&email)
		customer2 := model.NewCustomer2(name, gender, age, phone, email)
		this.customerService.Modify(id, customer2)
	} else {
		fmt.Println("-----------------输入id不存在，请重新输入--------------")
	}
}

func (this *customerView) delete() {
	fmt.Println("----------------删除客户----------------")
	fmt.Println("请选择待删除客户编号(-1退出):")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确认是否删除(Y/N):")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		// 调用customerService的Delete方法
		if this.customerService.Delete(id) {
			fmt.Println("----------------删除完成----------------")
		} else {
			fmt.Println("--------删除失败,输入的id不存在---------")
		}
	}
}

func (this *customerView) exit() {
	fmt.Println("确认是否退出(Y/N):")
	for {
		fmt.Scanln(&this.key)
		if this.key == "Y" || this.key == "y" || this.key == "N" || this.key == "n" {
			break
		}
		fmt.Println("你的输入有误,确认是否退出(Y/N):")
	}
	if this.key == "Y" || this.key == "y" {
		this.loop = false
	}
}
