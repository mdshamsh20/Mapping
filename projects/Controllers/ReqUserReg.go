package Controllers

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ReqUserReg struct {
	XMLName    xml.Name             `xml:"ReqUserReg"`
	Head       HeadReqUserReg       `xml:"Head"`
	ReqDetails ReqDetailsReqUserReg `xml:"ReqDetails"`
}

type HeadReqUserReg struct {
	Ver   string `xml:"ver,attr"`
	Ts    string `xml:"ts,attr"`
	OrgId string `xml:"orgId,attr"`
	MsgId string `xml:"msgId,attr"`
}

type ReqDetailsReqUserReg struct {
	Type string         `xml:"type,attr"`
	User UserReqUserReg `xml:"User"`
}

type UserReqUserReg struct {
	Device  DeviceReqUserReg   `xml:"Device"`
	Details []DetailReqUserReg `xml:"Details>Detail"`
}

type DeviceReqUserReg struct {
	Tags []TagReqUserReg `xml:"Tag"`
}

type TagReqUserReg struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type DetailReqUserReg struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

func ReqUserRegAPI(c *gin.Context) {
	reqBody := RequestBodyLogger(c)
	log.Info("@ReqUserRegAPI Request body:", reqBody)
	var req ReqUserReg
	err := xml.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		log.Error("@ReqUserRegAPI Error decoding request body:", err)
		InternalServerErrorResponse(c, err.Error())
		return
	}
	errCreatResp := ReqUserRegfunc(c)
	if errCreatResp != nil {
		log.Error("@ReqUserRegAPI Error: ", err)
		InternalServerErrorResponse(c, errCreatResp.Error())
		return
	}
	log.Info("@ReqUserRegAPI Successfully created request.")
}

type RespUserRegHd struct {
	XMLName    xml.Name             `xml:"token:RespUserReg"`
	XMLNS      string               `xml:"xmlns:token,attr"`
	Head       HeadReqUserReg       `xml:"Head"`
	ResDetails ResDetailsReqUserReg `xml:"ResDetails"`
}

// type Head struct {
// 	Ver   string `xml:"ver,attr"`
// 	Ts    string `xml:"ts,attr"`
// 	OrgId string `xml:"orgId,attr"`
// 	MsgId string `xml:"msgId,attr"`
// }

type ResDetailsReqUserReg struct {
	Type string      `xml:"type,attr"`
	Resp RespUserReg `xml:"Resp"`
}

type RespUserReg struct {
	Result  string `xml:"result,attr"`
	ErrCode string `xml:"errCode,attr"`
	Msg     string `xml:"msg,attr"`
}

func ReqUserRegfunc(c *gin.Context) error {
	resp := RespUserRegHd{
		XMLNS: "http://npci.org/token/schema/",
		Head: HeadReqUserReg{
			Ver:   "2.0",
			Ts:    "1656482881888",
			OrgId: "gaR2QLXH1p1YEOKP",
			MsgId: "001unSCnXWJ2v0ji3BnM1fwpimuhlSIrJeJ",
		},
		ResDetails: ResDetailsReqUserReg{
			Type: "Type2",
			Resp: RespUserReg{
				Result:  "SUCCESS",
				ErrCode: "00",
				Msg:     "Transaction Successful",
			},
		},
	}
	xmlData, err := xml.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Error("Failed to generate XML response: ", err)
		return err
	}

	SuccessResponseXml(c, xmlData)
	return nil
}
