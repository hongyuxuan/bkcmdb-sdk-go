package types

import (
	"encoding/json"

	"github.com/golang-module/carbon"
)

type ClassificationList []ClassificationData

type ClassificationData struct {
	Id                   int                    `json:"id"`
	BkClassificationId   string                 `json:"bk_classification_id"`
	BkClassificationName string                 `json:"bk_classification_name"`
	BkClassificationType string                 `json:"bk_classification_type"`
	BkClassificationIcon string                 `json:"bk_classification_icon"`
	Account              string                 `json:"bk_supplier_account"`
	BkObjects            []ClassificationObject `json:"bk_objects"`
}

type ClassificationObject struct {
	Id                 int             `json:"id"`
	BkClassificationId string          `json:"bk_classification_id"`
	BkObjIcon          string          `json:"bk_obj_icon"`
	BkObjId            string          `json:"bk_obj_id"`
	BkObjName          string          `json:"bk_obj_name"`
	BkIshidden         bool            `json:"bk_ishidden"`
	Ispre              bool            `json:"ispre"`
	BkIspaused         bool            `json:"bk_ispaused"`
	Position           string          `json:"position"`
	Account            string          `json:"bk_supplier_account"`
	Description        string          `json:"description"`
	Creator            string          `json:"creator"`
	Modifier           string          `json:"modifier"`
	CreateTime         carbon.DateTime `json:"create_time"`
	LastTime           carbon.DateTime `json:"last_time"`
	ObjSortNumber      int             `json:"obj_sort_number"`
}

func (d *ClassificationList) ToJsonString() string {
	b, _ := json.Marshal(d)
	return string(b)
}

func (d *ClassificationList) ToJsonStringPretty() string {
	b, _ := json.MarshalIndent(d, "", "  ")
	return string(b)
}
