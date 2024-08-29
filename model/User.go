package model

import (
	"database/sql"
	"fmt"
	"teacher2/middleware"
)

// CheckUser 查询用户是否存在
func CheckUser(number string) (int, string) {
	sqlStr := `select id from user where number=?
			union
			select id from admin where number=?`
	var u User
	db.QueryRow(sqlStr, number, number).Scan(&u.ID)
	if u.ID > 0 {
		return 500, "用户已存在"
	}
	return 200, "用户不存在"
}

// CheckLogin 检查登录
func CheckLogin(number string, password string) (int, string, map[string]any) {
	/*sqlstr1 := "select number,password from user"
	if sqlstr1 != "" {
		var u User
		err1 := db.Select(&u, sqlstr1, 1)
		if err1 != nil {
			return 0, 0, ""
		}
		return u.Role, 200, "登录成功"
	} else {
		sqlstr2 := "select number,password from admin"
		var a Admin
		err2 := db.Select(&a, sqlstr2, 1)
		if err2 != nil {
			return 0, 0, ""
		}
		return a.Role, 200, "登录成功"
	}*/
	//var token string
	//var name string
	//var role int
	//var sex string
	//var age string
	//var phone string
	//var email string
	//var address string
	//var employtime string
	//var eduction string
	//var undergraduate string
	//var graduate string
	//var doctorate string
	//
	//sqlStr := `select name,number,role,sex,age,phone,email,address,employtime,eduction,undergraduate,graduate,doctorate from user where number = ? and password = ?
	//UNION
	//select name,number,role,sex,age,phone,email,address,employtime,eduction,undergraduate,graduate,doctorate from admin where number = ? and password = ?`
	//
	//err1 := db.QueryRow(sqlStr, number, password, number, password).Scan(&name, &number, &role, &sex, &age, &phone, &email, &address, &employtime, &eduction, &undergraduate, &graduate, &doctorate)
	//err2 := db.QueryRow(sqlStr, number, EncryptPsw(password), number, EncryptPsw(password)).Scan(&name, &number, &role, &sex, &age, &phone, &email, &address, &employtime, &eduction, &undergraduate, &graduate, &doctorate)
	//
	//token, _, _ = middleware.SetToken(number, password)
	//var s = make(map[string]any)
	//s["token"] = token
	//s["name"] = name
	//s["number"] = number
	//s["role"] = role
	//s["sex"] = sex
	//s["age"] = age
	//s["phone"] = phone
	//s["email"] = email
	//s["address"] = address
	//s["employtime"] = employtime
	//s["eduction"] = eduction
	//s["undergraduate"] = undergraduate
	//s["graduate"] = graduate
	//s["doctorate"] = doctorate
	//
	//if err1 != nil && err2 != nil {
	//	if err1 == sql.ErrNoRows && err2 == sql.ErrNoRows {
	//		return 0, 500, "工号或密码错误", s
	//	} else {
	//		log.Fatal(err1)
	//		log.Fatal(err2)
	//	}
	//}
	//return role, 200, "登录成功", s

	token, _, _ := middleware.SetToken(number, password)
	//查询用户是否在user表中
	var u User
	sqlStr1 := `select id,name,number,role,imgurl,sex,age,phone,email,address,employtime,eduction,undergraduate,graduate,doctorate from user where number=? and password=?`
	err1 := db.QueryRow(sqlStr1, number, password).Scan(&u.ID, &u.Name, &u.Number, &u.Role, &u.Imgurl, &u.Sex, &u.Age, &u.Phone, &u.Email, &u.Address, &u.Employtime, &u.Eduction, &u.Undergraduate, &u.Graduate, &u.Doctorate)
	if err1 == nil {
		return 200, "登录成功", map[string]any{
			"data":    u,
			"token":   token,
			"message": "该用户是user",
		}
	}

	//查询用户是否在admin表中
	var a Admin
	sqlStr2 := `select id,name,number,role from admin where number=? and password=?`
	err2 := db.QueryRow(sqlStr2, number, password).Scan(&a.ID, &a.Name, &a.Number, &a.Role)
	if err2 == nil {
		return 200, "登录成功", map[string]any{
			"data":    a,
			"token":   token,
			"message": "该用户是admin",
		}
	} else if err != sql.ErrNoRows {
		return 500, "工号或密码错误", nil
	}
	return 500, "该用户不存在", nil
}

