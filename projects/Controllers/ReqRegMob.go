package Controllers

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ReqRegMob struct {
	XMLName xml.Name `xml:"ReqRegMob"`
	Head    Head     `xml:"Head"`
	Txn     Txn      `xml:"Txn"`
	Payer   Payer    `xml:"Payer"`
}

type HeadReqRegMob struct {
	Ver       string `xml:"ver,attr"`
	Timestamp string `xml:"ts,attr"`
	OrgID     string `xml:"orgId,attr"`
	MsgID     string `xml:"msgId,attr"`
	ProdType  string `xml:"prodType,attr"`
}

type TxnReqRegMob struct {
	ID     string `xml:"id,attr"`
	Note   string `xml:"note,attr"`
	RefID  string `xml:"refId,attr"`
	RefURL string `xml:"refUrl,attr"`
	Type   string `xml:"type,attr"`
	Op     string `xml:"op,attr"`
}

type PayerReqRegMob struct {
	Addr    string  `xml:"addr,attr"`
	Name    string  `xml:"name,attr"`
	SeqNum  int     `xml:"seqNum,attr"`
	Type    string  `xml:"type,attr"`
	KycType string  `xml:"kycType,attr"`
	Device  Device  `xml:"Device"`
	Ac      Account `xml:"Ac"`
}

type DeviceReqRegMob struct {
	Tags []Tag `xml:"Tag"`
}

type TagReqRegMob struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type AccountReqRegMob struct {
	Addr   string  `xml:"addrType,attr"`
	Detail []Field `xml:"Detail"`
}

type Field struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

func ReqRegMobAPI(c *gin.Context) {
	reqBody := RequestBodyLogger(c)
	log.Info("@ReqRegMobAPI Request body:", reqBody)
	var req ReqRegMob
	err := xml.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		log.Error("@ReqRegMobAPI Error decoding request body:", err)
		InternalServerErrorResponse(c, err.Error())
		return
	}
	errCreatResp := ReqRegMobfunc(c)
	if errCreatResp != nil {
		log.Error("@ReqRegMobAPI Error: ", err)
		InternalServerErrorResponse(c, errCreatResp.Error())
		return
	}
	log.Info("@ReqRegMobAPI Successfully created request.")
}

type HeadRegMob struct {
	Ver      string `xml:"ver,attr"`
	Ts       string `xml:"ts,attr"`
	OrgId    string `xml:"orgId,attr"`
	MsgId    string `xml:"msgId,attr"`
	ProdType string `xml:"prodType,attr"`
}

type TxnRegMob struct {
	ID     string `xml:"id,attr"`
	Note   string `xml:"note,attr"`
	RefId  string `xml:"refId,attr"`
	RefUrl string `xml:"refUrl,attr"`
	Ts     string `xml:"ts,attr"`
	Type   string `xml:"type,attr"`
}

type RespDataRegMob struct {
	ReqMsgId string `xml:"reqMsgId,attr"`
	Result   string `xml:"result,attr"`
}

type RespRegMob struct {
	XMLName xml.Name       `xml:"ns2:RespRegMob"`
	NS2     string         `xml:"xmlns:ns2,attr"`
	Head    HeadRegMob     `xml:"Head"`
	Txn     TxnRegMob      `xml:"Txn"`
	Resp    RespDataRegMob `xml:"Resp"`
}

func ReqRegMobfunc(c *gin.Context) error {
	resp := RespRegMob{
		XMLName: xml.Name{
			Local: "ns2:RespRegMob",
		},
		NS2: "http://npci.org/upi/schema/",
		Head: HeadRegMob{
			Ver:      "2.0",
			Ts:       "2022-11-02T01:35:39-04:00",
			OrgId:    "159022",
			MsgId:    "1vUWishpRrWlno0lbEX1",
			ProdType: "PSO",
		},
		Txn: TxnRegMob{
			ID:     "001QWeRx8y3IwbwsUYI0s8J5Y9wyRmYds8m",
			Note:   "FETCHLIST",
			RefId:  "001QWeRx8y3IwbwsUYI0s8J5Y9wyRmYds8m",
			RefUrl: "http://upi",
			Ts:     "2022-11-02T01:35:39-04:00",
			Type:   "PIN",
		},
		Resp: RespDataRegMob{
			ReqMsgId: "1vUWishpRrWlno0lbEX1",
			Result:   "SUCCESS",
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
