package repository

import (
	"fmt"

	"os"
	"strings"
)

type textUsersRepository struct {
	FilePath string
}

func NewTextUsersRepository(file_path string) *textUsersRepository {
	return &textUsersRepository{
		FilePath: file_path,
	}
}

func (t *textUsersRepository) CreateUser(userInfo string) {
	err := os.WriteFile(t.FilePath, []byte(userInfo+"\n"), 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func (t *textUsersRepository) ReadAllUsers() []string {
	data, err := os.ReadFile(t.FilePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return strings.Split(string(data), "\n")
}

func (t *textUsersRepository) FindByEmail(email string) map[string]string {
	users := t.ReadAllUsers()
	for _, user := range users {
		userData := strings.Split(user, ", ")
		userEmail := strings.Split(userData[1], ": ")[1]
		if userEmail == email {
			userInfo := make(map[string]string)
			for _, info := range userData {
				parts := strings.Split(info, ": ")
				userInfo[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
			}
			return userInfo
		}
	}
	return nil
}

func (t *textUsersRepository) CreateUserWithDetails(name, email, password string) {
	userInfo := fmt.Sprintf("Name: %s, Email: %s, Password: %s", name, email, password)
	t.CreateUser(userInfo)
}

func (t *textUsersRepository) UpdateUserWithDetails(userID int, name, email, password string) {
	users := t.ReadAllUsers()
	if userID >= 0 && userID < len(users) {
		userData := strings.Split(users[userID], ", ")
		updatedInfo := map[string]string{
			"Name":     name,
			"Email":    email,
			"Password": password,
		}

		for key, value := range updatedInfo {
			if value == "" {
				updatedInfo[key] = strings.Split(userData[getIndexUser(key)], ": ")[1]
			}
		}

		updatedUserInfo := fmt.Sprintf("Name: %s, Email: %s, Password: %s\n", updatedInfo["Name"], updatedInfo["Email"], updatedInfo["Password"])

		users[userID] = updatedUserInfo

		err := os.WriteFile(t.FilePath, []byte(strings.Join(users, "\n")), 0644)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("User ID out of range")
	}
}

func (t *textUsersRepository) DeleteUser(userID int) {
	users := t.ReadAllUsers()
	if userID >= 0 && userID < len(users) {
		users = append(users[:userID], users[userID+1:]...)
		err := os.WriteFile(t.FilePath, []byte(strings.Join(users, "\n")), 0644)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("User ID out of range")
	}
}

func getIndexUser(key string) int {
	switch key {
	case "Name":
		return 0
	case "Email":
		return 1
	case "Password":
		return 2
	default:
		return -1
	}
}
