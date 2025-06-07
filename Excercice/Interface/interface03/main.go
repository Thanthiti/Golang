package main

import "fmt"

type Notifier interface {
	Notify() string
}

type email struct {
	text string
}

func (e email) Notify() string {
	return ("Email sent " + e.text)
}

type sms struct {
	adress string
}

func (s sms) Notify() string {
	return ("SMS sent " + s.adress)
}

type PushNotification struct {
	deviceID string
}

func (p PushNotification) Notify() string {
	return ("Push Notification sent " + p.deviceID)
}

func SendNotification(n Notifier) {
	fmt.Println(n.Notify())
}

func main() {
	email := email{"palm@gmail.com"}
	sms := sms{"040404"}
	push := PushNotification{"Iphone"}
	SendNotification(email)
	SendNotification(sms)
	SendNotification(push)
}