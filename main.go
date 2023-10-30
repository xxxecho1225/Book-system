package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"project/db"
	"strconv"
)
type Book_zhoukai struct {
	Bid int `db:"bid"`
	Bookname string `db:"bookname"`
	Card string `db:"card"`
	Autor string `db:"autor"`
	Num int `db:"num"`
	Press string `db:"press"`
	Booktype string `db:"booktype"`
}
type Root_zhoukai struct {
	Rid int `db:"rid"`
	Name string `db:"name"`
	Password string `db:"password"`
}
type History_zhoukai struct {
	Hid int `db:"hid"`
	Aid int `db:"aid"`
	Bid int `db:"bid"`
	Card string `db:"card"`
	Bookname string `db:"bookname"`
	Username string `db:"username"`
	Begintime string `db:"begintime"`
	Endtime string `db:"endtime"`
	Status string `db:"status"`
}
type User_zhoukai struct {
	Aid int `db:"aid"`
	Username string `db:"username"`
	Readername string `db:"name"`
	Password string `db:"password"`
	Phone string `db:"phone"`
	Status string `db:"status"`
	Sendday string `db:"day"`
}
func main() {
	 db.InitDB()

	r := gin.Default()

	// 静态资源加载，本例为css,js以及资源图片
	r.LoadHTMLGlob("static/html/*")
	r.StaticFS("/static", http.Dir("./static"))


	r.GET("/",func(c*gin.Context){
		users :=ChanBook1()
		c.HTML(200,"book-list.html",gin.H{
			"a":users,
		})
	})
	r.GET("/add1",func(c*gin.Context){
		c.HTML(200,"add1.html",gin.H{})
	})
	r.GET("/add2",func(c*gin.Context){
		c.HTML(200,"add2.html",gin.H{})
	})
	r.GET("/add3",func(c*gin.Context){
		c.HTML(200,"add3.html",gin.H{})
	})
	r.GET("/update1",func(c*gin.Context){
		c.HTML(200,"update1.html",gin.H{})
	})
	r.GET("/update2",func(c*gin.Context){
		c.HTML(200,"update2.html",gin.H{})
	})
	r.GET("/update3",func(c*gin.Context){
		c.HTML(200,"update3.html",gin.H{})
	})

	//修改
	r.GET("/update1/:id2",func(c*gin.Context){
		c.HTML(200,"update1.html",gin.H{})
	})
	r.GET("/update2/:id2",func(c*gin.Context){
		c.HTML(200,"update2.html",gin.H{})
	})
	r.GET("/update3/:id2",func(c*gin.Context){
		c.HTML(200,"update3.html",gin.H{})
	})

	//增加图书
	r.POST("/insectPage1",InsertPage)
	//删除图书
	r.GET("/delete/:id",DelPage)
	//修改图书
	r.POST("/update1",UpdatePage)
	//查询图书
	r.POST("/query",QueryPage)

	//增加用户
	r.POST("/insectPage2",InsertPageone)
	//删除用户
	r.GET("/deleteone/:id",DelPageone)
	//修改用户
	r.POST("/update2",UpdatePageone)
	//查询用户
	r.POST("/queryone",QueryPageone)

	//增加借阅信息
	r.POST("/insectPage3",InsertPagetwo)
	//删除借阅信息
	r.GET("/deletetwo/:id",DelPagetwo)
	//修改借阅信息
	r.POST("/update3",UpdatePagetwo)
	//查询借阅信息
	r.POST("/querytwo",QueryPagetwo)


	//页面渲染
	r.GET("/book-list",func(c*gin.Context){
		users :=ChanBook1()
		c.HTML(200,"book-list.html",gin.H{
			"a":users,
		})
	})
	r.GET("/root-list",func(c*gin.Context){
		users :=ChanBook2()
		c.HTML(http.StatusOK,"root-list.html",gin.H{
			"b":users,
		})
	})
	r.GET("/sendbook-list",func(c*gin.Context){
		users :=ChanBook3()
		c.HTML(http.StatusOK,"sendbook-list.html",gin.H{
			"c":users,
		})
	})
	r.GET("/user-list",func(c*gin.Context){
		users :=ChanBook4()
		c.HTML(http.StatusOK,"user-list.html",gin.H{
			"d":users,
		})
	})

	r.Run(":8080")
}

//book-list页面渲染1
func ChanBook1()([]Book_zhoukai){
	var book []Book_zhoukai
	rows, errq := db.A.Query("select bid,bookname,card,autor,num,press,booktype from book")
	if errq != nil {
		log.Fatal(errq.Error)
		return book
	}
	for rows.Next() {
		var a Book_zhoukai
		errn := rows.Scan(&a.Bid, &a.Bookname, &a.Card, &a.Autor, &a.Num, &a.Press,&a.Booktype)
		if errn != nil {
			fmt.Printf("%v", errn)
		}
		book = append(book, a)
	}
	return book
}

