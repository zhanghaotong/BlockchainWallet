/**
  @Author : Zhang Haotong
*/

package controller

import (
	"net/http"
	"encoding/json"
	"github.com/kongyixueyuan.com/education/service"
	"fmt"
	"strconv"
)

var cuser User

func (app *Application) LoginView(w http.ResponseWriter, r *http.Request)  {

	ShowView(w, r, "login.html", nil)
}

func (app *Application) Index(w http.ResponseWriter, r *http.Request)  {
	if cuser.IsAdmin == "T"{
		ShowView(w, r, "index.html", nil)
	}else{
		result, err := app.Setup.FindAccInfoByEntityID(cuser.LoginName)
		var acc = service.Account{}
		json.Unmarshal(result, &acc)

		data := &struct {
			Edu service.Account
			CurrentUser User
			Msg string
			Flag bool
			History bool
		}{
			Edu:acc,
			CurrentUser:cuser,
			Msg:"",
			Flag:false,
			History:true,
		}

		if err != nil {
			data.Msg = err.Error()
			data.Flag = true
		}
		ShowView(w, r, "indexForUsers.html", data)
	}
	
}

func (app *Application) IndexForUsers(w http.ResponseWriter, r *http.Request)  {
	result, err := app.Setup.FindAccInfoByEntityID(cuser.LoginName)
	var acc = service.Account{}
	json.Unmarshal(result, &acc)

	data := &struct {
		Edu service.Account
		CurrentUser User
		Msg string
		Flag bool
		History bool
	}{
		Edu:acc,
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
		History:true,
	}

	if err != nil {
		data.Msg = err.Error()
		data.Flag = true
	}
	ShowView(w, r, "indexForUsers.html", data)
}

func (app *Application) Help(w http.ResponseWriter, r *http.Request)  {
	data := &struct {
		CurrentUser User
	}{
		CurrentUser:cuser,
	}
	ShowView(w, r, "help.html", data)
}

// 用户登录
func (app *Application) Login(w http.ResponseWriter, r *http.Request) {
	loginName := r.FormValue("loginName")
	password := r.FormValue("password")

	var flag bool
	for _, user := range users {
		if user.LoginName == loginName && user.Password == password {
			cuser = user
			flag = true
			break
		}
	}

	data := &struct {
		CurrentUser User
		Flag bool
	}{
		CurrentUser:cuser,
		Flag:false,
	}

	if flag{
		// 登录成功
		if cuser.IsAdmin == "T"{
			ShowView(w, r, "index.html", data)
		}else{

		result, err := app.Setup.FindAccInfoByEntityID(cuser.LoginName)
		var acc = service.Account{}
		json.Unmarshal(result, &acc)

		data := &struct {
			Edu service.Account
			CurrentUser User
			Msg string
			Flag bool
			History bool
		}{
			Edu:acc,
			CurrentUser:cuser,
			Msg:"",
			Flag:false,
			History:true,
		}

		if err != nil {
			data.Msg = err.Error()
			data.Flag = true
		}
		ShowView(w, r, "indexForUsers.html", data)

		}
	}else{
		// 登录失败
		data.Flag = true
		data.CurrentUser.LoginName = loginName
		ShowView(w, r, "login.html", data)
	}
}

// 用户登出
func (app *Application) LoginOut(w http.ResponseWriter, r *http.Request)  {
	cuser = User{}
	ShowView(w, r, "login.html", nil)
}

// 显示添加信息页面
func (app *Application) AddAccShow(w http.ResponseWriter, r *http.Request)  {
	data := &struct {
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
	}
	ShowView(w, r, "addEdu.html", data)
}

