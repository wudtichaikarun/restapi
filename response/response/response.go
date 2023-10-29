package response

import (
	"net/http"

	"github.com/wudtichaikarun/restapi"
)

const (
	SuccessStatus = "success"
	FailStatus    = "fail"
)

type SuccessResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	// Meta    *paginate.Meta `json:"meta,omitempty"`
}

type FailResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Message string `json:"message"`
}

type BindError struct {
	Message string
}

func (e *BindError) Error() string {
	return e.Message
}

func (f *FailResponse) Error() string {
	return f.Message
}

func AttachError(ctx restapi.Context, err error) {
	ctx.AttachError(err)
}

func NewFail(message string, code int) error {
	return &FailResponse{
		Code:    code,
		Status:  FailStatus,
		Message: message,
		Errors:  nil,
	}
}

func NotContent(ctx restapi.Context) {
	Success(ctx, &SuccessResponse{
		Code: http.StatusNoContent,
	})
}

func Created(ctx restapi.Context) {
	Success(ctx, &SuccessResponse{
		Code: http.StatusCreated,
	})
}

func Data(ctx restapi.Context, data interface{}) {
	Success(ctx, &SuccessResponse{
		Code: http.StatusOK,
		Data: data,
	})
}

func Success(ctx restapi.Context, r *SuccessResponse) {
	r.Status = SuccessStatus
	if r.Code == 0 {
		r.Code = http.StatusOK
	}
	if r.Message == "" {
		r.Message = ""
	}

	ctx.JSON(r.Code, r)
}

func UnprocessableEntity(ctx restapi.Context, errors interface{}) {
	Fail(ctx, &FailResponse{
		Code:    http.StatusUnprocessableEntity,
		Status:  FailStatus,
		Message: "",
		Errors:  errors,
	})
}

func BadRequest(ctx restapi.Context, message string) {
	Fail(ctx, NewFail(message, http.StatusBadRequest))
}

func Forbidden(ctx restapi.Context, message string) {
	Fail(ctx, NewFail(message, http.StatusForbidden))
}

func NotFound(ctx restapi.Context, message string) {
	Fail(ctx, NewFail(message, http.StatusNotFound))
}

func InternalServerError(ctx restapi.Context, message string) {
	Fail(ctx, NewFail(message, http.StatusInternalServerError))
}

func BadGateway(ctx restapi.Context, message string) {
	Fail(ctx, NewFail(message, http.StatusBadGateway))
}

func Fail(ctx restapi.Context, err error) {
	var response *FailResponse

	if e, ok := err.(*FailResponse); err != nil && ok {
		response = e
	} else {
		var message string
		if err == nil {
			message = "Server error occurred"
		} else {
			message = err.Error()
		}

		response = &FailResponse{
			Code:    500,
			Status:  FailStatus,
			Message: message,
		}
	}

	ctx.JSON(response.Code, response)
}
