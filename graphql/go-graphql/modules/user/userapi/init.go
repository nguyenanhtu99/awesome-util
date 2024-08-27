package userapi

import (
	"awesome-util/graphql/go-graphql/modules/user/userservice"
	"fmt"
	"github.com/labstack/echo/v4"
)

type userAPI struct {
	userSvc *userservice.Service
}

func (api *userAPI) initService() {
	api.userSvc = userservice.New()
}

func New(g *echo.Group) error {
	api := userAPI{}
	api.initService()

	g.GET("/:id", api.Get)

	return nil
}

func (api *userAPI) Get(c echo.Context) error {
	user, err := api.userSvc.Get(c.Request().Context(), c.Param("id"))
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	return c.JSON(200, user)
}
