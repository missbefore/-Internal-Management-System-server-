package controller

import (
	"github.com/kataras/iris"
	"net/login/services"
	"net/login/datamodels"
)

type UsersController struct {
	Ctx iris.Context

	Service services.UserService
}

func (c *UserController) Get() (results []datamodels.User) {
	return c.Service.GetAll()
}

func (c *UserController) GetBy(id int64) (user datamodels.User, found bool)  {
	u, found := c.Service.GetByID(id)
	if !found {
		c.Ctx.Values().Set("message", "找不到这个用户！")
	}
	return u, found
}

func (c *UserController) PutBy(id int64) (datamodels.User, error) {
	u := datamodels.User{}

	if err := c.Ctx.ReadForm(&u); err != nil {
		return u, err
	}

	return c.Service.Update(id, u)
}

func (c *UserController) DeleteBy(id int64) interface{} {
	wasDel := c.Service.DeleteByID(id)
	if wasDel {
		return map[string]interface{}{"delete": id}
	}

	return iris.StatusBadRequest
}