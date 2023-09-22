package main

import (
	"fmt"
  "net/smtp"
)


	//understand hoe to send mail
  //first we create a func sendMail and create a auth 
  func sendMailSample(){
    auth:= smtp.PlainAuth(
      "",
      "vk663434@gmail.com",//mail address
      "dnrqrokzxnllyqdh",// app password
      "smtp.gmail.com",//server host
    )
    msg:="Subject : hi this is subject\nthis is a mail body"
    err:=smtp.SendMail(
      "smtp.gmail.com:587",//host
      auth,
      "vk663434@gmail.com",//from
      []string{"vk663434@gmail.com"},//to
      []byte(msg),//msg
    )
    if err!=nil{
      fmt.Println(err)
    }
  }
func main(){
  sendMailSample()
}
  
  
