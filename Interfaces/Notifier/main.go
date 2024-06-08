package main

import "fmt"

type Notifier interface {
	Notify()
}

type EmailNotifier string
type MessageNotifier string

func (e EmailNotifier) Notify() {
	fmt.Println("Email Notifier")
}
func (m MessageNotifier) Notify() {
	fmt.Println("Message  Notifier")
}

func sendNotification(n Notifier) {
	n.Notify()
}

func main() {
	var e EmailNotifier
	var m MessageNotifier
	sendNotification(e)
	sendNotification(m)

}
