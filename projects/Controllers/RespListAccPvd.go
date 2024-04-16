package Controllers

import (
	"encoding/xml"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ReqListAccPvd struct {
	XMLName xml.Name `xml:"http://npci.org/upi/schema/ ReqListAccPvd"`
	Head    Head     `xml:"Head"`
	Txn     Txn      `xml:"Txn"`
}

func RespListAccPvdApi(c *gin.Context) {
	reqBody := RequestBodyLogger(c)
	log.Info("@RespListAccPvdApi Request body:", reqBody)
	var req ReqListAccPvd
	err := xml.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		log.Error("@RespListAccPvdApi Error decoding request body:", err)
		InternalServerErrorResponse(c, err.Error())
		return
	}
	errInCrResp := RespListAccPvdFunc(c)
	if errInCrResp != nil {
		log.Error("@RespListAccPvdFunc Error: ", err)
		InternalServerErrorResponse(c, errInCrResp.Error())
		return
	}
	log.Info("@RespListAccPvdApi Successfully created request.")
}

type RespListAccPvd struct {
	XMLName    xml.Name   `xml:"ns2:RespListAccPvd"`
	XmlnsNs2   string     `xml:"xmlns:ns2,attr"`
	Head       Head       `xml:"Head"`
	Resp       RespPvd    `xml:"Resp"`
	Txn        TxnPvd     `xml:"Txn"`
	AccPvdList AccPvdList `xml:"AccPvdList"`
}

type RespPvd struct {
	ReqMsgID string `xml:"reqMsgId,attr"`
	Result   string `xml:"result,attr"`
	// ErrCode  string `xml:"errCode,attr"`
	// Msg      string `xml:"msg,attr"`
}

type TxnPvd struct {
	Id       string `xml:"id,attr"`
	Note     string `xml:"note,attr"`
	PspOrgId string `xml:"pspOrgId,attr"`
	RefId    string `xml:"refId,attr"`
	RefUrl   string `xml:"refUrl,attr"`
	Ts       string `xml:"ts,attr"`
}

type AccPvdList struct {
	AccPvds []AccPvd `xml:"AccPvd"`
}

type AccPvd struct {
	Name             string           `xml:"name,attr"`
	Iin              string           `xml:"iin,attr"`
	Ifsc             string           `xml:"ifsc,attr"`
	Active           string           `xml:"active,attr"`
	Url              string           `xml:"url,attr"`
	SpocName         string           `xml:"spocName,attr"`
	SpocEmail        string           `xml:"spocEmail,attr"`
	SpocPhone        string           `xml:"spocPhone,attr"`
	TspId            string           `xml:"tspId,attr"`
	OrgWalletAddress string           `xml:"orgWalletAddress,attr"`
	Prods            string           `xml:"prods,attr"`
	LastModifedTs    string           `xml:"lastModifedTs,attr"`
	MobRegFormat     string           `xml:"mobRegFormat,attr"`
	VersionSupported VersionSupported `xml:"VersionSupported"`
}

type VersionSupported struct {
	Versions []Version `xml:"Version"`
}

type Version struct {
	No          string `xml:"no,attr"`
	Description string `xml:"description,attr"`
	Mandatory   string `xml:"mandatory,attr"`
}

