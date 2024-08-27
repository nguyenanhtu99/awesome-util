package userapi

import (
	"awesome-util/graphql/gqlgen/modules/user"
	"awesome-util/graphql/gqlgen/modules/user/userservice"
	"fmt"

	"github.com/labstack/echo/v4"
)

type userAPI struct {
	userSvc *userservice.Service
}

func (api *userAPI) initService() error {
	var err error
	api.userSvc, err = user.New()
	if err != nil {
		return fmt.Errorf("new user service, err: %w", err)
	}

	return nil
}

func New(g *echo.Group) error {
	api := userAPI{}
	if err := api.initService(); err != nil {
		return fmt.Errorf("init service, err: %w", err)
	}

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
