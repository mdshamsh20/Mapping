package Controllers

import (
	"encoding/xml"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type HeadCr struct {
	Ver      string `xml:"ver,attr"`
	Ts       string `xml:"ts,attr"`
	OrgID    string `xml:"orgId,attr"`
	MsgID    string `xml:"msgId,attr"`
	ProdType string `xml:"prodType,attr"`
}

type TxnCr struct {
	ID     string `xml:"id,attr"`
	Note   string `xml:"note,attr"`
	RefID  string `xml:"refId,attr"`
	RefUrl string `xml:"refUrl,attr"`
	Ts     string `xml:"ts,attr"`
	Type   string `xml:"type,attr"`
	TspID  string `xml:"tspId,attr"`
}

type LinkCr struct {
	Type  string `xml:"type,attr"`
	Value string `xml:"value,attr"`
}

type TagCr struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type DeviceCr struct {
	Tags []Tag `xml:"Tag"`
}

type AcDetail struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type AcCr struct {
	AddrType string     `xml:"addrType,attr"`
	Details  []AcDetail `xml:"Detail"`
}

type PayerAccCr struct {
	Addr   string `xml:"addr,attr"`
	Name   string `xml:"name,attr"`
	SeqNum string `xml:"seqNum,attr"`
	Type   string `xml:"type,attr"`
	Code   string `xml:"code,attr"`
	Device Device `xml:"Device"`
	Ac     Ac     `xml:"Ac"`
}

type ReqListAccountCr struct {
	XMLName xml.Name `xml:"http://npci.org/token/schema/ ReqListAccount"`
	Head    Head     `xml:"Head"`
	Txn     Txn      `xml:"Txn"`
	Link    Link     `xml:"Link"`
	Payer   Payer    `xml:"Payer"`
}

func ReqListAccountCreateApi(c *gin.Context) {
	reqBody := RequestBodyLogger(c)
	log.Info("@ReqListAccountCreateApi Request body:", reqBody)
	var req ReqListAccountCr
	err := xml.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		log.Error("@ReqListAccountCreateApi Error decoding request body:", err)
		InternalServerErrorResponse(c, err.Error())
		return
	}
	errCreatResp := ReqListAccountCreatefunc(c)
	if errCreatResp != nil {
		log.Error("@ReqListAccountCreateApi Error: ", err)
		InternalServerErrorResponse(c, errCreatResp.Error())
		return
	}
	log.Info("@ReqListAccountCreateApi Successfully created request.")
}

type RespCr struct {
	ReqMsgId string `xml:"reqMsgId,attr"`
	Result   string `xml:"result,attr"`
	ErrCode  string `xml:"errCode,attr"`
}

type CredsAllowedCr struct {
	Type    string `xml:"type,attr"`
	SubType string `xml:"subType,attr"`
	DType   string `xml:"dType,attr"`
	DLength string `xml:"dLength,attr"`
}

type AccountCr struct {
	AccType       string         `xml:"accType,attr"`
	WalletAddress string         `xml:"walletAddress,attr"`
	TspId         string         `xml:"tspId,attr"`
	Name          string         `xml:"name,attr"`
	KycStatus     string         `xml:"kycStatus,attr"`
	CredsAllowed  []CredsAllowed `xml:"CredsAllowed"`
}

type AccountListCr struct {
	Accounts []Account `xml:"Account"`
}

type RespListAccount struct {
	XMLName     xml.Name      `xml:"token:RespListAccount"`
	XmlnsToken  string        `xml:"xmlns:token,attr"`
	Head        HeadCr        `xml:"Head"`
	Txn         TxnCr         `xml:"Txn"`
	Resp        RespCr        `xml:"Resp"`
	AccountList AccountListCr `xml:"AccountList"`
}

func ReqListAccountCreatefunc(c *gin.Context) error {
	resp := RespListAccount{
		XmlnsToken: "http://npci.org/token/schema/",
		Head: HeadCr{
			Ver:      "1.0|2.0",
			Ts:       "",
			OrgID:    "",
			MsgID:    "",
			ProdType: "PSO",
		},
		Txn: TxnCr{
			ID:     "",
			Note:   "",
			RefID:  "",
			RefUrl: "",
			Ts:     "",
			Type:   "CREATE",
			TspID:  "400011",
		},
		Resp: RespCr{
			ReqMsgId: "",
			Result:   "SUCCESS",
			ErrCode:  "",
		},
		AccountList: AccountListCr{
			Accounts: []Account{
				{
					AccType:       "WALLET",
					WalletAddress: "sad32n3d2nd2f2qfn0i23",
					TspId:         "400011",
					Name:          "Harsh Malik",
					KycStatus:     "NO",
					CredsAllowed: []CredsAllowed{
						{Type: "PIN", SubType: "TPIN", DType: "", DLength: ""},
						{Type: "OTP", SubType: "SMS", DType: "", DLength: ""},
					},
				},
			},
		},
	}

	xmlData, err := xml.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Error(err.Error(), http.StatusInternalServerError)
		return err
	}
	SuccessResponseXml(c, xmlData)
	return nil
}
