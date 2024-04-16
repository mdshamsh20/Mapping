package Controllers

import (
	"encoding/json"
	"fmt"
	models "my-project/Models"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type AccountClosureRequestV2 struct {
	ClientId string `json:"client_id" validate:"required"`
	// SupportType          string json:"support_type" validate:"required,oneof=feedback other 'trading & marketing' sgb reports orders login kyc fund brokerage 'account related' 'account opening' 'account closure'"
	// InvestorName         string json:"investor_name" validate:"required_if=SupportType feedback,required_if=SupportType other,required_if=SupportType 'trading & marketing',required_if=SupportType sgb,required_if=SupportType reports,required_if=SupportType orders,required_if=SupportType login,required_if=SupportType kyc,required_if=SupportType fund,required_if=SupportType brokerage,required_if=SupportType 'account related',required_if=SupportType 'account opening'"
	// EmailId              string json:"email_id" validate:"required_if=SupportType feedback,required_if=SupportType other,required_if=SupportType 'trading & marketing',required_if=SupportType sgb,required_if=SupportType reports,required_if=SupportType orders,required_if=SupportType login,required_if=SupportType kyc,required_if=SupportType fund,required_if=SupportType brokerage,required_if=SupportType 'account related',required_if=SupportType 'account opening'"
	// Message              string json:"message" validate:"required_if=SupportType feedback,required_if=SupportType other,required_if=SupportType 'trading & marketing',required_if=SupportType sgb,required_if=SupportType reports,required_if=SupportType orders,required_if=SupportType login,required_if=SupportType kyc,required_if=SupportType fund,required_if=SupportType brokerage,required_if=SupportType 'account related',required_if=SupportType 'account opening'"
	// AccClosureType       string json:"acc_closure_type" validate:"required_if=SupportType 'account closure'"
	// AccClosureReason     string json:"acc_closure_reason" validate:"required_if=SupportType 'account closure'"
	// AccClosureReasonSpec string json:"acc_closure_reason_spec" validate:"required_if=AccClosureReason others"
	// CreatedBy            string json:"createdBy"
}

func AccClosureRequestV2(c *gin.Context) {

	var accClosureReqV2 AccountClosureRequestV2
	json.NewDecoder(c.Request.Body).Decode(&accClosureReqV2)

	fmt.Println("pppppppppppppppppppppppppppppppppppp : ", accClosureReqV2)

}

// errStore := StoreData(req)
//
//	if errStore != nil {
//		log.Error("Error storing data to DB, error: ", errStore)
//		InternalServerErrorResponse(c, errStore.Error())
//		return
//	}
// func StoreData(data ReqStruct) error {
// 	db, err := GormDB()
// 	if err != nil {
// 		log.Error("Eror in DB connection Error: ", err)
// 		return err
// 	}
// 	defer db.Close()
// 	fmt.Println("data: ", data)
// 	err = db.Debug().Create(models.User{Id: data.Id, Name: data.Name, Email: data.Email, Mobile: data.Mobile, Address: data.Address}).Error
// 	if err != nil {
// 		log.Error("Error inserting data in DB, error: ", err)
// 		return err
// 	}
// 	return nil
// }

func FetchFromDB(c *gin.Context) {
	db, err := GormDB()
	if err != nil {
		log.Error("Eror in DB connection Error: ", err)
		InternalServerErrorResponse(c, err.Error())
		return
	}
	defer db.Close()
	var users []models.User
	err = db.Debug().Select("*").Find(&users).Error
	if err != nil {
		log.Error("Error fetching data in DB, error: ", err)
		InternalServerErrorResponse(c, err.Error())
		return
	}
	fmt.Println("users: ", users)
	jsonData, err := json.Marshal(users)
	if err != nil {
		log.Println("Error marshaling JSON: ", err)
		return
	}
	fmt.Println("users2 : ", users, jsonData)

	// SuccessResponse(c, "success", map[string]interface{}{"data: ": jsonData})
}
