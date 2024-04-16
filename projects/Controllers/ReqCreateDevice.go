package Controllers

import (
	"encoding/xml"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ReqCreateDevice struct {
	XMLName    xml.Name `xml:"http://npci.org/token/schema/ ReqCreateDevice"`
	ReqDetails ReqDetails
}

type ReqDetails struct {
	Type string `xml:"type,attr"`
	User User
}

type User struct {
	DeviceInfo DeviceInfo
	Strategies Strategies
}

type DeviceInfo struct {
	Tag []Tag `xml:"Tag"`
}

type Tag struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type Strategies struct {
	Strategy Strategy
}

type Strategy struct {
	Detail Detail
}

type Detail struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

func ReqCreateDeviceApi(c *gin.Context) {
	reqBody := RequestBodyLogger(c)
	log.Info("Request body:", reqBody)
	var req ReqCreateDevice
	err := xml.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		log.Error("Error decoding request body:", err)
		InternalServerErrorResponse(c, err.Error())
		return
	}
	RespCreateDevicefunc(c)
	log.Info("Successfully created request.")

}

type RespCreateDevice struct {
	XMLName    xml.Name `xml:"token:RespCreateDevice"`
	XMLNS      string   `xml:"xmlns:token,attr"`
	Head       Head     `xml:"Head"`
	ResDetails ResDetails
}

type Head struct {
	Ver   string `xml:"ver,attr"`
	Ts    string `xml:"ts,attr"`
	OrgID string `xml:"orgId,attr"`
	MsgID string `xml:"msgId,attr"`
}

type ResDetails struct {
	Type string `xml:"type,attr"`
	Resp Resp
	User UserResp
}

type Resp struct {
	ReqMsgID string `xml:"reqMsgId,attr"`
	Result   string `xml:"result,attr"`
	ErrCode  string `xml:"errCode,attr"`
	Msg      string `xml:"msg,attr"`
}

type UserResp struct {
	Details    Details `xml:"Details"`
	SMSDetails SMSDetails
}
type Details struct {
	Detail []Detail `xml:"Detail"`
}
type SMSDetails struct {
	Detail []Detail `xml:"Detail"`
}

func RespCreateDevicefunc(c *gin.Context) {
	resp := RespCreateDevice{
		XMLNS: "http://npci.org/token/schema/",
		Head: Head{
			Ver:   "1.0|2.0",
			Ts:    "",
			OrgID: "",
			MsgID: "",
		},
		ResDetails: ResDetails{
			Type: "Type1|Type2",
			Resp: Resp{
				ReqMsgID: "",
				Result:   "SUCCESS|FAILURE",
				ErrCode:  "",
				Msg:      "OTP SENT/CAN'T SENT",
			},
			User: UserResp{
				Details: Details{
					Detail: []Detail{
						{Name: "Id", Value: ""},
						{Name: "token", Value: ""}, //7739517851
						{Name: "expiresIn", Value: ""},
						{Name: "verified", Value: ""},
						{Name: "keyword", Value: ""},
					},
				},
				SMSDetails: SMSDetails{
					Detail: []Detail{
						{Name: "servieProviderNum", Value: ""},
						{Name: "content", Value: ""},
						{Name: "priority", Value: ""},
						{Name: "channel", Value: ""},
						{Name: "to", Value: ""},
					},
				},
			},
		},
	}

	xmlData, err := xml.MarshalIndent(resp, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling XML:", err)
		return
	}
	SuccessResponseXml(c, xmlData)
}
