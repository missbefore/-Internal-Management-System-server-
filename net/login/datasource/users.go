package datasource

import (
	"net/login/datamodels"
	"github.com/kataras/iris/core/errors"
)

type Engine uint32

const (
	Memory Engine = iota
	Bolt
	MySQL
)

func LoadUsers(engine Engine) (map[int64]datamodels.User, error)  {
	if engine != Memory {
		return nil, errors.New("走简单的方法")
	}

	return make(map[int64]datamodels.User), nil
}
