package types

import (
	"encoding/json"

	"github.com/golang-module/carbon"
)

type Object struct {
	Account            string `json:"bk_supplier_account"`
	BkClassificationId string `json:"bk_classification_id"`
	BkObjIcon          string `json:"bk_obj_icon"`
	BkObjId            string `json:"bk_obj_id"`
	BkObjName          string `json:"bk_obj_name"`
	Username           string `json:"userName"`
}

type ObjectAttrGroup struct {
	Id           int    `json:"id"`
	BkBizId      int    `json:"bk_biz_id"`
	BkGroupId    string `json:"bk_group_id"`
	BkGroupName  string `json:"bk_group_name"`
	BkGroupIndex int    `json:"bk_group_index"`
	BkObjId      string `json:"bk_obj_id"`
	Account      string `json:"bk_supplier_account"`
	BkIsdefault  bool   `json:"bk_isdefault"`
	Ispre        bool   `json:"ispre"`
	IsCollapse   bool   `json:"is_collapse"`
}

type ObjectAttrGroupList []ObjectAttrGroup

func (o *ObjectAttrGroupList) ToJsonString() string {
	b, _ := json.Marshal(o)
	return string(b)
}

func (o *ObjectAttrGroupList) ToJsonStringPretty() string {
	b, _ := json.MarshalIndent(o, "", "  ")
	return string(b)
}

type ObjectAttr struct {
	Id                  int              `json:"id,omitempty"`
	BkBizId             int              `json:"bk_biz_id,omitempty"`
	BkPropertyName      string           `json:"bk_property_name"`
	BkPropertyId        string           `json:"bk_property_id"`
	BkPropertyGroup     string           `json:"bk_property_group"`
	BkPropertyGroupName string           `json:"bk_property_group_name,omitempty"`
	BkPropertyIndex     int              `json:"bk_property_index,omitempty"`
	Unit                string           `json:"unit"`
	Placeholder         string           `json:"placeholder"`
	BkPropertyType      string           `json:"bk_property_type"`
	Editable            bool             `json:"editable"`
	Ispre               bool             `json:"ispre"`
	Isrequired          bool             `json:"isrequired"`
	Isreadonly          bool             `json:"isreadonly,omitempty"`
	Isonly              bool             `json:"isonly,omitempty"`
	BkIssystem          bool             `json:"bk_issystem,omitempty"`
	BkIsapi             bool             `json:"bk_isapi,omitempty"`
	IsMultiple          bool             `json:"ismultiple"`
	Option              interface{}      `json:"option"`
	Creator             string           `json:"creator"`
	BkObjId             string           `json:"bk_obj_id"`
	Account             string           `json:"bk_supplier_account"`
	Description         string           `json:"description,omitempty"`
	CreateTime          *carbon.DateTime `json:"create_time,omitempty"`
	LastTime            *carbon.DateTime `json:"last_time,omitempty"`
	BkTemplateId        int              `json:"bk_template_id,omitempty"`
}

type ObjectAttrList []ObjectAttr

func (o *ObjectAttrList) ToJsonString() string {
	b, _ := json.Marshal(o)
	return string(b)
}

func (o *ObjectAttrList) ToJsonStringPretty() string {
	b, _ := json.MarshalIndent(o, "", "  ")
	return string(b)
}

type ObjectAssociation struct {
	Id            int         `json:"id,omitempty"`
	Account       string      `json:"bk_supplier_account,omitempty"`
	BkObjAsstId   string      `json:"bk_obj_asst_id"`
	BkObjAsstName string      `json:"bk_obj_asst_name"`
	BkObjId       string      `json:"bk_obj_id"`
	BkAsstObjId   string      `json:"bk_asst_obj_id"`
	BkAsstId      string      `json:"bk_asst_id"`
	Mapping       string      `json:"mapping"`
	OnDelete      string      `json:"on_delete,omitempty"`
	Ispre         interface{} `json:"ispre,omitempty"`
}

type ObjectAssociationList []ObjectAssociation

func (o *ObjectAssociationList) ToJsonString() string {
	b, _ := json.Marshal(o)
	return string(b)
}

func (o *ObjectAssociationList) ToJsonStringPretty() string {
	b, _ := json.MarshalIndent(o, "", "  ")
	return string(b)
}
