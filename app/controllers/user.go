package controllers

import "github.com/revel/revel"
import "fmt"
import "emspace/repositories"

type User struct {
	*revel.Controller
}

func (c User) GetUser(id int) revel.Result {
	u := repositories.GetUsers("user_id", fmt.Sprintf("%v", id))
	return c.RenderJson(u)
}
