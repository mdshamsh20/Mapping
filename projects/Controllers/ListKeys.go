package Controllers

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Txn struct {
	Id       string `xml:"id,attr"`
	Note     string `xml:"note,attr"`
	PspOrgId string `xml:"pspOrgId,attr"`
	RefId    string `xml:"refId,attr"`
	RefUrl   string `xml:"refUrl,attr"`
	Ts       string `xml:"ts,attr"`
	Type     string `xml:"type,attr"`
}

type ReqListKeys struct {
	XMLName xml.Name `xml:"http://npci.org/token/schema/ ReqListKeys"`
	XMLNS   string   `xml:"xmlns:token,attr"`
	Head    Head     `xml:"Head"`
	Txn     Txn      `xml:"Txn"`
}

func ListKeysApi(c *gin.Context) {
	reqBody := RequestBodyLogger(c)
	log.Info("Request body:", reqBody)
	var req ReqListKeys
	err := xml.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		log.Error("Error decoding request body:", err)
		InternalServerErrorResponse(c, err.Error())
		return
	}
	ListKeyFunc(c)

	log.Info("Successfully created request.")
}

type RespKeyList struct {
	ReqMsgId string `xml:"reqMsgId,attr"`
	Result   string `xml:"result,attr"`
}

type Key struct {
	Code     string `xml:"code,attr"`
	Owner    string `xml:"owner,attr"`
	Type     string `xml:"type,attr"`
	Ki       string `xml:"ki,attr"`
	KeyValue KeyValue
}
type KeyValue struct {
	XMLName xml.Name `xml:"keyValue"`
	XSI     string   `xml:"xmlns:xsi,attr"`
	XS      string   `xml:"xmlns:xs,attr"`
	Type    string   `xml:"xsi:type,attr"`
	Value   string   `xml:",chardata"`
}

type KeyList struct {
	Keys []Key `xml:"key"`
}

type RespListKeys struct {
	XMLName xml.Name `xml:"ns2:RespListKeys"`
	NS2     string   `xml:"xmlns:ns2,attr"`
	Head    Head
	Resp    RespKeyList
	Txn     Txn
	KeyList KeyList
}

func ListKeyFunc(c *gin.Context) {
	resp := RespListKeys{
		NS2: "http://npci.org/token/schema/",
		Head: Head{
			Ver:   "2.0",
			Ts:    time.Now().Format(time.RFC3339),
			OrgID: "NPCI",
			MsgID: "1vUWishpRrWikuqfrFiO",
		},
		Resp: RespKeyList{
			ReqMsgId: "001cpHRvI5XH8hc2ezt7PI2xebpuHpN8eO0",
			Result:   "SUCCESS",
		},
		Txn: Txn{
			Id:       "001cpHRvI5XH8hc2ezt7PI2xebpuHpN8eO0",
			Note:     "NA",
			RefId:    "001cpHRvI5XH8hc2ezt7PI2xebpuHpN8eO0",
			RefUrl:   "http://upi",
			Ts:       time.Now().Format(time.RFC3339),
			Type:     "ListKeys",
			PspOrgId: "159022",
		},
		KeyList: KeyList{
			Keys: []Key{
				{
					Code:  "NPCI",
					Owner: "NPCI",
					Type:  "PKI",
					Ki:    "20150822",
					KeyValue: KeyValue{
						XSI:   "http://www.w3.org/2001/XMLSchema-instance",
						XS:    "http://www.w3.org/2001/XMLSchema",
						Type:  "xs:string",
						Value: "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuMKxWfy0WcP p98muBWa6yhpmb6ZGZGSKHRIOv05UlIN5TbUPl6yEerh7Wj0+JyKfsOntRdAVhkLJGRoHwH6gEEe FNHge7kPea/B33cQAbqa39mnP5F1aaZT3tjJnKrfI1Wum0crdb7dAMzft4JILOEa+s3Uh7OdYEl/ Xp7EisdSoJ345Cj0LTfLZEQzRdVGovXZrfLByJysH11V9tDrIVv75C/3UndwjHt3NrqzNBoUMh5V ZRFkcwuebUAkhIed5gvoysJwd0yYGrAUXNrXJJDTAj5diCuasWyfWZR9lsX5l14hdxF+lqadR/pg II53DW5oEy2LMXgvt2u/qmSml8wIDAQAB",
					},
				},
			},
		},
	}

	// Encode the XML response
	encodedResponse, err := xml.MarshalIndent(resp, "", "    ")
	if err != nil {
		fmt.Printf("Error encoding XML: %v", err)
		return
	}
	SuccessResponseXml(c, encodedResponse)
}
