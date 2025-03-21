// Package tasks provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// Task defines model for Task.
type Task struct {
	Id     *uint  `json:"id"`
	IsDone *bool  `json:"is_done"`
	Task   string `json:"task"`
}

// UpdateTaskJSONBody defines parameters for UpdateTask.
type UpdateTaskJSONBody map[string]interface{}

// PostTasksJSONRequestBody defines body for PostTasks for application/json ContentType.
type PostTasksJSONRequestBody = Task

// UpdateTaskJSONRequestBody defines body for UpdateTask for application/json ContentType.
type UpdateTaskJSONRequestBody UpdateTaskJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Получить все задачи
	// (GET /tasks)
	GetTasks(ctx echo.Context) error
	// Создать новую задачу
	// (POST /tasks)
	PostTasks(ctx echo.Context) error
	// Удалить задачу по ID
	// (DELETE /tasks/{id})
	DeleteTask(ctx echo.Context, id uint) error
	// Получить задачу по ID
	// (GET /tasks/{id})
	GetTaskByID(ctx echo.Context, id uint) error
	// Обновить задачу по ID
	// (PATCH /tasks/{id})
	UpdateTask(ctx echo.Context, id uint) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetTasks converts echo context to params.
func (w *ServerInterfaceWrapper) GetTasks(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTasks(ctx)
	return err
}

// PostTasks converts echo context to params.
func (w *ServerInterfaceWrapper) PostTasks(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostTasks(ctx)
	return err
}

// DeleteTask converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteTask(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id uint

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteTask(ctx, id)
	return err
}

// GetTaskByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetTaskByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id uint

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTaskByID(ctx, id)
	return err
}

// UpdateTask converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateTask(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id uint

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateTask(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/tasks", wrapper.GetTasks)
	router.POST(baseURL+"/tasks", wrapper.PostTasks)
	router.DELETE(baseURL+"/tasks/:id", wrapper.DeleteTask)
	router.GET(baseURL+"/tasks/:id", wrapper.GetTaskByID)
	router.PATCH(baseURL+"/tasks/:id", wrapper.UpdateTask)

}

type GetTasksRequestObject struct {
}

type GetTasksResponseObject interface {
	VisitGetTasksResponse(w http.ResponseWriter) error
}

type GetTasks200JSONResponse []Task

func (response GetTasks200JSONResponse) VisitGetTasksResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostTasksRequestObject struct {
	Body *PostTasksJSONRequestBody
}

type PostTasksResponseObject interface {
	VisitPostTasksResponse(w http.ResponseWriter) error
}

type PostTasks200JSONResponse Task

func (response PostTasks200JSONResponse) VisitPostTasksResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostTasks422Response struct {
}

func (response PostTasks422Response) VisitPostTasksResponse(w http.ResponseWriter) error {
	w.WriteHeader(422)
	return nil
}

type PostTasks500Response struct {
}

func (response PostTasks500Response) VisitPostTasksResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type DeleteTaskRequestObject struct {
	Id uint `json:"id"`
}

type DeleteTaskResponseObject interface {
	VisitDeleteTaskResponse(w http.ResponseWriter) error
}

type DeleteTask204Response struct {
}

func (response DeleteTask204Response) VisitDeleteTaskResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type DeleteTask404Response struct {
}

func (response DeleteTask404Response) VisitDeleteTaskResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type DeleteTask500Response struct {
}

func (response DeleteTask500Response) VisitDeleteTaskResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type GetTaskByIDRequestObject struct {
	Id uint `json:"id"`
}

type GetTaskByIDResponseObject interface {
	VisitGetTaskByIDResponse(w http.ResponseWriter) error
}

type GetTaskByID200JSONResponse Task

func (response GetTaskByID200JSONResponse) VisitGetTaskByIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetTaskByID404Response struct {
}

func (response GetTaskByID404Response) VisitGetTaskByIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetTaskByID500Response struct {
}

func (response GetTaskByID500Response) VisitGetTaskByIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type UpdateTaskRequestObject struct {
	Id   uint `json:"id"`
	Body *UpdateTaskJSONRequestBody
}

type UpdateTaskResponseObject interface {
	VisitUpdateTaskResponse(w http.ResponseWriter) error
}

type UpdateTask200JSONResponse Task

func (response UpdateTask200JSONResponse) VisitUpdateTaskResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type UpdateTask400Response struct {
}

func (response UpdateTask400Response) VisitUpdateTaskResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type UpdateTask404Response struct {
}

func (response UpdateTask404Response) VisitUpdateTaskResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type UpdateTask500Response struct {
}

func (response UpdateTask500Response) VisitUpdateTaskResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Получить все задачи
	// (GET /tasks)
	GetTasks(ctx context.Context, request GetTasksRequestObject) (GetTasksResponseObject, error)
	// Создать новую задачу
	// (POST /tasks)
	PostTasks(ctx context.Context, request PostTasksRequestObject) (PostTasksResponseObject, error)
	// Удалить задачу по ID
	// (DELETE /tasks/{id})
	DeleteTask(ctx context.Context, request DeleteTaskRequestObject) (DeleteTaskResponseObject, error)
	// Получить задачу по ID
	// (GET /tasks/{id})
	GetTaskByID(ctx context.Context, request GetTaskByIDRequestObject) (GetTaskByIDResponseObject, error)
	// Обновить задачу по ID
	// (PATCH /tasks/{id})
	UpdateTask(ctx context.Context, request UpdateTaskRequestObject) (UpdateTaskResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetTasks operation middleware
func (sh *strictHandler) GetTasks(ctx echo.Context) error {
	var request GetTasksRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTasks(ctx.Request().Context(), request.(GetTasksRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTasks")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTasksResponseObject); ok {
		return validResponse.VisitGetTasksResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostTasks operation middleware
func (sh *strictHandler) PostTasks(ctx echo.Context) error {
	var request PostTasksRequestObject

	var body PostTasksJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostTasks(ctx.Request().Context(), request.(PostTasksRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostTasks")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostTasksResponseObject); ok {
		return validResponse.VisitPostTasksResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// DeleteTask operation middleware
func (sh *strictHandler) DeleteTask(ctx echo.Context, id uint) error {
	var request DeleteTaskRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteTask(ctx.Request().Context(), request.(DeleteTaskRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteTask")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteTaskResponseObject); ok {
		return validResponse.VisitDeleteTaskResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetTaskByID operation middleware
func (sh *strictHandler) GetTaskByID(ctx echo.Context, id uint) error {
	var request GetTaskByIDRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTaskByID(ctx.Request().Context(), request.(GetTaskByIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTaskByID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTaskByIDResponseObject); ok {
		return validResponse.VisitGetTaskByIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// UpdateTask operation middleware
func (sh *strictHandler) UpdateTask(ctx echo.Context, id uint) error {
	var request UpdateTaskRequestObject

	request.Id = id

	var body UpdateTaskJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.UpdateTask(ctx.Request().Context(), request.(UpdateTaskRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateTask")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(UpdateTaskResponseObject); ok {
		return validResponse.VisitUpdateTaskResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
