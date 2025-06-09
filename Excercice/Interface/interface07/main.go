package main

import "fmt"

type Notifier interface {
    Notify(msg string)
}

type EmailNotifier struct {
    email string
}

func (e EmailNotifier) Notify(msg string) {
    fmt.Printf("Sending email to %s: %s\n", e.email, msg)
}

type SMSNotifier struct {
    phone string
}

func (s SMSNotifier) Notify(msg string) {
    fmt.Printf("Sending SMS to %s: %s\n", s.phone, msg)
}

type User struct {
    name     string
    notifier Notifier
}

func (u User) Alert(msg string) {
    u.notifier.Notify(msg)
}

func main() {
    emailNotifier := EmailNotifier{email: "user@example.com"}
    smsNotifier := SMSNotifier{phone: "1234567890"}

    user1 := User{name: "Alice", notifier: emailNotifier}
    user2 := User{name: "Bob", notifier: smsNotifier}

    user1.Alert("You have a new message!")
    user2.Alert("Your OTP is 1234")
}
