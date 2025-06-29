package userservices

func getUserbyEmail(email string) (User, error) {
	user := User{
		Name:  "Алексей",
		Role:  "user",
		Email: "bboy23@mail.ru",
		ID:    1,
	}
	// user = findUserByEmail(email)
	return user, nil
}

func removeFileFromListOfUser(fileId int) {
	/* нужно подключение к БД */

}

func refreshFileListofUser(u User) {

}

func getAuntification(u User)

type User struct {
	ID          int
	Password    string
	Name        string
	Sername     string
	Role        string
	PhoneNumber int
	Email       string
}
