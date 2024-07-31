package model

import "database/sql"

type User struct {
	ID            int            `gorm:"int;not null;unique" json:"id"`
	Name          string         `gorm:"type:varchar(50);not null" json:"name" validate:"required,min=2,max=4" label:"姓名"`
	Number        string         `gorm:"type:varchar(50);not null;unique;primary_key" json:"number" validate:"required,len=6" label:"工号"`
	Password      string         `gorm:"type:varchar(50);not null" json:"password" label:"密码"`
	Role          int            `gorm:"int;not null" json:"role" validate:"required,eq=1" label:"权限"` //1:普通教师 2:管理者
	Imgurl        sql.NullString `gorm:"type:varchar(255)" json:"imgurl" label:"头像地址"`
	Sex           string         `gorm:"type:varchar(50)" json:"sex"  label:"性别"`
	Age           int            `gorm:"type:int" json:"age"  label:"年龄"`
	Phone         string         `gorm:"type:varchar(50);not null" json:"phone" validate:"required,len=11" label:"手机号"`
	Email         string         `gorm:"type:varchar(50);not null" json:"email" validate:"required,email" label:"邮箱"`
	Address       string         `gorm:"type:varchar(50)" json:"address"  label:"家庭住址"`
	Employtime    string         `gorm:"type:varchar(50)" json:"employtime"  label:"任职时间"`
	Eduction      string         `gorm:"type:varchar(50)" json:"eduction"  label:"学历"`
	Undergraduate string         `gorm:"type:varchar(50)" json:"undergraduate"  label:"本科毕业院校"`
	Graduate      string         `gorm:"type:varchar(50)" json:"graduate"  label:"研究生毕业院校"`
	Doctorate     string         `gorm:"type:varchar(50)" json:"doctorate"  label:"博士毕业院校"`
}

type Admin struct {
	ID            int    `gorm:"int;not null;unique" json:"id"`
	Name          string `gorm:"type:varchar(50);not null" json:"name" validate:"required,min=2,max=4" label:"姓名"`
	Number        string `gorm:"type:varchar(50);not null;unique;primary_key" json:"number" validate:"required,len=6" label:"工号"`
	Password      string `gorm:"type:varchar(50);not null" json:"password" label:"密码"`
	Role          int    `gorm:"int;not null" json:"role" validate:"required,eq=2" label:"权限"`
	Sex           string `gorm:"type:varchar(50);not null" json:"sex"  label:"性别"`
	Age           int    `gorm:"type:int;not null" json:"age"  label:"年龄"`
	Phone         string `gorm:"type:varchar(50);not null" json:"phone" validate:"required,len=11" label:"手机号"`
	Email         string `gorm:"type:varchar(50);not null" json:"email" validate:"required,email" label:"邮箱"`
	Address       string `gorm:"type:varchar(50)" json:"address"  label:"家庭住址"`
	Employtime    string `gorm:"type:varchar(50)" json:"employtime"  label:"任职时间"`
	Eduction      string `gorm:"type:varchar(50)" json:"eduction"  label:"学历"`
	Undergraduate string `gorm:"type:varchar(50)" json:"undergraduate"  label:"本科毕业院校"`
	Graduate      string `gorm:"type:varchar(50)" json:"graduate"  label:"研究生毕业院校"`
	Doctorate     string `gorm:"type:varchar(50)" json:"doctorate"  label:"博士毕业院校"`
}

type Book struct {
	ID              int    `gorm:"int;not null;unique" json:"id"`
	Name            string `gorm:"type:varchar(50);not null" json:"name" label:"姓名"`
	Number          string `gorm:"type:varchar(50);not null" json:"number" label:"工号"`
	BookName        string `gorm:"type:varchar(50)" json:"bookname" label:"书籍名称"`
	Bookpublishdate string `gorm:"type:varchar(50)" json:"bookpublishdate" label:"书籍发布时间"`
	BookContent     string `gorm:"type:varchar(50)" json:"bookcontent" label:"书籍内容"`
	BookIsbn        string `gorm:"type:varchar(50)" json:"bookisbn" label:"书籍编号"`
}

type Course struct {
	ID     int    `gorm:"int;not null;unique" json:"id"`
	Name   string `gorm:"type:varchar(50);not null" json:"name" label:"姓名"`
	Number string `gorm:"type:varchar(50);not null" json:"number" label:"工号"`
	Course string `gorm:"type:varchar(50)" json:"course" label:"课程名称"`
}

type Project struct {
	ID             int    `gorm:"int;not null;unique" json:"id"`
	Name           string `gorm:"type:varchar(50);not null" json:"name" label:"姓名"`
	Number         string `gorm:"type:varchar(50);not null" json:"number" label:"工号"`
	ProjectName    string `gorm:"type:varchar(50)" json:"projectname" label:"项目名称"`
	ProjectContent string `gorm:"type:varchar(50)" json:"projectcontent" label:"项目内容"`
}
