package Controllers

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ReqListKeysGetWallet struct {
	XMLName xml.Name `xml:"http://npci.org/token/schema/ ReqListKeys"`
	Head    Head     `xml:"Head"`
	Txn     Txn      `xml:"Txn"`
	Ac      Ac       `xml:"Ac"`
	Creds   Creds    `xml:"Creds"`
}

type HeadGetWallet struct {
	Ver      string `xml:"ver,attr"`
	Ts       string `xml:"ts,attr"`
	OrgID    string `xml:"orgId,attr"`
	MsgID    string `xml:"msgId,attr"`
	ProdType string `xml:"prodType,attr"`
}

type TxnGetWallet struct {
	ID       string `xml:"id,attr"`
	Note     string `xml:"note,attr"`
	RefID    string `xml:"refId,attr"`
	RefUrl   string `xml:"refUrl,attr"`
	Ts       string `xml:"ts,attr"`
	Type     string `xml:"type,attr"`
	PspOrgID string `xml:"pspOrgId,attr"`
	Mobile   string `xml:"mobile,attr"`
}

type AcGetWallet struct {
	AddrType string   `xml:"addrType,attr"`
	Detail   []Detail `xml:"Detail"`
}

type DetailGetWallet struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type CredsGetWallet struct {
	Cred Cred `xml:"Cred"`
}

type CredGetWallet struct {
	Type    string `xml:"type,attr"`
	SubType string `xml:"subType,attr"`
	Data    Data   `xml:"Data"`
}

type DataGetWallet struct {
	Code string `xml:"code,attr"`
	Ki   string `xml:"ki,attr"`
}

func RespListKeysGetWalletAPI(c *gin.Context) {
	reqBody := RequestBodyLogger(c)
	log.Info("@RespListKeysGetWallet Request body:", reqBody)
	var req ReqListKeysGetWallet
	err := xml.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		log.Error("@RespListKeysGetWallet Error decoding request body:", err)
		InternalServerErrorResponse(c, err.Error())
		return
	}
	errInCrResp := RespListKeysGetWalletfunc(c)
	if errInCrResp != nil {
		log.Error("@RespListKeysGetWallet Error: ", err)
		InternalServerErrorResponse(c, errInCrResp.Error())
		return
	}
	log.Info("@RespListKeysGetWallet Successfully created request.")
}

type RespListKeysGetWallet struct {
	XMLName            xml.Name            `xml:"ns2:RespListKeys"`
	NS2                string              `xml:"xmlns:ns2,attr"`
	NS3                string              `xml:"xmlns:ns3,attr"`
	Head               HeadRespGetWallet   `xml:"Head"`
	Resp               RespGetWallet       `xml:"Resp"`
	Txn                TxnRespGetWallet    `xml:"Txn"`
	Ac                 AcGetWallet         `xml:"Ac"`
	ParamsList         ParamsListGetWallet `xml:"paramsList"`
	ParamListSignature ParamListSignature  `xml:"paramListSignature"`
	Signature          SignatureGetWallet  `xml:"Signature"`
}

type HeadRespGetWallet struct {
	MsgId string `xml:"msgId,attr"`
	OrgId string `xml:"orgId,attr"`
	Ts    string `xml:"ts,attr"`
	Ver   string `xml:"ver,attr"`
}

type RespGetWallet struct {
	ReqMsgId string `xml:"reqMsgId,attr"`
	Result   string `xml:"result,attr"`
}

type TxnRespGetWallet struct {
	Id       string `xml:"id,attr"`
	Note     string `xml:"note,attr"`
	PspOrgId string `xml:"pspOrgId,attr"`
	RefId    string `xml:"refId,attr"`
	RefUrl   string `xml:"refUrl,attr"`
	Ts       string `xml:"ts,attr"`
	Type     string `xml:"type,attr"`
}

// type AcGetWallet struct {
// 	AddrType string   `xml:"addrType,attr"`
// 	Details  []Detail `xml:"Detail"`
// }

// type Detail struct {
// 	Name  string `xml:"name,attr"`
// 	Value string `xml:"value,attr"`
// }

type ParamsListGetWallet struct {
	Params []ParamGetWallet `xml:"param"`
}

