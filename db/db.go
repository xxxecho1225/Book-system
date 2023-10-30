package db

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

//数据库指针
var A *sql.DB
var err error


//参数含义:数据库用户名、密码、主机ip、端口号、连接的数据库
func InitDB() {
    // 1.打开数据库
    sqlStr := "root:123456@tcp(127.0.0.1:3306)/book?charset=utf8&parseTime=true&loc=Local"
    var err error
    A, err = sql.Open("mysql", sqlStr)
    if err != nil {
        fmt.Println("数据库打开出现了问题", err)
        return
    }
    A.SetMaxIdleConns(10)
    A.SetMaxOpenConns(5)
    // 2.测试数据库是否连接成功
    err = A.Ping()
    if err != nil {
        fmt.Println("数据库连接出现了问题", err)
        return
    }
}

//增加图书
func InsertPage1(bookname string,card string,autor string,num int,press string,booktype string) (err error)  {
    sql := `INSERT into book(bookname,card,autor,num,press,booktype) VALUES (?,?,?,?,?,?)`
    _, err = A.Exec(sql,bookname,card,autor,num,press,booktype)
    if err != nil {
        fmt.Println(err)
        return err
    }
    return nil
}

//删除图书
func DelPage1(bid int) (err error)  {
    sql := `delete FROM book where bid = ?`
    _, err = A.Exec(sql,bid)
    if err != nil {
        fmt.Println(err)
        return err
    }
    return nil
}

//修改图书
func UpdatePage1(bid int,bookname string,card string,autor string,num int,press string,booktype string) (err error) {
    sql := `Update book set bookname = ?,card = ?,autor = ?,num = ?,press = ?,booktype = ? where bid = ?`
    _, err = A.Exec(sql,bookname,card,autor,num,press,booktype,bid)
    if err != nil {
        fmt.Println(err)
        return err
    }
    return nil
}

//增加用户
func InsertPageone1(username string,readername string,password string,phone string,status string,sendday string) (err error)  {
   sql := `INSERT into user(username,readername,password,phone,status,sendday) VALUES (?,?,?,?,?,?)`
   _, err = A.Exec(sql,username,readername,password,phone,status,sendday)
   if err != nil {
       fmt.Println(err)
       return err
   }
   return nil
}

//删除用户
func DelPageone1(aid int) (err error)  {
   sql := `delete FROM user where aid = ?`
   _, err = A.Exec(sql,aid)
   if err != nil {
       fmt.Println(err)
       return err
   }
   return nil
}

//修改用户
func UpdatePageone1(aid int,username string,readername string,password string,phone string,status string,sendday string) (err error) {
   sql := `Update user set username = ?,readername = ?,password = ?,phone = ?,status = ?,sendday = ? where aid = ?`
   _, err = A.Exec(sql,username,readername,password,phone,status,sendday,aid)
   if err != nil {
       fmt.Println(err)
       return err
   }
   return nil
}

//增加借阅信息
func InsertPagetwo1(aid int,bid int,card string,bookname string,username string,begintime string,endtime string,status string) (err error)  {
    sql := `INSERT into history(aid,bid,card,bookname,username,begintime,endtime,status) VALUES (?,?,?,?,?,?,?,?)`
    _, err = A.Exec(sql,aid,bid,card,bookname,username,begintime,endtime,status)
    if err != nil {
        fmt.Println(err)
        return err
    }
    return nil
}

//删除借阅信息
func DelPagetwo1(hid int) (err error)  {
    sql := `delete FROM history where hid = ?`
    _, err = A.Exec(sql,hid)
    if err != nil {
        fmt.Println(err)
        return err
    }
    return nil
}

//修改借阅信息
func UpdatePagetwo1(hid int,aid int,bid int,card string,bookname string,username string,begintime string,endtime string,status string) (err error) {
    sql := `Update history set aid = ?,bid = ?,card = ?,bookname = ?,username = ?,begintime = ?,endtime = ? ,status = ? where hid = ?`
    _, err = A.Exec(sql,aid,bid,card,bookname,username,begintime,endtime,status,hid)
    if err != nil {
        fmt.Println(err)
        return err
    }
    return nil
}
