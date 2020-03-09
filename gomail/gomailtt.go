package main

import (
	"logistics/src/gopkg.in/gomail.v2"
)

func main() {
	m := gomail.NewMessage()

	m.SetHeader("From", "244800087@qq.com")
	m.SetHeader("To", "1142835240@qq.com")
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>hmc 23333</i>!")
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.qq.com", 587, "244800087@qq.com", "laetmwxdubylbjci")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		println("panic!!")
		panic(err)
	}
	println("end!!")
}
