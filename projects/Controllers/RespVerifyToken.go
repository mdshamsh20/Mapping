package Controllers

import (
	"encoding/xml"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ReqVerifyToken struct {
	XMLName    xml.Name `xml:"http://npci.org/token/schema/ ReqVerifyToken"`
	ReqDetails ReqDetails
}

type UserVerifyTocken struct {
	Details DetailTc `xml:"Details"`
}
type DetailTc struct {
	Details []Detail `xml:"Detail"`
}

type RespVerifyToken struct {
	XMLName    xml.Name         `xml:"token:RespVerifyToken"`
	XMLNS      string           `xml:"xmlns:token,attr"`
	Head       Head             `xml:"Head"`
	ResDetails ResDetailsTocken `xml:"ResDetails"`
}
type ResDetailsTocken struct {
	Type string `xml:"type,attr"`
	Resp Resp
	User UserVerifyTocken
}

func RespVerifyTokenApi(c *gin.Context) {
	reqBody := RequestBodyLogger(c)
	log.Info("@RespVerifyTokenApi Request body:", reqBody)
	var req ReqVerifyToken
	err := xml.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		log.Error("@RespVerifyTokenApi Error decoding request body:", err)
		InternalServerErrorResponse(c, err.Error())
		return
	}
	RespVerifyTokenfunc(c)
	log.Info("@RespVerifyTokenApi Successfully created request.")
}

func RespVerifyTokenfunc(c *gin.Context) {
	resp := RespVerifyToken{
		XMLNS: "http://npci.org/token/schema/",
		Head: Head{
			Ver:   "1.0|2.0",
			Ts:    "",
			OrgID: "",
			MsgID: "",
		},
		ResDetails: ResDetailsTocken{
			Type: "Type1|Type2",
			Resp: Resp{
				ReqMsgID: "",
				Result:   "SUCCESS|FAILURE",
				ErrCode:  "",
				Msg:      "OTP SENT/CAN'T SENT",
			},
			User: UserVerifyTocken{
				Details: DetailTc{
					[]Detail{
						{Name: "Id", Value: ""},
						{Name: "token", Value: ""},
						{Name: "verified", Value: ""},
						{Name: "phone", Value: ""},
						{Name: "preExistingUser", Value: "TRUE|FALSE"},
						{Name: "name", Value: "SomeUser"},
						{Name: "vpa", Value: "729349823@oktoken"},
					}},
			},
		},
	}

	output, err := xml.MarshalIndent(resp, "", "    ")
	if err != nil {
		fmt.Printf("Error marshaling XML: %v", err)
		return
	}
	SuccessResponseXml(c, output)
}
