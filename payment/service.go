package payment

import (
	"bwa/golang/user"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"strconv"
)

var sn = snap.Client{}

type PaymentService interface {
	GetURLPayment(transaction TransactionPayment, user user.User) (string, error)
}

type service struct {
}

func NewPaymentService() PaymentService {
	return &service{}
}

func (s *service) GetURLPayment(transaction TransactionPayment, user user.User) (string, error) {
	midtrans.ServerKey = ""
	midtrans.Environment = midtrans.Sandbox

	sn.New("", midtrans.Sandbox)

	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.Id),
			GrossAmt: int64(transaction.Amount),
		},
		//CreditCard: &snap.CreditCardDetails{
		//	Secure: true,
		//},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Name,
			Email: user.Email,
		},
	}

	snapToken, _ := sn.CreateTransactionUrl(snapReq)

	return snapToken, nil
}

// func (s *service) GetToken(transaction transaction.Transaction,req *snap.Request) (string,*midtrans.Error) {
// 	// s.Options.SetPaymentOverrideNotification("https://example.com/url2")

// 	// resp, err := s.CreateTransactionToken(GenerateSnapReq())
// 	// if err != nil {
// 	// 	fmt.Println("Error :", err.GetMessage())
// 	// }
// 	// fmt.Println("Response : ", resp)

// }

//func GenerateSnapReq(transaction TransactionPayment, user user.User) *snap.Request {
//
//	// Initiate Snap Request
//	snapReq := &snap.Request{
//		TransactionDetails: midtrans.TransactionDetails{
//			OrderID:  strconv.Itoa(transaction.Id),
//			GrossAmt: int64(transaction.Amount),
//		},
//		CreditCard: &snap.CreditCardDetails{
//			Secure: true,
//		},
//		CustomerDetail: &midtrans.CustomerDetails{
//			FName: user.Name,
//			Email: user.Email,
//		},
//	}
//
//	return snapReq
//}

//func setupGlobalMidtransConfig() {
//	midtrans.ServerKey = "example.SandboxServerKey1"
//	midtrans.Environment = midtrans.Sandbox
//
//	// // Optional : here is how if you want to set append payment notification globally
//	// midtrans.SetPaymentAppendNotification("https://example.com/append")
//	// // Optional : here is how if you want to set override payment notification globally
//	// midtrans.SetPaymentOverrideNotification("https://example.com/override")
//
//	//// remove the comment bellow, in cases you need to change the default for Log Level
//	// midtrans.DefaultLoggerLevel = &midtrans.LoggerImplementation{
//	//	 LogLevel: midtrans.LogInfo,
//	// }
//}

//func initializeSnapClient() {
//	s.New("example.SandboxServerKey1", midtrans.Sandbox)
//}
