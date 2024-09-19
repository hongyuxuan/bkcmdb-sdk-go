package errorx

import (
	"fmt"
	"net/http"

	"github.com/hongyuxuan/bkcmdb-sdk-go/core/constant"
)

type BkcmdbError struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (e *BkcmdbError) Error() string {
	return e.Message
}

func NewError(code int64, message string, data interface{}) error {
	return &BkcmdbError{Code: code, Message: message, Data: data}
}

func NewDefaultError(message string, a ...any) error {
	return &BkcmdbError{Code: http.StatusInternalServerError, Message: fmt.Sprintf(message, a...)}
}

func IsDuplicateAssociationError(err error) bool {
	defer func() {
		if er := recover(); er != nil {
			return
		}
	}()
	e := err.(*BkcmdbError)
	return e.Code == constant.ERR_DUPLICATE_ASSOCIATION_ERROR
}

func IsMissingBkObjIdError(err error) bool {
	defer func() {
		if er := recover(); er != nil {
			return
		}
	}()
	e := err.(*BkcmdbError)
	return e.Code == constant.ERR_MISSING_BKOBJID
}

func IsUniqKeyError(err error) bool {
	defer func() {
		if er := recover(); er != nil {
			return
		}
	}()
	e := err.(*BkcmdbError)
	return e.Code == constant.ERR_UNIQ_KEY_ERROR
}
