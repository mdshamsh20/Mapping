package Controllers

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ReqSetCre struct {
	Head  HeadReqSetCre  `xml:"Head"`
	Txn   TxnReqSetCre   `xml:"Txn"`
	Payer PayerReqSetCre `xml:"Payer"`
}

type HeadReqSetCre struct {
	Ver      string `xml:"ver,attr"`
	Ts       string `xml:"ts,attr"`
	OrgId    string `xml:"orgId,attr"`
	MsgId    string `xml:"msgId,attr"`
	ProdType string `xml:"prodType,attr"`
}

type TxnReqSetCre struct {
	Id     string `xml:"id,attr"`
	Note   string `xml:"note,attr"`
	RefId  string `xml:"refId,attr"`
	RefUrl string `xml:"refUrl,attr"`
	Ts     string `xml:"ts,attr"`
	Type   string `xml:"type,attr"`
}

type PayerReqSetCre struct {
	Addr   string          `xml:"addr,attr"`
	Name   string          `xml:"name,attr"`
	SeqNum string          `xml:"seqNum,attr"`
	Type   string          `xml:"type,attr"`
	Code   string          `xml:"code,attr"`
	Device DeviceReqSetCre `xml:"Device"`
	Ac     AcReqSetCre     `xml:"Ac"`
}

type DeviceReqSetCre struct {
	Tags []TagReqSetCre `xml:"Tag"`
}

type TagReqSetCre struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type AcReqSetCre struct {
	Addr    string           `xml:"addrType,attr"`
	Details DetailsReqSetCre `xml:"Details"`
}

type DetailsReqSetCre struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type CredsReqSetCre struct {
	Cred []CredReqSetCre `xml:"Cred"`
}

type CredReqSetCre struct {
	Type    string        `xml:"type,attr"`
	SubType string        `xml:"subType,attr"`
	Data    DataReqSetCre `xml:"Data"`
}

type DataReqSetCre struct {
	Code  string `xml:"code,attr"`
	Ki    string `xml:"ki,attr"`
	Value string `xml:",chardata"`
}

func ReqSetCreAPI(c *gin.Context) {
	reqBody := RequestBodyLogger(c)
	log.Info("@ReqSetCreAPI Request body:", reqBody)
	var req ReqSetCre
	err := xml.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		log.Error("@ReqSetCreAPI Error decoding request body:", err)
		InternalServerErrorResponse(c, err.Error())
		return
	}
	errCreatResp := RespSetCreFunc(c)
	if errCreatResp != nil {
		log.Error("@ReqRegMobAPI Error: ", err)
		InternalServerErrorResponse(c, errCreatResp.Error())
		return
	}
	log.Info("@ReqSetCreAPI Successfully created request.")
}

type RespSetCre struct {
	XMLName xml.Name       `xml:"ns2:RespSetCre"`
	Ns2     string         `xml:"xmlns:ns2,attr"`
	Ns3     string         `xml:"xmlns:ns3,attr"`
	Head    HeadReqSetCre  `xml:"Head"`
	Resp    RespRespSetCre `xml:"Resp"`
	Txn     TxnReqSetCre   `xml:"Txn"`
}

// type Head struct {
// 	MsgId    string `xml:"msgId,attr"`
// 	OrgId    string `xml:"orgId,attr"`
// 	ProdType string `xml:"prodType,attr"`
// 	Ts       string `xml:"ts,attr"`
// 	Ver      string `xml:"ver,attr"`
// }

type RespRespSetCre struct {
	ReqMsgId string `xml:"reqMsgId,attr"`
	Result   string `xml:"result,attr"`
}

// type Txn struct {
// 	Id     string `xml:"id,attr"`
// 	Note   string `xml:"note,attr"`
// 	RefId  string `xml:"refId,attr"`
// 	RefUrl string `xml:"refUrl,attr"`
// 	Ts     string `xml:"ts,attr"`
// 	Type   string `xml:"type,attr"`
// }

func RespSetCreFunc(c *gin.Context) error {
	resp := RespSetCre{
		Ns2: "http://npci.org/token/schema/",
		Ns3: "http://npci.org/cm/schema/",
		Head: HeadReqSetCre{
			MsgId:    "1vUWishpRrWlqaOlWobc",
			OrgId:    "NPCI",
			ProdType: "PSO",
			Ts:       "2022-11-03T14:21:34+05:30",
			Ver:      "2.0",
		},
		Resp: RespRespSetCre{
			ReqMsgId: "001OnMMukQsOCTS7FsG8vJW8YmBD5Z4tSz9",
			Result:   "SUCCESS",
		},
		Txn: TxnReqSetCre{
			Id:     "001OnMMukQsOCTS7FsG8vJW8YmBD5Z4tSz9",
			Note:   "FETCHLIST",
			RefId:  "001OnMMukQsOCTS7FsG8vJW8YmBD5Z4tSz9",
			RefUrl: "http://upi",
			Ts:     "2022-11-03T04:51:34-04:00",
			Type:   "SetCre",
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
