package main

import "time"

func main() {
	eventBus := NewEventBus()

	userRegisteredChan := make(chan Event)
	resetPasswordChan := make(chan Event)

	eventBus.Subscribe("UserRegistered", userRegisteredChan)
	eventBus.Subscribe("ResetPassword", resetPasswordChan)

	go UserRegisteredHandler(userRegisteredChan)
	go ResetPasswordHandler(resetPasswordChan)

	userService := NewUserRegistrationService(eventBus)

	user := userService.RegisterUser("John Doe", "jd@gmail.com")

	userService.ResetPassword(user)

	time.Sleep(time.Hour)
}