//root-list页面渲染2
func ChanBook2()([]Root_zhoukai){
	var root []Root_zhoukai
	rows, errq := db.A.Query("select rid,name,password from root")
	if errq != nil {
		log.Fatal(errq.Error)
		return root
	}
	for rows.Next() {
		var a Root_zhoukai
		errn := rows.Scan(&a.Rid, &a.Name, &a.Password)
		if errn != nil {
			fmt.Printf("%v", errn)
		}
		root = append(root, a)
	}
	return root
}

//sendbook-list页面渲染3
func ChanBook3()([]History_zhoukai){
	var history []History_zhoukai
	rows, errq := db.A.Query ("select hid,aid,bid,card,bookname,username,begintime,endtime,status from history")
	if errq != nil {
		log.Fatal(errq.Error)
		return history
	}
	for rows.Next() {
		var a History_zhoukai
		errn := rows.Scan(&a.Hid, &a.Aid, &a.Bid, &a.Card, &a.Bookname, &a.Username, &a.Begintime, &a.Endtime, &a.Status)
		if errn != nil {
			fmt.Printf("%v", errn)
		}
		history = append(history, a)
	}
	return history
}

//user-list页面渲染4
func ChanBook4()([]User_zhoukai){
	var user []User_zhoukai
	rows, errq := db.A.Query("select aid,username,readername,password,phone,status,sendday from user")
	if errq != nil {
		log.Fatal(errq.Error)
		return user
	}
	for rows.Next() {
		var a User_zhoukai
		errn := rows.Scan(&a.Aid, &a.Username, &a.Readername,&a.Password, &a.Phone,&a.Status,&a.Sendday)
		if errn != nil {
			fmt.Printf("%v", errn)
		}
		user = append(user, a)
	}
	return user
}

//增加图书
func InsertPage(c *gin.Context) {
	bookname := c.DefaultPostForm("bookname", "anonymous")
	card := c.DefaultPostForm("card", "anonymous")
	autor := c.DefaultPostForm("autor", "anonymous")
	num := c.DefaultPostForm("num", "anonymous")
	press := c.DefaultPostForm("press", "anonymous")
	booktype := c.DefaultPostForm("booktype", "anonymous")
	id, _ := strconv.Atoi(num)
	db.InsertPage1(bookname,card,autor,id,press,booktype)
	c.Redirect(http.StatusMovedPermanently,"/book-list")
}

//删除图书
func DelPage(c *gin.Context)  {
	var aa = c.Param("id")
	a1 , _ :=strconv.Atoi(aa)
	db.DelPage1(a1)
	c.Redirect(http.StatusMovedPermanently,"/book-list")
}

//修改图书
func UpdatePage(c *gin.Context) {
	bid := c.DefaultPostForm("bid", "anonymous")
	bids,_ := strconv.Atoi(bid)
	bookname := c.DefaultPostForm("bookname", "anonymous")
	card := c.DefaultPostForm("card", "anonymous")
	autor := c.DefaultPostForm("autor", "anonymous")
	num := c.DefaultPostForm("num", "anonymous")
	nums,_ := strconv.Atoi(num)
	press := c.DefaultPostForm("press", "anonymous")
	booktype := c.DefaultPostForm("booktype", "anonymous")
	db.UpdatePage1(bids,bookname,card,autor,nums,press,booktype)
	c.Redirect(http.StatusMovedPermanently,"/book-list")
}

//查询图书
func QueryPage1(bookname string) (book []Book_zhoukai,err error)  {
	//sql := `select * from book where bookname like "%'+bookname+'%"`
	sql := `SELECT * FROM book where bookname like concat('%',?,'%')`
	rows, err := db.A.Query(sql,bookname)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		res := Book_zhoukai{}
		err = rows.Scan(&res.Bid,&res.Bookname,&res.Card,&res.Autor,&res.Num,&res.Press,&res.Booktype)
		if err != nil {
			return nil, err
		}
		fmt.Println(res)
		book = append(book,res)
	}
	return
}
func QueryPage(c *gin.Context) {
	bookname := c.DefaultPostForm("bookname", "anonymous")
	res, err := QueryPage1(bookname)
	if err != nil {
		return
	}
	c.HTML(200,"book-list.html",gin.H{
		"a" :res,
	})
}


//增加用户
func InsertPageone(c *gin.Context) {
	username := c.DefaultPostForm("username", "anonymous")
	readername := c.DefaultPostForm("readername", "anonymous")
	password := c.DefaultPostForm("password", "anonymous")
	phone := c.DefaultPostForm("phone", "anonymous")
	status := c.DefaultPostForm("status", "anonymous")
	sendday := c.DefaultPostForm("sendday", "anonymous")
	db.InsertPageone1(username,readername,password,phone,status,sendday)
	c.Redirect(http.StatusMovedPermanently,"/user-list")
}

