package userservices

func getUserbyEmail(email string) User {
	user := User{}
	// user = findUserByEmail(email)
	return user
}

// func getPathToFile(selected bool) string {

// 	if selected == true {
// 		path := ""
// 		return path
// 		//
// 	}

// }

// func isSelectedLocalStorage(isLocal bool) bool {
// 	if isLocal != false {
// 		return true
// 	}
// }

// func isSelectedRemoteStorage(isRemoteSelected bool) bool {
// 	if isRemoteSelected != false {
// 		return true
// 	}
// }

// func getFileListOfUser(u User) []string {
// 	// listOfFile := []string
// 	// listOfFile = getFileListByUserId(u.ID)
// 	// return listOfFile
// }

// func sendFileToServer(f File) {
// 	// нужен метод по отправке файла не сервер
// }

func removeFileFromListOfUser(fileId int) {
	// нужно подключение к БД

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
