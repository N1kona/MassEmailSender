package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"time"
)

//Структура подписчика

type EmailPerson struct {
	Email    string
	Name     string
	Surname  string
	Birthday string
}

//Подписчика 1

var name1 = EmailPerson{
	Email:    "mail1",
	Name:     "name1",
	Surname:  "name1",
	Birthday: "01.01",
}

//Подписчика 2

var name2 = EmailPerson{
	Email:    "mail2",
	Name:     "name2",
	Surname:  "name2",
	Birthday: "02.02",
}

// Вызов переданной функции раз в сутки в указанное время.
func callAt(hour, min, sec int, f func()) error {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return err
	}

	// Вычисляем время первого запуска.
	now := time.Now().Local()
	firstCallTime := time.Date(
		now.Year(), now.Month(), now.Day(), hour, min, sec, 0, loc)
	if firstCallTime.Before(now) {
		// Если получилось время раньше текущего, прибавляем сутки.
		firstCallTime = firstCallTime.Add(time.Hour * 24)
	}

	// Вычисляем временной промежуток до запуска.
	duration := firstCallTime.Sub(time.Now().Local())

	go func() {
		time.Sleep(duration)
		for {
			f()
			// Следующий запуск через сутки.
			time.Sleep(time.Hour * 24)
		}
	}()

	return nil
}

// Ваша отправки mail

func SendMail() {

	// Слайс с подписчиками
	var slice []EmailPerson
	slice = append(slice, name1, name2)

	sender := ""   //Аккаунт откуда будет идти рассылка
	password := "" //Пароль аккаунт откуда будет идти рассылка

	//Перебор подписчиков и отправка mail
	for _, el := range slice {

		m := gomail.NewMessage()
		m.SetHeader("From", sender)
		m.SetHeader("To", el.Email)
		m.SetAddressHeader("Cc", sender, "NEVERMORE")
		m.SetHeader("Subject", "Hello"+" "+el.Name)
		m.SetBody("text/html", "<h1>Тут есть всё</h>\n<p>Заходи</p>")

		d := gomail.NewDialer("smtp.gmail.com", 587, sender, password)

		if err := d.DialAndSend(m); err != nil {
			panic(err)
		}
	}
}

func main() {
	//Отложенный вызов отправки mail
	err := callAt(03, 29, 0, SendMail) //Выбор времени отправки
	if err != nil {
		fmt.Println(err)
	}
}
