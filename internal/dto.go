package internal

type User struct {
	userName string
}

func NewUser(userName string) User {
	return User{
		userName: userName,
	}
}

func (u *User) GetUsername() string {
	return u.userName
}
