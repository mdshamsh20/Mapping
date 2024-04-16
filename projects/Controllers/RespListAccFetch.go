package Controllers

import (
	"encoding/xml"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Link struct {
	Type  string `xml:"type,attr"`
	Value string `xml:"value,attr"`
}

type Device struct {
	Tags []Tag `xml:"Tag"`
}

type Ac struct {
	AddrType string `xml:"addrType,attr"`
	Detail   Detail `xml:"Detail"`
}

type Payer struct {
	Addr   string `xml:"addr,attr"`
	Name   string `xml:"name,attr"`
	SeqNum string `xml:"seqNum,attr"`
	Type   string `xml:"type,attr"`
	Code   string `xml:"code,attr"`
	Device Device `xml:"Device"`
	Ac     Ac     `xml:"Ac"`
}

type ReqListAccountFetch struct {
	XMLName xml.Name `xml:"http://npci.org/token/schema/ ReqListAccount"`
	Head    Head     `xml:"Head"`
	Txn     Txn      `xml:"Txn"`
	Link    Link     `xml:"Link"`
	Payer   Payer    `xml:"Payer"`
}

func RespListAccFetchApi(c *gin.Context) {
	reqBody := RequestBodyLogger(c)
	log.Info("@RespListAccFetchApi Request body:", reqBody)
	var req ReqListAccountFetch
	err := xml.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		log.Error("@RespListAccFetchApi Error decoding request body:", err)
		InternalServerErrorResponse(c, err.Error())
		return
	}
	errInCrResp := RespListAccFetchFunc(c)
	if errInCrResp != nil {
		log.Error("@RespListAccFetchApi Error: ", err)
		InternalServerErrorResponse(c, errInCrResp.Error())
		return
	}
	log.Info("@RespListAccFetchApi Successfully created request.")
}

type HeadFetch struct {
	MsgId    string `xml:"msgId,attr"`
	OrgId    string `xml:"orgId,attr"`
	ProdType string `xml:"prodType,attr"`
	Ts       string `xml:"ts,attr"`
	Ver      string `xml:"ver,attr"`
}

type RespFetch struct {
	Ac       string `xml:"ac,attr"`
	Lk       string `xml:"lk,attr"`
	ReqMsgId string `xml:"reqMsgId,attr"`
	Result   string `xml:"result,attr"`
	Sa       string `xml:"sa,attr"`
}

type TxnFetch struct {
	Id     string `xml:"id,attr"`
	Note   string `xml:"note,attr"`
	RefId  string `xml:"refId,attr"`
	RefUrl string `xml:"refUrl,attr"`
	Ts     string `xml:"ts,attr"`
	Type   string `xml:"type,attr"`
}

type Account struct {
	AccType              string         `xml:"accType,attr"`
	IsLinkedDefault      string         `xml:"isLinkedDefault,attr"`
	KycStatus            string         `xml:"kycStatus,attr"`
	LinkedAccIfsc        string         `xml:"linkedAccIfsc,attr"`
	LinkedAccMaskedNum   string         `xml:"linkedAccMaskedNum,attr"`
	LinkedAccRefNum      string         `xml:"linkedAccRefNum,attr"`
	LinkedAccType        string         `xml:"linkedAccType,attr"`
	Name                 string         `xml:"name,attr"`
	TspId                string         `xml:"tspId,attr"`
	WalletAddress        string         `xml:"walletAddress,attr"`
	IsPCBDCWalletEnabled string         `xml:"isP-CBDCWalletEnabled,attr"`
	CredsAllowed         []CredsAllowed `xml:"CredsAllowed"`
}

type CredsAllowed struct {
	DLength string `xml:"dLength,attr"`
	DType   string `xml:"dType,attr"`
	SubType string `xml:"subType,attr"`
	Type    string `xml:"type,attr"`
}

type AccountList struct {
	Account Account `xml:"Account"`
}

type Signature struct {
	SignedInfo     SignedInfo `xml:"SignedInfo"`
	Xmlns          string     `xml:"xmlns,attr"`
	SignatureValue string     `xml:"SignatureValue"`
}

type SignedInfo struct {
	CanonicalizationMethod CanonicalizationMethod `xml:"CanonicalizationMethod"`
	SignatureMethod        SignatureMethod        `xml:"SignatureMethod"`
	Reference              Reference              `xml:"Reference"`
}

type CanonicalizationMethod struct {
	Algorithm string `xml:"Algorithm,attr"`
}

type SignatureMethod struct {
	Algorithm string `xml:"Algorithm,attr"`
}

type Reference struct {
	URI          string       `xml:"URI,attr"`
	Transforms   Transforms   `xml:"Transforms"`
	DigestMethod DigestMethod `xml:"DigestMethod"`
	DigestValue  string       `xml:"DigestValue"`
}

type Transforms struct {
	Transform Transform `xml:"Transform"`
}

