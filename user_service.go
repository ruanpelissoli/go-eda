package main

import (
	"fmt"
	"time"
)

type User struct {
	ID    int
	Name  string
	Email string
}

type ResetPassword struct {
	Email string
	Link  string
}

type UserRegistrationService struct {
	eventBus *EventBus
}

func NewUserRegistrationService(eventBus *EventBus) *UserRegistrationService {
	return &UserRegistrationService{
		eventBus: eventBus,
	}
}

func (urs *UserRegistrationService) RegisterUser(name, email string) User {
	user := User{
		ID:    1,
		Name:  name,
		Email: email,
	}

	event := Event{
		Type:      "UserRegistered",
		Timestamp: time.Now(),
		Data: UserRegisteredEvent{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}

	urs.eventBus.Publish(event)

	return user
}

func (urs *UserRegistrationService) ResetPassword(user User) {
	resetPassword := ResetPassword{
		Email: user.Email,
		Link:  fmt.Sprintf("user/%v/resetpassword", user.ID),
	}

	event := Event{
		Type:      "ResetPassword",
		Timestamp: time.Now(),
		Data: ResetPasswordEvent{
			Email: resetPassword.Email,
			Link:  resetPassword.Link,
		},
	}

	urs.eventBus.Publish(event)
}

func UserRegisteredHandler(eventChan <-chan Event) {
	for event := range eventChan {
		userRegisteredEvent, ok := event.Data.(UserRegisteredEvent)
		if !ok {
			fmt.Println("Invalid event data")
			continue
		}

		fmt.Printf("New user Registered: %+v \n", userRegisteredEvent)
	}
}

func ResetPasswordHandler(eventChan <-chan Event) {
	for event := range eventChan {
		resetPasswordEvent, ok := event.Data.(ResetPasswordEvent)
		if !ok {
			fmt.Println("Invalid event data")
			continue
		}

		fmt.Printf("Password link sent to %v, with link %v \n", resetPasswordEvent.Email, resetPasswordEvent.Link)
	}
}
