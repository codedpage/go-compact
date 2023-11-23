package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your required driver
)

// Model Struct
type Employee struct {
	Id   int
	Name string `orm:"size(100)"`
}

type Post struct {
	Id       int       `orm:"auto"`
	Title    string    `orm:"size(100)"`
	Employee *Employee `orm:"rel(fk)"`
}

func init() {
	// register model
	orm.RegisterModel(new(Employee))
	orm.RegisterModel(new(Post))

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8", 30)

}

func main() {
	basic()
	relational()
	raw_join()
	raw()

}

func relational() {

	var posts []*Post
	orm.Debug = true
	o := orm.NewOrm()
	qs := o.QueryTable("post")
	//num, err := qs.Filter("id__gt", "0").All(&posts)
	num, err := qs.All(&posts)
	fmt.Printf("Num: %d, ERR: %v\n", num, err)

	for _, m := range posts {

		fmt.Println(m.Id, m.Title, m.Employee.Id)
	}
}

func raw_join() {

	var posts []*Post
	orm.Debug = true
	o := orm.NewOrm()
	//qs := o.QueryTable("post")
	//num, err := qs.Filter("id__gt", "0").All(&posts)
	num, err := o.Raw("SELECT * FROM employee e join post p on e.id = p.employee_id WHERE e.id = ?", "1").QueryRows(&posts)
	fmt.Printf("Num: %d, ERR: %v\n", num, err)

	for _, m := range posts {

		fmt.Println(m.Id, m.Title, m.Employee.Id)
	}
}

func raw() {
	o := orm.NewOrm()
	var maps []orm.Params
	num, _ := o.Raw("SELECT id FROM employee WHERE id = ?", "1").Values(&maps)
	if num > 0 {
		fmt.Println(maps[0]["id"])
	}
}

func basic() {

	//debug
	orm.Debug = true

	o := orm.NewOrm()

	user := Employee{Name: "slene"}

	// insert
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	user.Name = "manojks"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// read one
	u := Employee{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	// delete
	//num, err = o.Delete(&u)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}

//https://go.dev/play/p/8QKp7zeLNXu