type Transform struct {
	Algorithm string `xml:"Algorithm,attr"`
}

type DigestMethod struct {
	Algorithm string `xml:"Algorithm,attr"`
}

func RespListAccFetchFunc(c *gin.Context) error {

	head := HeadFetch{
		MsgId:    "1vUWishtkFdgpufdLd3U",
		OrgId:    "NPCI",
		ProdType: "PSO",
		Ts:       time.Now().Format("2006-01-02T15:04:05-07:00"),
		Ver:      "2.0",
	}

	resp := RespFetch{
		Ac:       "",
		Lk:       "",
		ReqMsgId: "001kYnce9j4KCkfsXGKHbzBFMlkw6Xf3rT8",
		Result:   "SUCCESS",
		Sa:       "",
	}

	txn := TxnFetch{
		Id:     "001kYnce9j4KCkfsXGKHbzBFMlkw6Xf3rT8",
		Note:   "FETCH",
		RefId:  "0",
		RefUrl: "http://upi",
		Ts:     time.Now().Format("2006-01-02T15:04:05-07:00"),
		Type:   "FETCH",
	}

	account := Account{
		AccType:              "WALLET",
		IsLinkedDefault:      "true",
		KycStatus:            "NO|MIN|FULL",
		LinkedAccIfsc:        "BNKD0000001",
		LinkedAccMaskedNum:   "***********535767",
		LinkedAccRefNum:      "535767",
		LinkedAccType:        "SAVINGS",
		Name:                 "TEST USER_535767",
		TspId:                "159022",
		WalletAddress:        "9a0dbbf53ae84f42852da66296a5c7d2",
		IsPCBDCWalletEnabled: "TRUE|FALSE",
		CredsAllowed: []CredsAllowed{
			{
				DLength: "6",
				DType:   "NUM",
				SubType: "TPIN",
				Type:    "PIN",
			},
			{
				DLength: "6",
				DType:   "NUM",
				SubType: "SMS",
				Type:    "OTP",
			},
		}}

	accountList := AccountList{
		Account: account,
	}

	signature := Signature{
		Xmlns: "http://www.w3.org/2000/09/xmldsig#",
		SignedInfo: SignedInfo{
			CanonicalizationMethod: CanonicalizationMethod{
				Algorithm: "http://www.w3.org/TR/2001/REC-xml-c14n-20010315",
			},
			SignatureMethod: SignatureMethod{
				Algorithm: "http://www.w3.org/2001/04/xmldsig-more#rsa-sha256",
			},
			Reference: Reference{
				URI: "",
				Transforms: Transforms{
					Transform: Transform{
						Algorithm: "http://www.w3.org/2000/09/xmldsig#enveloped-signature",
					},
				},
				DigestMethod: DigestMethod{
					Algorithm: "http://www.w3.org/2001/04/xmlenc#sha256",
				},
				DigestValue: "ko0qt2fqQP+xV/xLaaXy+KWA97F2gXRBrygZxBY2ooo=",
			},
		},
		SignatureValue: "UahNNyPs1OTrSbBvPZB3S3A6W54R/DytWGcPIYM2Iq9mxHoS4bD41PRrIx04nABN7YTdU2WqlaL6+sIjMCN2QYVK/TrOLOzx9jR2zpuwEpNcMtYT1pqQh0gUT42HEews iILxfOw95/0C0mVua/qk3AsK7OU8excgtufuJpSL03qQSQbGE7 VKfQ1omIhHUFEq+wcKKK1i00NY kOO5VybUk0Q132gHj1ZgVQ66In2T4Msvg4eyhYCFdYn4nTRA5cZ9UgSK+wiCv53vafLbd4+bEjKW UY1cl6xE+1609BnDQNAW6QBkeo+jrxrWSN3Fyxextz/hS8QXdS/3oBswNwSfeA==",
	}

	respListAccount := struct {
		XMLName     xml.Name    `xml:"ns2:RespListAccount"`
		XmlnsNs2    string      `xml:"xmlns:ns2,attr"`
		XmlnsNs3    string      `xml:"xmlns:ns3,attr"`
		Head        HeadFetch   `xml:"Head"`
		Resp        RespFetch   `xml:"Resp"`
		Txn         TxnFetch    `xml:"Txn"`
		AccountList AccountList `xml:"AccountList"`
		Signature   Signature   `xml:"Signature"`
	}{
		XmlnsNs2:    "http://npci.org/token/schema/",
		XmlnsNs3:    "http://npci.org/cm/schema/",
		Head:        head,
		Resp:        resp,
		Txn:         txn,
		AccountList: accountList,
		Signature:   signature,
	}

	xmlData, err := xml.MarshalIndent(respListAccount, "", "    ")
	if err != nil {
		log.Error("Failed to generate XML response: ", err)
		return err
	}

	SuccessResponseXml(c, xmlData)
	return nil
}