func RespListAccPvdFunc(c *gin.Context) error {
	xmlData := RespListAccPvd{
		XmlnsNs2: "http://npci.org/token/schema/",
		Head: Head{
			Ver:   "2.0",
			Ts:    "2022-07-28T21:41:52+05:30",
			OrgID: "NPCI",
			MsgID: "1vU7XpDnWNfX4r12JzBh",
		},
		Resp: RespPvd{
			ReqMsgID: "b2ZbOEWcnkDOzNDW6AXJNOPv1uO6OZxGijb",
			Result:   "SUCCESS",
		},
		Txn: TxnPvd{
			Id:     "JZWRxIizybKtYB3M714bSnsrDtTh9KJhnJE",
			Note:   "LISTACCPVD",
			RefId:  "b2ZbOEWcnkDOzNDW6AXJNOPv1uO6OZxGijb",
			RefUrl: "http://upi",
			Ts:     "2022-07-28T21:41:52+05:30",
		},
		AccPvdList: AccPvdList{
			AccPvds: []AccPvd{
				{
					Name:             "BANK B",
					Iin:              "600010",
					Ifsc:             "BNKB",
					Active:           "Y",
					Url:              "http://10.40.136.35:9002",
					SpocName:         "BANK spoc",
					SpocEmail:        "spoc@npci.com",
					SpocPhone:        "9999999999",
					TspId:            "159015",
					OrgWalletAddress: "0x000000000000000000001",
					Prods:            "UPI,PSO",
					LastModifedTs:    "2022-06-14T14:27:20+05:30",
					MobRegFormat:     "FORMAT2",
					VersionSupported: VersionSupported{
						Versions: []Version{
							{No: "2.0", Description: "ITS A FIR VERSION", Mandatory: "true"},
							{No: "2.95", Description: "CENTRAL_MAPPER_VERSION", Mandatory: "false"},
						},
					},
				},
				{
					Name:             "BANKA",
					Iin:              "600008",
					Ifsc:             "BNKA",
					Active:           "Y",
					Url:              "http://10.40.136.46:9006",
					SpocName:         "BANKA spoc",
					SpocEmail:        "spoc@npci.com",
					SpocPhone:        "9999999999",
					OrgWalletAddress: "0x000000000000000000003",
					Prods:            "UPI,PSO",
					LastModifedTs:    "2022-06-01T12:40:20+05:30",
					MobRegFormat:     "FORMAT2",
					VersionSupported: VersionSupported{
						Versions: []Version{
							{No: "2.0", Description: "ITS A FIR VERSION", Mandatory: "true"},
							{No: "2.95", Description: "CENTRAL_MAPPER_VERSION", Mandatory: "false"},
						},
					},
				},
				{
					Name:             "ICICI",
					Iin:              "100228",
					Ifsc:             "ICIC",
					Active:           "Y",
					Url:              "http://10.40.136.46:9006",
					SpocName:         "ICIC SPOC",
					SpocEmail:        "spoc@npci.com",
					SpocPhone:        "9999999999",
					TspId:            "400011",
					OrgWalletAddress: "0x000000000000000000004",
					Prods:            "UPI,PSO",
					LastModifedTs:    "2022-06-14T02:12:45+05:30",
					MobRegFormat:     "FORMAT2",
					VersionSupported: VersionSupported{
						Versions: []Version{
							{No: "2.0", Description: "ITS A FIR VERSION", Mandatory: "true"},
							{No: "2.95", Description: "CENTRAL_MAPPER_VERSION", Mandatory: "false"},
						},
					},
				},
				{
					Name:             "SBI",
					Iin:              "100229",
					Ifsc:             "SBIB",
					Active:           "Y",
					Url:              "http://10.40.136.35:9006",
					SpocName:         "SBI SPOC",
					SpocEmail:        "spoc@npci.com",
					SpocPhone:        "9999999999",
					TspId:            "159002",
					OrgWalletAddress: "0x000000000000000000004",
					Prods:            "UPI,PSO",
					LastModifedTs:    "2022-06-14T02:10:43+05:30",
					MobRegFormat:     "FORMAT2",
					VersionSupported: VersionSupported{
						Versions: []Version{
							{No: "2.0", Description: "ITS A FIR VERSION", Mandatory: "true"},
							{No: "2.95", Description: "CENTRAL_MAPPER_VERSION", Mandatory: "false"},
						},
					},
				},
			},
		},
	}
	output, err := xml.MarshalIndent(xmlData, "", "    ")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}
	SuccessResponseXml(c, output)
	return nil
}