//删除用户
func DelPageone(c *gin.Context)  {
	var aa = c.Param("id")
	a1 , _ :=strconv.Atoi(aa)
	db.DelPageone1(a1)
	c.Redirect(http.StatusMovedPermanently,"http://localhost:8080/user-list")
}
//修改用户
func UpdatePageone(c *gin.Context) {
	aid := c.DefaultPostForm("aid", "anonymous")
	aids,_ := strconv.Atoi(aid)
	username := c.DefaultPostForm("username", "anonymous")
	readername := c.DefaultPostForm("readername", "anonymous")
	password := c.DefaultPostForm("password", "anonymous")
	phone := c.DefaultPostForm("phone", "anonymous")
	status := c.DefaultPostForm("status", "anonymous")
	sendday := c.DefaultPostForm("sendday", "anonymous")
	db.UpdatePageone1(aids,username,readername,password,phone,status,sendday)
	c.Redirect(http.StatusMovedPermanently,"/user-list")
}

//查询用户
func QueryPageone1(readername string) (user []User_zhoukai,err error)  {
	sql := `SELECT * FROM user where readername like concat('%',?,'%')`
	rows, err := db.A.Query(sql,readername)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		res := User_zhoukai{}
		err = rows.Scan(&res.Aid,&res.Username,&res.Readername,&res.Password,&res.Phone,&res.Status,&res.Sendday)
		if err != nil {
			return nil, err
		}
		fmt.Println(res)
		user = append(user,res)
	}
	return
}

func QueryPageone(c *gin.Context) {
	readername := c.DefaultPostForm("readername", "anonymous")
	res, err := QueryPageone1(readername)
	if err != nil {
		return
	}
	c.HTML(200,"user-list.html",gin.H{
		"d" :res,
	})
}

//增加借阅信息
func InsertPagetwo(c *gin.Context) {
	aid := c.DefaultPostForm("aid", "anonymous")
	aids,_ := strconv.Atoi(aid)
	bid := c.DefaultPostForm("bid", "anonymous")
	bids,_ := strconv.Atoi(bid)
	card := c.DefaultPostForm("card", "anonymous")
	bookname := c.DefaultPostForm("bookname", "anonymous")
	username := c.DefaultPostForm("username", "anonymous")
	begintime := c.DefaultPostForm("begintime", "anonymous")
	endtime := c.DefaultPostForm("endtime", "anonymous")
	status := c.DefaultPostForm("status", "anonymous")
	err1 :=judge(username)
	if err1 !=nil{
		c.HTML(200,"error.html",gin.H{
			"remind": "该账号未归还书籍,无法再次借阅书籍,请归还书籍后再借阅书籍",
		})
		return
	}
	db.InsertPagetwo1(aids,bids,card,bookname,username,begintime,endtime,status)
	c.Redirect(http.StatusMovedPermanently,"/sendbook-list")
}

//删除借阅信息
func DelPagetwo(c *gin.Context)  {
	var aa = c.Param("id")
	a1 , _ :=strconv.Atoi(aa)
	db.DelPagetwo1(a1)
	c.Redirect(http.StatusMovedPermanently,"/sendbook-list")
}
//修改借阅信息
func UpdatePagetwo(c *gin.Context) {
	hid := c.DefaultPostForm("hid", "anonymous")
	hids,_ := strconv.Atoi(hid)
	aid := c.DefaultPostForm("aid", "anonymous")
	aids,_ := strconv.Atoi(aid)
	bid := c.DefaultPostForm("bid", "anonymous")
	bids,_ := strconv.Atoi(bid)
	card := c.DefaultPostForm("card", "anonymous")
	bookname := c.DefaultPostForm("bookname", "anonymous")
	username := c.DefaultPostForm("username", "anonymous")
	begintime := c.DefaultPostForm("begintime", "anonymous")
	endtime := c.DefaultPostForm("endtime", "anonymous")
	status := c.DefaultPostForm("status", "anonymous")
	db.UpdatePagetwo1(hids,aids,bids,card,bookname,username,begintime,endtime,status)
	c.Redirect(http.StatusMovedPermanently,"/sendbook-list")
}

//查询借阅信息
func QueryPagetwo1(username string) (user []History_zhoukai,err error)  {
	sql := `SELECT * FROM history where username like concat('%',?,'%')`
	rows, err := db.A.Query(sql,username)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		res := History_zhoukai{}
		err = rows.Scan(&res.Hid,&res.Aid,&res.Bid,&res.Card,&res.Bookname,&res.Username,&res.Begintime,&res.Endtime,&res.Status)
		if err != nil {
			return nil, err
		}
		fmt.Println(res)
		user = append(user,res)
	}
	return
}

func QueryPagetwo(c *gin.Context) {
	username := c.DefaultPostForm("username", "anonymous")
	res, err := QueryPagetwo1(username)
	if err != nil {
		return
	}
	c.HTML(200,"sendbook-list.html",gin.H{
		"c" :res,
	})
}
func judge(usernames string)( error){
	var count int
	sql := `select count(*) as a from history where username = ?`
	err := db.A.QueryRow(sql,usernames).Scan(&count)
	if err != nil {
		return err
	}
	if count != 0 {
		return errors.New("无法借阅")
	}
	return nil
}