type ParamGetWallet struct {
	Code       string `xml:"code,attr"`
	Owner      string `xml:"owner,attr"`
	ParamValue ParamValue
}
type ParamValue struct {
	XMLName   xml.Name `xml:"paramValue"`
	XMLNS_XS  string   `xml:"xmlns:xs,attr"`
	XMLNS_XSI string   `xml:"xmlns:xsi,attr"`
	XSI_Type  string   `xml:"xsi:type,attr"`
	Value     string   `xml:",chardata"`
}
type ParamListSignature struct {
	Code           string         `xml:"code,attr"`
	Owner          string         `xml:"owner,attr"`
	Type           string         `xml:"type,attr"`
	SignatureValue SignatureValue `xml:"signatureValue"`
}

type SignatureGetWallet struct {
	XMLName        xml.Name            `xml:"Signature"`
	SignedInfo     SignedInfoGetWallet `xml:"SignedInfo"`
	SignatureValue string              `xml:"signatureValue"`
}
type SignatureValue struct {
	XMLName   xml.Name `xml:"signatureValue"`
	XMLNS_XS  string   `xml:"xmlns:xs,attr"`
	XMLNS_XSI string   `xml:"xmlns:xsi,attr"`
	XSI_Type  string   `xml:"xsi:type,attr"`
	Value     string   `xml:",chardata"`
}

type SignedInfoGetWallet struct {
	CanonicalizationMethod CanonicalizationMethodGetWallet `xml:"CanonicalizationMethod"`
	SignatureMethod        SignatureMethodGetWallet        `xml:"SignatureMethod"`
	Reference              ReferenceGetWallet              `xml:"Reference"`
}

type CanonicalizationMethodGetWallet struct {
	Algorithm string `xml:"Algorithm,attr"`
}

type SignatureMethodGetWallet struct {
	Algorithm string `xml:"Algorithm,attr"`
}

type ReferenceGetWallet struct {
	URI          string                `xml:"URI,attr"`
	Transforms   []TransformGetWallet  `xml:"Transforms>Transform"`
	DigestMethod DigestMethodGetWallet `xml:"DigestMethod"`
	DigestValue  string                `xml:"DigestValue"`
}

type TransformGetWallet struct {
	Algorithm string `xml:"Algorithm,attr"`
}

type DigestMethodGetWallet struct {
	Algorithm string `xml:"Algorithm,attr"`
}

