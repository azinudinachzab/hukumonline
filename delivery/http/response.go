package delivery

import (
	"net/http"

	"github.com/azinudinachzab/hukumonline/model"
	"github.com/azinudinachzab/hukumonline/pkg/errs"
	"github.com/azinudinachzab/hukumonline/pkg/strcase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type httpResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func responseData(c *gin.Context, status int, res httpResponse) {
	if status == 0 {
		status = http.StatusOK
	}
	c.JSON(status, gin.H{"message": res.Message, "data": res.Data})
	return
}

func responseError(c *gin.Context, err error) {
	cerr, ok := err.(*errs.Error)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if cerr.Code == model.ECodeValidateFail && cerr.Err != nil && cerr.Attributes == nil {
		data := convertValidatorErrToAttribute(cerr)
		c.JSON(http.StatusBadRequest, gin.H{"message": data.Message, "code": data.Code})
		return
	}

	stsCode := http.StatusBadRequest
	if cerr.Code == model.ECodeInternal {
		stsCode = http.StatusInternalServerError
	}

	c.JSON(stsCode, gin.H{"message": cerr.Message, "code": cerr.Code})
}

func convertValidatorErrToAttribute(cerr *errs.Error) *errs.Error {
	if _, ok := cerr.Err.(*validator.InvalidValidationError); ok {
		return errs.New(model.ECodeInternal, "unknown error")
	}

	attrs := make([]errs.Attribute, 0)
	for _, fe := range cerr.Err.(validator.ValidationErrors) {
		fld := strcase.ToSnakeCase(fe.Field())
		msg := tagToMsg(fld, fe)
		attrs = append(attrs, errs.Attribute{
			Field:   fld,
			Message: msg,
		})
	}

	return errs.NewWithAttribute(cerr.Code, cerr.Message, attrs)
}

func tagToMsg(field string, fe validator.FieldError) string {
	switch fe.Tag() {
	case "min":
		return "less than minimum length"
	case "max":
		return "over than maximum length max: " + fe.Param()
	case "required":
		return "cannot be empty."
	case "required_if":
		return "cannot be empty if field " + fe.Param()
	case "oneof":
		return "value must be one of: " + fe.Param()
	case "len":
		return "value must have length " + fe.Param()
	}

	return field + " is failed to validate"
}
