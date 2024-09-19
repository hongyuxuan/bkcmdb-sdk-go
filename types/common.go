package types

import "encoding/json"

type BaseResponse struct {
	Code       int64       `json:"bk_error_code"`
	Message    string      `json:"bk_error_msg"`
	Permission interface{} `json:"permission"`
	Result     bool        `json:"result"`
}

type PostResponse struct {
	BaseResponse
	Data map[string]interface{} `json:"data"`
}

type Rule struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

type Conditions struct {
	Condition string `json:"condition"`
	Rules     []Rule `json:"rules,omitempty"`
}

type Page struct {
	Limit int    `json:"limit"`
	Start int    `json:"start"`
	Sort  string `json:"sort"`
}

type ListOption struct {
	Fields     []string    `json:"fields,omitempty"`
	Page       *Page       `json:"page,omitempty"`
	Conditions *Conditions `json:"conditions,omitempty"`
}

type ListDataInfo struct {
	Info []interface{} `json:"info"`
}

type ListDataCount struct {
	Count int64 `json:"count"`
}

type ListInfoResponse struct {
	BaseResponse
	Data ListDataInfo `json:"data"`
}

type ListCountResponse struct {
	BaseResponse
	Data ListDataCount `json:"data"`
}

type ListResponse struct {
	ListDataInfo
	ListDataCount
}

func (r *ListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListResponse) ToJsonStringPretty() string {
	b, _ := json.MarshalIndent(r, "", "  ")
	return string(b)
}

type ListInstAssociationRequest struct {
	Condition InstAssociationCondition `json:"condition"`
	BkObjId   string                   `json:"bk_obj_id"`
}

type InstAssociationCondition struct {
	BkObjId      string `json:"bk_obj_id,omitempty"`
	BkInstId     int64  `json:"bk_inst_id,omitempty"`
	BkObjAsstId  string `json:"bk_obj_asst_id,omitempty"`
	BkAsstObjId  string `json:"bk_asst_obj_id,omitempty"`
	BkAsstInstId int64  `json:"bk_asst_inst_id,omitempty"`
}