// 添加信息
func (app *Application) AddAcc(w http.ResponseWriter, r *http.Request)  {

	balance, err := strconv.ParseInt(r.FormValue("balance"), 10, 64)
	if err != nil{
		fmt.Printf("Illegal type of balance")
	}

	acc := service.Account{
		Name:r.FormValue("name"),
		Balance:balance,
		EntityID:r.FormValue("entityID"),
	}

	//	在登陆界面添加新用户
	newUser := User{LoginName:r.FormValue("entityID"), Password:"123456", IsAdmin:"F"}
	users = append(users, newUser)

	app.Setup.SaveAcc(acc)
	r.Form.Set("entity", acc.EntityID)
	app.FindByID(w, r)
}


func (app *Application) QueryPage2(w http.ResponseWriter, r *http.Request)  {
	data := &struct {
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
	}
	ShowView(w, r, "query2.html", data)
}

// 根据EntityID查询信息
func (app *Application) FindByID(w http.ResponseWriter, r *http.Request)  {
	entityID := r.FormValue("entityID")
	result, err := app.Setup.FindAccInfoByEntityID(entityID)
	var acc = service.Account{}
	json.Unmarshal(result, &acc)

	data := &struct {
		Edu service.Account
		CurrentUser User
		Msg string
		Flag bool
		History bool
	}{
		Edu:acc,
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
		History:true,
	}

	if err != nil {
		data.Msg = err.Error()
		data.Flag = true
	}

	ShowView(w, r, "queryResult.html", data)
}

// 根据EntityID查询History
func (app *Application) FindHistoryByID(w http.ResponseWriter, r *http.Request)  {
	entityID := cuser.LoginName
	result, err := app.Setup.FindAccInfoByEntityID(entityID)
	var acc = service.Account{}
	json.Unmarshal(result, &acc)

	data := &struct {
		Edu service.Account
		CurrentUser User
		Msg string
		Flag bool
		History bool
	}{
		Edu:acc,
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
		History:true,
	}

	if err != nil {
		data.Msg = err.Error()
		data.Flag = true
	}

	ShowView(w, r, "newHistory.html", data)
}

// 修改/添加新信息
func (app *Application) ModifyShow(w http.ResponseWriter, r *http.Request){
	// 根据证书编号与姓名查询信息
	entityID := r.FormValue("entityID")
	result, err := app.Setup.FindAccInfoByEntityID(entityID)

	var acc = service.Account{}
	json.Unmarshal(result, &acc)

	data := &struct {
		Edu service.Account
		CurrentUser User
		Msg string
		Flag bool
	}{
		Edu:acc,
		CurrentUser:cuser,
		Flag:true,
		Msg:"",
	}

	if err != nil {
		data.Msg = err.Error()
		data.Flag = true
	}

	ShowView(w, r, "modify.html", data)
}

// 修改/添加新信息
func (app *Application) Modify(w http.ResponseWriter, r *http.Request) {

	// 将 string 强制转换为 int64
	balance, err := strconv.ParseInt(r.FormValue("balance"), 10, 64)
	if err != nil{
		fmt.Printf("Illegal Input!")
	}

	acc := service.Account{
		Name:r.FormValue("name"),
		Balance:balance,
		EntityID:r.FormValue("entityID"),
	}
	app.Setup.ModifyAcc(acc)
	r.Form.Set("entityID", acc.EntityID)
	//r.Form.Set("name", acc.Name)
	app.FindByID(w, r)
}


func (app *Application) TransferShow(w http.ResponseWriter, r *http.Request){
	data := &struct {
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
	}

	if cuser.IsAdmin == "T"{
		ShowView(w, r, "transfer.html", data)
	}else{
		ShowView(w, r, "transferForUsers.html", data)
	}
}

func (app *Application) Transfer(w http.ResponseWriter, r *http.Request) {
	balance, err := strconv.ParseInt(r.FormValue("balance"), 10, 64)
	if err != nil{
		fmt.Printf("Illegal Input!")
	}
	from := r.FormValue("from")

	if cuser.IsAdmin == "F"{
		from = cuser.LoginName
	}
	app.Setup.TransferAcc(from,r.FormValue("to"),balance)
	r.Form.Set("entityID", from)
	app.FindByID(w, r)
}