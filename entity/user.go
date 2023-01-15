package entity

type User struct {
	Name     string
	Password string
}

type Validator interface {
	validate() error
}

type NotificationHandler interface {
	HandleError(err error)
}

type notificationHandler struct {
}

type UserValidator struct {
	user                User
	notificationHandler NotificationHandler
}

func (a *UserValidator) validate() error {
	return nil
}

func NewUserValidator(
	u User,
	nh NotificationHandler,
) Validator {
	return &UserValidator{
		user:                u,
		notificationHandler: nh,
	}
}
