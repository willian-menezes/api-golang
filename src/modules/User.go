package modules

import (
	"errors"
	"strings"
	"time"

	"github.com/fatih/structs"
)

// User struct is a representation
type User struct {
	ID       uint64    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	CreateAt time.Time `json:"createAt,omitempty"`
}

func (user *User) Preparement(typeProcess string) error {
	if erro := user.validation(typeProcess); erro != nil {
		return erro
	}

	user.formatting()

	return nil
}

func (user *User) validation(typeProcess string) error {

	var labels []string
	data := structs.Map(user)

	for index, value := range data {
		if value == "" {
			if typeProcess != "create" && index == "Password" {
				continue
			}
			labels = append(labels, index)
		}
	}

	if len(labels) > 0 {
		return errors.New("Values requiered: " + strings.Join(labels, ", "))
	}

	return nil

}

func (user *User) formatting() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