// AddUser 添加用户信息
func AddUser(u *User) (int, string) {
	sqlStr := `insert into user(name,number,password,role,sex,age,phone,email,address,employtime,eduction,undergraduate,graduate,doctorate) 
			   values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	_, err := db.Exec(sqlStr, u.Name, u.Number, EncryptPsw(u.Password), u.Role, u.Sex, u.Age, u.Phone, u.Email, u.Address, u.Employtime, u.Eduction, u.Undergraduate, u.Graduate, u.Doctorate)
	if err != nil {
		fmt.Println("err:", err)
		return 500, "用户添加失败"
	}
	return 200, "用户添加成功"
}

// DeleteUser 删除用户信息
func DeleteUser(number string) (int, string) {
	sqlStr := `delete from user where number=?`
	_, err := db.Exec(sqlStr, number)
	if err != nil {
		fmt.Println("err:", err)
		return 500, "删除失败"
	}
	return 200, "删除成功"
}

// UpdateUser 教师修改教师信息
//func UpdateUser(token string, u *User) (int, string) {
//
//	_, _, _, number := middleware.ParseToken(token)
//
//	sqlStr := `update user set name=?,number=?,role=?,sex=?,age=?,phone=?,email=?,address=?,employtime=?,eduction=?,undergraduate=?,graduate=?,doctorate=? where number=?`
//
//	_, err := db.Exec(sqlStr, u.Name, u.Number, u.Role, u.Sex, u.Age, u.Phone, u.Email, u.Address, u.Employtime, u.Eduction, u.Undergraduate, u.Graduate, u.Doctorate, number)
//	if err != nil {
//		fmt.Println("err:", err)
//		return 500, "更新失败"
//	}
//	return 200, "更新成功"
//}

// UpdateUser 教师修改自己信息
func UpdateUser(token string, u *User) (int, string) {

	_, _, _, number := middleware.ParseToken(token)

	sqlStr := `update user set name=?,number=?,role=?,sex=?,age=?,phone=?,email=?,address=?,employtime=?,eduction=?,undergraduate=?,graduate=?,doctorate=? where number=?`

	_, err := db.Exec(sqlStr, u.Name, u.Number, u.Role, u.Sex, u.Age, u.Phone, u.Email, u.Address, u.Employtime, u.Eduction, u.Undergraduate, u.Graduate, u.Doctorate, number)
	if err != nil {
		fmt.Println("err:", err)
		return 500, "更新失败"
	}
	return 200, "更新成功"
}

// Update 管理员修改教师信息
func Update(number string, u *User) (int, string) {

	//_, _, _, number := middleware.ParseToken(token)

	sqlStr := `update user set name=?,number=?,role=?,sex=?,age=?,phone=?,email=?,address=?,employtime=?,eduction=?,undergraduate=?,graduate=?,doctorate=? where number=?`

	_, err := db.Exec(sqlStr, u.Name, u.Number, u.Role, u.Sex, u.Age, u.Phone, u.Email, u.Address, u.Employtime, u.Eduction, u.Undergraduate, u.Graduate, u.Doctorate, number)
	if err != nil {
		fmt.Println("err:", err)
		return 500, "更新失败"
	}
	return 200, "更新成功"
}

// SelectUserById 教师根据token查看自己信息
func SelectUserById(token string) (int, string, map[string]any) {
	/*sqlStr := `select id,name,number,password,role,sex,age,phone,email,address,employtime,eduction,undergraduate,graduate,doctorate from user where id=?`
	var u User
	err := db.QueryRow(sqlStr, id).Scan(&u.ID, &u.Name, &u.Number, &u.Password, &u.Role, &u.Sex, &u.Age, &u.Phone, &u.Email, &u.Address, &u.Employtime, &u.Eduction, &u.Undergraduate, &u.Graduate, &u.Doctorate)*/
	//if number == u.Number {
	//	return u, 500, "该用户不存在"
	//}
	var u User
	var b Book
	var c Course
	var p Project
	_, _, _, number := middleware.ParseToken(token)
	fmt.Println("number:", number)

	//sqlStr := `select * from user where number=?`
	sqlStr := `select u.name,u.number,u.role,u.imgurl,u.sex,u.age,u.phone,u.email,u.address,u.employtime,u.eduction,u.undergraduate,u.graduate,u.doctorate,
       b.bookname,b.bookpublishdate,b.bookcontent,b.bookisbn,
       c.course,
       p.projectname,p.projectcontent
	   from user u 
	       join book b on u.number=b.number  
		   join course c on u.number=c.number 
		   join project p on u.number=p.number 
	   where u.number=?`
	err := db.QueryRow(sqlStr, number).Scan(&u.Name, &u.Number, &u.Role, &u.Imgurl, &u.Sex, &u.Age, &u.Phone, &u.Email, &u.Address, &u.Employtime, &u.Eduction, &u.Undergraduate, &u.Graduate, &u.Doctorate,
		&b.BookName, &b.Bookpublishdate, &b.BookContent, &b.BookIsbn,
		&c.Course,
		&p.ProjectName, &p.ProjectContent)
	if err != nil {
		fmt.Println("err:", err)
		return 500, "该用户不存在", nil
	}
	s := make(map[string]any)
	s["user"] = u
	s["book"] = map[string]interface{}{
		"bookname":        b.BookName,
		"bookpublishdate": b.Bookpublishdate,
		"bookcontent":     b.BookContent,
		"bookisbn":        b.BookIsbn,
	}
	s["course"] = map[string]interface{}{
		"course": c.Course,
	}
	s["project"] = map[string]interface{}{
		"projectname":    p.ProjectName,
		"projectcontent": p.ProjectContent,
	}
	//s["bookname"] = b.BookName
	//s["bookpublishdate"] = b.Bookpublishdate
	//s["bookcontent"] = b.BookContent
	//s["bookisbn"] = b.BookIsbn
	//s["name"] = u.Name
	//s["number"] = u.Number
	//s["role"] = u.Role
	//s["imgurl"] = u.Imgurl
	//s["sex"] = u.Sex
	//s["age"] = u.Age
	//s["phone"] = u.Phone
	//s["email"] = u.Email
	//s["address"] = u.Address
	//s["employtime"] = u.Employtime
	//s["eduction"] = u.Eduction
	//s["undergraduate"] = u.Undergraduate
	//s["graduate"] = u.Graduate
	//s["doctorate"] = u.Doctorate
	//s["bookname"] = b.BookName
	//s["bookpublishdate"] = b.Bookpublishdate
	//s["bookcontent"] = b.BookContent
	//s["bookisbn"] = b.BookIsbn

	return 200, "查看用户信息成功", s
	//return 200, "查看用户信息成功", map[string]any{
	//	"user": User{
	//Name:   u.Name,
	//Number: u.Number,
	//Role:          u.Role,
	//Imgurl:        u.Imgurl,
	//Sex:           u.Sex,
	//Age:           u.Age,
	//Phone:         u.Phone,
	//Email:         u.Email,
	//Address:       u.Address,
	//Employtime:    u.Employtime,
	//Eduction:      u.Eduction,
	//Undergraduate: u.Undergraduate,
	//Graduate:      u.Graduate,
	//Doctorate:     u.Doctorate
	//},
	//"book": Book{
	//	BookName:        b.BookName,
	//	Bookpublishdate: b.Bookpublishdate,
	//	BookContent:     b.BookContent,
	//	BookIsbn:        b.BookIsbn,
	//},
	//}
}

//type Num struct {
//	Number string `json:"number"`
//}

// SelectUserOne 管理员根据number查看教师信息
func SelectUserOne(token string, number string) (int, string, map[string]any) {
	/*sqlStr := `select id,name,number,password,role,sex,age,phone,email,address,employtime,eduction,undergraduate,graduate,doctorate from user where id=?`
	var u User
	err := db.QueryRow(sqlStr, id).Scan(&u.ID, &u.Name, &u.Number, &u.Password, &u.Role, &u.Sex, &u.Age, &u.Phone, &u.Email, &u.Address, &u.Employtime, &u.Eduction, &u.Undergraduate, &u.Graduate, &u.Doctorate)*/
	var u User
	var b Book
	var c Course
	var p Project
	_, _, _, number1 := middleware.ParseToken(token)
	fmt.Println("number1:", number1)
	//fmt.Println("Number:", num.Number)

	//sqlStr := `select * from user where number=?`
	sqlStr := `select u.*,
       b.bookname,b.bookpublishdate,b.bookcontent,b.bookisbn,
       c.course,
       p.projectname,p.projectcontent
	   from user u 
	       join book b on u.number=b.number  
		   join course c on u.number=c.number 
		   join project p on u.number=p.number 
	   where u.number=?`

	err := db.QueryRow(sqlStr, number).Scan(&u.ID, &u.Name, &u.Number, &u.Password, &u.Role, &u.Imgurl, &u.Sex, &u.Age, &u.Phone, &u.Email, &u.Address, &u.Employtime, &u.Eduction, &u.Undergraduate, &u.Graduate, &u.Doctorate,
		&b.BookName, &b.Bookpublishdate, &b.BookContent, &b.BookIsbn,
		&c.Course,
		&p.ProjectName, &p.ProjectContent)
	if err != nil {
		fmt.Println("err:", err)
		return 500, "该用户不存在", nil
	}
	s := make(map[string]any)
	s["user"] = u
	s["book"] = map[string]interface{}{
		"bookname":        b.BookName,
		"bookpublishdate": b.Bookpublishdate,
		"bookcontent":     b.BookContent,
		"bookisbn":        b.BookIsbn,
	}
	s["course"] = map[string]interface{}{
		"course": c.Course,
	}
	s["project"] = map[string]interface{}{
		"projectname":    p.ProjectName,
		"projectcontent": p.ProjectContent,
	}
	return 200, "查看用户信息成功", s
}

// SelectAllUser 管理员查看所有教师信息
func SelectAllUser(pageNum int, pageSize int) (int, string, any) {

	offset := (pageNum - 1) * pageSize
	//sqlStr := `select * from user limit ? offset ?`
	sqlStr := `select u.*,
 	      b.id,b.name,b.number,b.bookname,b.bookpublishdate,b.bookcontent,b.bookisbn,
	      c.id,c.name,c.number,c.course,
	      p.id,p.name,p.number,p.projectname,p.projectcontent
		   from user u
		       join book b on u.number=b.number
			   join course c on u.number=c.number
			   join project p on u.number=p.number
			   limit ? offset ?`

	rows, err := db.Query(sqlStr, pageSize, offset)
	if err != nil {
		fmt.Printf("err:%s\n", err)
	}
	defer rows.Close()

	var alldatas []struct {
		User    User    `json:"user"`
		Book    Book    `json:"book"`
		Course  Course  `json:"course"`
		Project Project `json:"project"`
	}

	for rows.Next() {
		//var u User
		//var b Book
		//var c Course
		//var p Project
		/*err := rows.Scan(
			&u.ID, &u.Name, &u.Number, &u.Password, &u.Role, &u.Imgurl, &u.Sex, &u.Age, &u.Phone, &u.Email, &u.Address, &u.Employtime, &u.Eduction, &u.Undergraduate, &u.Graduate, &u.Doctorate,
			&b.BookName, &b.Bookpublishdate, &b.BookContent, &b.BookIsbn,
			&c.Course,
			&p.ProjectName, &p.ProjectContent,
		)*/
		/*alldata = append(alldata, struct {

			Project Project
		}{User: u, Book: b, Course: c, Project: p})*/
		var alldata struct {
			User    User    `json:"user"`
			Book    Book    `json:"book"`
			Course  Course  `json:"course"`
			Project Project `json:"project"`
		}
		err := rows.Scan(
			&alldata.User.ID, &alldata.User.Name, &alldata.User.Number, &alldata.User.Password, &alldata.User.Role, &alldata.User.Imgurl, &alldata.User.Sex, &alldata.User.Age, &alldata.User.Phone, &alldata.User.Email, &alldata.User.Address, &alldata.User.Employtime, &alldata.User.Eduction, &alldata.User.Undergraduate, &alldata.User.Graduate, &alldata.User.Doctorate,
			&alldata.Book.ID, &alldata.Book.Name, &alldata.Book.Number, &alldata.Book.BookName, &alldata.Book.Bookpublishdate, &alldata.Book.BookContent, &alldata.Book.BookIsbn,
			&alldata.Course.ID, &alldata.Course.Name, &alldata.Course.Number, &alldata.Course.Course,
			&alldata.Project.ID, &alldata.Project.Name, &alldata.Project.Number, &alldata.Project.ProjectName, &alldata.Project.ProjectContent,
		)

		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
			return 500, "查看用户信息失败", nil
		}
		alldatas = append(alldatas, alldata)

	}

	//s := make(map[string]any)
	/*s["user"] = user
	s["book"] = book
	s["course"] = course
	s["project"] = project*/
	/*s["user"] = alldata.User
	s["book"] = map[string]interface{}{
		"bookname":        alldata.Book.BookName,
		"bookpublishdate": alldata.Book.Bookpublishdate,
		"bookcontent":     alldata.Book.BookContent,
		"bookisbn":        alldata.Book.BookIsbn,
	}
	s["course"] = map[string]interface{}{
		"course": alldata.Course.Course,
	}
	s["project"] = map[string]interface{}{
		"projectname":    alldata.Project.ProjectName,
		"projectcontent": alldata.Project.ProjectContent,
	}*/
	return 200, "查看用户信息成功", alldatas
}
