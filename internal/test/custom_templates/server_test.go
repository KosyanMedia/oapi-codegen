package custom_templates

import (
	"github.com/labstack/echo/v4"
	"testing"
)

func TestCustomTemplates(t *testing.T) {
	t.Parallel()

	// It just needs to be successfully compiled
	var _ ServerInterface = &echoController{}
	var _ CustomArgument = "arg"
}

type echoController struct {
}

func (e echoController) GetEveryTypeOptional(ctx echo.Context) error {
	panic("implement me")
}

func (e echoController) CreateEveryTypeOptional(ctx echo.Context, params CreateEveryTypeOptionalParams) error {
	panic("implement me")
}

func (e echoController) GetSimple(ctx echo.Context) error {
	panic("implement me")
}

func (e echoController) GetWithArgs(ctx echo.Context, params GetWithArgsParams) error {
	panic("implement me")
}

func (e echoController) GetWithReferences(ctx echo.Context, globalArgument int64, argument Argument) error {
	panic("implement me")
}

func (e echoController) GetWithContentType(ctx echo.Context, contentType GetWithContentTypeParamsContentType) error {
	panic("implement me")
}

func (e echoController) GetReservedKeyword(ctx echo.Context) error {
	panic("implement me")
}

func (e echoController) CreateResource(ctx echo.Context, argument Argument) error {
	panic("implement me")
}

func (e echoController) CreateResource2(ctx echo.Context, inlineArgument int, params CreateResource2Params) error {
	panic("implement me")
}

func (e echoController) UpdateResource3(ctx echo.Context, pFallthrough int) error {
	panic("implement me")
}

func (e echoController) GetResponseWithReference(ctx echo.Context) error {
	panic("implement me")
}
