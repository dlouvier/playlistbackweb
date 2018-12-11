package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/buger/jsonparser"
	"github.com/dlouvier/playlistbackweb/models"
	"github.com/dlouvier/playlistbackweb/spotify"
)

// SpotifyController to work around with */spotify/* requests
type SpotifyController struct {
	beego.Controller
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (p *SpotifyController) Get() {

	uid := p.GetString(":uid")
	fmt.Println("Hola")
	fmt.Println(uid)
	response := spotify.GetTracks(uid)

	var pl []models.Track

	jsonparser.ArrayEach([]byte(response), func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		var track models.Track
		track.Title, _ = jsonparser.GetString(value, "track", "name")
		jsonparser.ArrayEach(value, func(lala []byte, dataType jsonparser.ValueType, offset int, err error) {
			t, _ := jsonparser.GetString(lala, "name") // Artist es un array. Una canción puede tener multiples artistas
			track.Author = append(track.Author, t)
		}, "track", "artists")
		pl = models.AddTrack(pl, track)
	}, "items")
	fmt.Println(pl)

	p.Data["json"] = pl

	p.ServeJSON()
}

// // @Title Update
// // @Description update the user
// // @Param	uid		path 	string	true		"The uid you want to update"
// // @Param	body		body 	models.User	true		"body for user content"
// // @Success 200 {object} models.User
// // @Failure 403 :uid is not int
// // @router /:uid [put]
// func (u *UserController) Put() {
// 	uid := u.GetString(":uid")
// 	if uid != "" {
// 		var user models.User
// 		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
// 		uu, err := models.UpdateUser(uid, &user)
// 		if err != nil {
// 			u.Data["json"] = err.Error()
// 		} else {
// 			u.Data["json"] = uu
// 		}
// 	}
// 	u.ServeJSON()
// }

// // @Title Delete
// // @Description delete the user
// // @Param	uid		path 	string	true		"The uid you want to delete"
// // @Success 200 {string} delete success!
// // @Failure 403 uid is empty
// // @router /:uid [delete]
// func (u *UserController) Delete() {
// 	uid := u.GetString(":uid")
// 	models.DeleteUser(uid)
// 	u.Data["json"] = "delete success!"
// 	u.ServeJSON()
// }

// // @Title Login
// // @Description Logs user into the system
// // @Param	username		query 	string	true		"The username for login"
// // @Param	password		query 	string	true		"The password for login"
// // @Success 200 {string} login success
// // @Failure 403 user not exist
// // @router /login [get]
// func (u *UserController) Login() {
// 	username := u.GetString("username")
// 	password := u.GetString("password")
// 	if models.Login(username, password) {
// 		u.Data["json"] = "login success"
// 	} else {
// 		u.Data["json"] = "user not exist"
// 	}
// 	u.ServeJSON()
// }

// // @Title logout
// // @Description Logs out current logged in user session
// // @Success 200 {string} logout success
// // @router /logout [get]
// func (u *UserController) Logout() {
// 	u.Data["json"] = "logout success"
// 	u.ServeJSON()
// }