func RespListKeysGetWalletfunc(c *gin.Context) error {
	resp := RespListKeysGetWallet{
		NS2: "http://npci.org/upi/schema/",
		NS3: "http://npci.org/cm/schema/",
		Head: HeadRespGetWallet{
			MsgId: "1vUWishpRqKuwFUpbnR7",
			OrgId: "NPCI",
			Ts:    time.Now().Format(time.RFC3339),
			Ver:   "2.0",
		},
		Resp: RespGetWallet{
			ReqMsgId: "HBiIbtPlMZhjHeBhs5ZOzB7OkmqGmzQy",
			Result:   "SUCCESS",
		},
		Txn: TxnRespGetWallet{
			Id:       "HBiIbtPlMZhjHeBhs5ZOzB7OkmqGmzQy",
			Note:     "GetWalletKey",
			PspOrgId: "159022",
			RefId:    "HBiIbtPlMZhjHeBhs5ZOzB7OkmqGmzQy",
			RefUrl:   "http://upi",
			Ts:       "2022-12-01T18:53:38.964",
			Type:     "GetWalletKey",
		},
		Ac: AcGetWallet{
			AddrType: "WALLET",
			Detail: []Detail{
				{Name: "WALLETADDRESS", Value: "db554515f93846d59dcee483fde14262"},
				{Name: "TSPID", Value: "159022"},
				{Name: "TYPE", Value: "WALLET"},
			},
		},
		ParamsList: ParamsListGetWallet{
			Params: []ParamGetWallet{
				{Code: "DEVICE-CERT", Owner: "NPCI", ParamValue: ParamValue{
					XMLNS_XS:  "http://www.w3.org/2001/XMLSchema",
					XMLNS_XSI: "http://www.w3.org/2001/XMLSchema-instance",
					XSI_Type:  "xs:string",
					Value:     "VJnapEm/LArltAVm0KkeJJb1Pej74v6WxFsp1Mvxmg/ulLD60mR3C5O Tb0SxAvBzSDenu5pNC3LXlZ7IYnMyEZt1jOEXi6ST+2lEUmQyEe0/f5tyVOfcbaK+BoIghJRJ2Xr wxGeDJ/vPG4UReZX1lkP9QaVvFmbhibiXv6X+suj665xKSHc2BSVs9nPEF6yYxoSCEw2pIkH+E50 HcqTUFAeuGWDSEXaYTMwfnJnfDgSwMbrZ+OVHxmFJT+ar9jYJgjRVCtSNgI/jhVsBRBNNRueoHVF BZKC1cg3iqeoehWydGf4Nq/JdOmHRCCfUbjH0vJqctJ3uxtrjLRptI76QM/5k0IcQXGXI3ZC198u N5fW9YhPJwD29X13tY9GY5II/x7L4s55HVfFygxthDw+AFIFeLUQStQ4ZI04V3OeZXG2RhSnXQr5 R+2rXL/xfD3qQUO0Pu7Tkv1X8Yue5L/monpuIxcL5jJdlaGCWzv05nLw1shosRTGjITklsm5JZcx t2fVTF4EV+W4Tw7v50260qNYYQq4bNMAEFmuao07jXtWyvm3zlAIMf/TbBk1+AvgO1urZiSh3X6d /ft8iLradv+i2ntFwRQpyQMK5gyqFG4U1DEwaXs+Do6qmRRMd0CeoSi32gMMLmFNdKCcUzxmH0e6 IQl/EeeXI3RRwjFtqlEp3Mm8yMAZXFeTY0qpD5WpnLwBUIYTN3VuXgGddLpyoI6PvAfbbbvyu5wm 6fTN+/6jEBw5UORAkhOVhdThGpiSukSqyNKksQujfJfzwospc4hZmPaza6k+JsNbalk12ZyEaUzN pa/wusUaFcpibA5Qok7iCqhLnZsVZxig5FIMtuPBjR4+V3PfFSBXugTQRXsERdpQVn9fp44ZdZ7G AdjVIpBaD2FVXNIu5ySJkXD0N6ohh06kyY/qx6rdGuGyBdyltni9HpOFK/7yX/Ua4m4Y/RnOJdrd lOw4NRRskWV9LJWR/9yi7uVk1vqGqiLYqnvyZhBgGnQlfsXxcPQ5wIGIOhWOkqHdW/eIj9Lubgtm OWGlFIKtGlnzkXJmP/kycVMXMzhKTf0OkSGoLjKQQxdYstXYhuS9+YbARS69RqsKZ3Lrp3hFk0aO k7mHoZzpmBvNNyjvAanBfbqHPGq+HZ0PVzVyiqnDg90pqJRwuIWnDTA1e30hAcC4kpFqt9WS5x3+ fzHWdtAfM8F878SbQgzL7J3V+WXV1ENbR1kgi0S7TQ464NkV5l6vVy7CeUfEsx3Hmz+glDVqLpgk EwwS3fNtHImOpyWstesVTeh6JZA1XGvIB08LXS4cnyIVn7PCtw+JWRfu1JyMcgEvo6cgzYjkDgBl ct8cega11vvSZ2gFQiE8FwmiuL7Sqp454RL5YNWdBr9MPIFVZpFkaqJR15EXCW8IxY3sBxmuhaYx JZif0R6GWKkBJae2h1B6Hb5cXWesolEuPK8cKA3XnuWAJ5hclU+rxxjaTaYVCkwqtWrajA8xg3Wc xI+b9j6Rg63J/w5V/lA==|Mfa5NYTbjjVfA+tQSxoqiA==",
				}},
				{Code: "STATE", Owner: "NPCI", ParamValue: ParamValue{
					XMLNS_XS:  "http://www.w3.org/2001/XMLSchema",
					XMLNS_XSI: "http://www.w3.org/2001/XMLSchema-instance",
					XSI_Type:  "xs:string",
					Value:     "nmZRB2VQrBuSGwwpKjOvCoEqDRvYUXdDyy5pKxMraI890h1uQx8t4zf /s7yaldK6MpGy1vZuCDbPCAOvtLLX/N5d91zYrbuXfZhrG9qSwKiDwOkKyP5pjAtDS/Pp3D8EZAC Je3vwJ+sVF/CPr88J6cru0JyQjW13LUl5K58se2OI6VdnB11ohZKiDWi2r8yo7h67TJXJWCgBuFY YCbVc63zZWINx6Z2ylMOIPo07ZdnrZlpLOxh8pFzhXw2GUU3yWLuyYHUTOYMcPSKA6zqmGTYymNI 0gxZQEi4bwXibyfZh140CXmqqAO6/416iKjPQN8O2rePIpdvwKegmL7snVVsh79Sh+FrUxwgmoV4 xZ1oy/BtWbWObwPPHdT7gZ1oFVPuZ4OCzTSGNaqAoa4iGyM4AhaTB7tPcHY0x8cmAH7s+CbfuJkd KjUQ+gjo14PIq+VXnwvucErEiWRgBMx3txaBfJTg0mej3fDCT196YtBRD8LnnmW7x5PqRyBW8ppG Fmvl4qng5raZNqPYuyP1zB1rITTz3y6gztK/TtDN2VUW5RvAwkM3MOqJBV5ommXKGKDU8CMnRB42 1ASix6zco67BYMi2KBdgsiCKg9ASwkgMfr03AaLx2apMVX762CMeRSSPXC/6vGnr+MU97bnUYo6e 46U+Iye+KOxJ2qv3wL1WeSlurUhzjBgZVB/yoXcCzER3RGwfzHKb69mG4ObtL8Cz3kBjAnwTp/gx 66Ai+Fc0J50Tp9Wym4p/Xy+aeKjr1sDYscPkK40GvDtVZ7zxwNI50EFoxoeGDsyvujaYpWowQGbA 5m0Y+EebeTQ==|JrTMXAwQvsYJpVgKyjsy+A==",
				}},
			},
		},
		ParamListSignature: ParamListSignature{
			Code:  "PRIMARY",
			Owner: "NPCI",
			Type:  "SIGNATURE",
			SignatureValue: SignatureValue{
				XMLNS_XS:  "http://www.w3.org/2001/XMLSchema",
				XMLNS_XSI: "http://www.w3.org/2001/XMLSchema-instance",
				XSI_Type:  "xs:string",
				Value:     "YZnUZxVne/jQNUNhCg3dbCRN0wWtAJ1VRT5B0B9+7VEVZsdQ73fxVcO ixrO2qPRQACBF1OVJFmsuXn+qV/22VPttcwhZ47mipmqm7xX1eb2KoPruWfDUbs2K1N67LK/WZu1 J7saBXpf4Rsup9tLdZlv9StGoMP8gynw3smRnjRdDpxF1uMQbbQGlzNXQnE2KBj7fDPlJm1daH3A Y1kT1E0nQrLqBLMzbWGAYlzgUGjfKHN7W+hMKj4hHc74NQNavVKiKjF9Fj1IUXdwxGLApWLcpJ2+ OnWJ92P8NOrmx2Ji4lcobSw7jgM72+YsCyY0iQ5FpSQb4IebRzvQP2KYx9w==",
			},
		},
		Signature: SignatureGetWallet{
			SignedInfo: SignedInfoGetWallet{
				CanonicalizationMethod: CanonicalizationMethodGetWallet{Algorithm: "http://www.w3.org/TR/2001/REC-xml-c14n-20010315"},
				SignatureMethod:        SignatureMethodGetWallet{Algorithm: "http://www.w3.org/2001/04/xmldsig-more#rsa-sha256"},
				Reference: ReferenceGetWallet{
					URI:          "",
					Transforms:   []TransformGetWallet{{Algorithm: "http://www.w3.org/2000/09/xmldsig#enveloped-signature"}},
					DigestMethod: DigestMethodGetWallet{Algorithm: "http://www.w3.org/2001/04/xmlenc#sha256"},
					DigestValue:  "3PBjBADGnOBirHq6jLB+r6fslBV4W3ee8WXltR1ZA0s=",
				},
			},
			SignatureValue: "g0zmXWIc9PDFKf0sbnF5YBsFPrlotJquw3WRXHYZQmLA24EGPxWZ1s86fA+XL9xYap6smk8M92LGBfUVuRFwCi0C017t30WzAHC/I7yiZzqfq5bMiO2jsbS6XCoG2dcZX2A0WSD70V//EdigDwZYO75EWtuFUmLEGdTTd486ukWkNDUpuiKs7NIYXnc4MEjQ/vXfRtPtJRa/oA+9aYMRkmryo+G9oHAVb7CHomu0BXwgERH4JNZ3qFqag0pXVpmAs60z5OOXY4cfQBCbWZHxZnoT8trekqdGUZkx43Fcwgp1UMgvuhX1pEcLHKgqEMjVCzX978s3KE1xTbRE4eeoYA==",
		},
	}
	xmlResponse, err := xml.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Error(err.Error(), http.StatusInternalServerError)
		return err
	}
	SuccessResponseXml(c, xmlResponse)
	return nil

}
