package types

import "encoding/json"

type InstAssociation struct {
	Id           int64  `json:"id,omitempty"`
	Account      string `json:"bk_supplier_account,omitempty"`
	BkAsstInstId int64  `json:"bk_asst_inst_id"`
	BkAsstObjId  string `json:"bk_asst_obj_id"`
	BkAsstId     string `json:"bk_asst_id"`
	BkObjAsstId  string `json:"bk_obj_asst_id"`
	BkObjId      string `json:"bk_obj_id"`
	BkInstId     int64  `json:"bk_inst_id"`
}

type InstAssociationList []InstAssociation

func (i *InstAssociationList) ToJsonString() string {
	b, _ := json.Marshal(i)
	return string(b)
}

func (i *InstAssociationList) ToJsonStringPretty() string {
	b, _ := json.MarshalIndent(i, "", "  ")
	return string(b)
}
