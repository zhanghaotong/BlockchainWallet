/**
  @Author : Zhang Haotong
*/

package controller

import "github.com/kongyixueyuan.com/education/service"

type Application struct {
	Setup *service.ServiceSetup
}

type User struct {
	LoginName string
	Password  string
	IsAdmin   string
}

var users []User

func init() {

	admin := User{LoginName: "admin", Password: "admin", IsAdmin: "T"}
	bob := User{LoginName: "101", Password: "123456", IsAdmin: "F"}
	jack := User{LoginName: "102", Password: "123456", IsAdmin: "F"}

	users = append(users, admin)
	users = append(users, bob)
	users = append(users, jack)

}
