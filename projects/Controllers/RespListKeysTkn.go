package Controllers

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ReqListKeysTkn struct {
	XMLName xml.Name `xml:"http://npci.org/token/schema/ ReqListKeys"`
	Head    HeadTkn  `xml:"Head"`
	Txn     TxnTkn   `xml:"Txn"`
	Creds   Creds    `xml:"Creds"`
}

type HeadTkn struct {
	Ver      string `xml:"ver,attr"`
	Ts       string `xml:"ts,attr"`
	OrgId    string `xml:"orgId,attr"`
	MsgId    string `xml:"msgId,attr"`
	ProdType string `xml:"prodType,attr"`
}

type TxnTkn struct {
	Id       string `xml:"id,attr"`
	Note     string `xml:"note,attr"`
	RefId    string `xml:"refId,attr"`
	RefUrl   string `xml:"refUrl,attr"`
	Ts       string `xml:"ts,attr"`
	Type     string `xml:"type,attr"`
	PspOrgId string `xml:"pspOrgId,attr"`
	Mobile   string `xml:"mobile,attr"`
	TspId    string `xml:"tspId,attr"`
}

type Creds struct {
	Cred Cred `xml:"Cred"`
}

type Cred struct {
	Type    string `xml:"type,attr"`
	SubType string `xml:"subType,attr"`
	Data    Data   `xml:"Data"`
}

type Data struct {
	Code string `xml:"code,attr"`
	Ki   string `xml:"ki,attr"`
}

func RespListKeysTknApi(c *gin.Context) {
	reqBody := RequestBodyLogger(c)
	log.Info("@RespListKeysTknApi Request body:", reqBody)
	var req ReqListKeysTkn
	err := xml.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		log.Error("@RespListKeysTknApi Error decoding request body:", err)
		InternalServerErrorResponse(c, err.Error())
		return
	}
	errInCrResp := RespListKeysTknfunc(c)
	if errInCrResp != nil {
		log.Error("@RespListKeysTknApi Error: ", err)
		InternalServerErrorResponse(c, errInCrResp.Error())
		return
	}
	log.Info("@RespListKeysTknApi Successfully created request.")
}

type RespTkn struct {
	ReqMsgID string `xml:"reqMsgId,attr"`
	Result   string `xml:"result,attr"`
}
type RespListKeysTkn struct {
	XMLName  xml.Name `xml:"ns2:RespListKeys"`
	Head     Head     `xml:"Head"`
	Resp     RespTkn  `xml:"Resp"`
	Txn      Txn      `xml:"Txn"`
	KeyList  KeyList  `xml:"keyList"`
	XmlnsNs2 string   `xml:"xmlns:ns2,attr"`
}

func RespListKeysTknfunc(c *gin.Context) error {
	resp := RespListKeysTkn{
		Head: Head{
			Ver:   "2.0",
			Ts:    time.Now().Format(time.RFC3339),
			OrgID: "NPCI",
			MsgID: "1vUWishpRrWikzn8WWPp",
		},
		Resp: RespTkn{
			ReqMsgID: "0018RbHc6yphC2o8xsiVJKFtmb768NC3z8N",
			Result:   "SUCCESS",
		},
		Txn: Txn{
			Id:       "0018RbHc6yphC2o8xsiVJKFtmb768NC3z8N",
			Note:     "NA",
			RefId:    "0018RbHc6yphC2o8xsiVJKFtmb768NC3z8N",
			RefUrl:   "http://upi",
			Ts:       time.Now().Format(time.RFC3339),
			Type:     "GetToken",
			PspOrgId: "159022",
		},
		KeyList: KeyList{
			Keys: []Key{
				{
					Code:  "NPCI",
					Owner: "NPCI",
					Type:  "CLF",
					Ki:    "20220817",
					KeyValue: KeyValue{
						XSI:   "http://www.w3.org/2001/XMLSchema-instance",
						XS:    "http://www.w3.org/2001/XMLSchema",
						Type:  "xs:string",
						Value: "E2jSxkturCtGWjDqI3MnOyyNbVOgDYIjD7 FxVYsRI7E=",
					},
				},
			},
		},
		XmlnsNs2: "http://npci.org/token/schema/",
	}

	output, err := xml.MarshalIndent(resp, "", "    ")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}
	SuccessResponseXml(c, output)
	return nil
}
