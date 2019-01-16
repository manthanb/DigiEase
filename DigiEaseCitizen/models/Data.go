package models

type UserLoginRequest struct {
	UserEmail string `json:"userEmail"`
	Password  string `json:"password"`
}

type UserLoginResponse struct {
	UserID           int64  `json:"userID"`
	UserEmail        string `json:"userEmail"`
	UserName         string `json:"userName"`
	Password         string `json:"password"`
	Status           string `json:"status"`
	ExceptionMessage string `json:"exceptionMessage"`
}

type UserRegistrationRequest struct {
	UserName       string `json:"userName"`
	UserEmail      string `json:"userEmail"`
	Password       string `json:"password"`
	FirstName      string `json:"firstName"`
	MiddleName     string `json:"middleName"`
	LastName       string `json:"lastName"`
	Gender         string `json:"gender"`
	Date           string `json:"date"`
	Month          string `json:"month"`
	Year           string `json:"year"`
	Address1       string `json:"address1"`
	Address2       string `json:"address2"`
	Address3       string `json:"address3"`
	City           string `json:"city"`
	State          string `json:"state"`
	PinCode        string `json:"pinCode"`
	Mobile         string `json:"mobile"`
	Landline       string `json:"landline"`
	AadharCard     string `json:"aadharCard"`
	PanCard        string `json:"panCard"`
	Passport       string `json:"passport"`
	DrivingLicense string `json:"drivingLicense"`
	Photo          string `json:"photo"`
	AddressProof   string `json:"addressProof"`
}

type TenantFormRequest struct {
	UserEmail     string       `json:"userEmail"`
	Owner         OwnerData    `json:"owner"`
	Property      PropertyData `json:"property"`
	Tenant        TenantData   `json:"tenant"`
	Agent         AgentData    `json:"agent"`
	OwnerPhoto    string       `json:"ownerPhoto"`
	AgentPhoto    string       `json:"agentPhoto"`
	TenantPhoto   string       `json:"tenantPhoto"`
	OwnerPhotoID  string       `json:"ownerPhotoID"`
	TenantPhotoID string       `json:"tenantPhotoID"`
	ErrorMessage  string       `json:"errorMessage"`
}

// type TenantFormResponse struct {
// 	Owner        OwnerData    `json:"owner"`
// 	Property     PropertyData `json:"property"`
// 	Tenant       TenantData   `json:"tenant"`
// 	Agent        AgentData    `json:"agent"`
// 	ErrorMessage string       `json:"errorMessage"`
// }

type TenantFormResponse struct {
	Status           string `json:"status"`
	ExceptionMessage string `json:"exceptionMessage"`
}

type OwnerData struct {
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
	Age        string `json:"age"`
	Address1   string `json:"address1"`
	Address2   string `json:"address2"`
	Address3   string `json:"address3"`
	City       string `json:"city"`
	State      string `json:"state"`
	PinCode    string `json:"pinCode"`
	Mobile     string `json:"mobile"`
	Landline   string `json:"landline"`
}

type PropertyData struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	Address3 string `json:"address3"`
	City     string `json:"city"`
	State    string `json:"state"`
	PinCode  string `json:"pinCode"`
	Use      string `json:"use"`
}

type TenantData struct {
	FirstName        string `json:"firstName"`
	MiddleName       string `json:"middleName"`
	LastName         string `json:"lastName"`
	Age              string `json:"age"`
	Address1         string `json:"address1"`
	Address2         string `json:"address2"`
	Address3         string `json:"address3"`
	City             string `json:"city"`
	State            string `json:"state"`
	PinCode          string `json:"pinCode"`
	Mobile           string `json:"mobile"`
	Landline         string `json:"landline"`
	PermanentAddress string `json:"permanentAddress"`
	OfficeAddress    string `json:"officeAddress"`
}

type AgentData struct {
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
	Age        string `json:"age"`
	Address1   string `json:"address1"`
	Address2   string `json:"address2"`
	Address3   string `json:"address3"`
	City       string `json:"city"`
	State      string `json:"state"`
	PinCode    string `json:"pinCode"`
	Mobile     string `json:"mobile"`
	Landline   string `json:"landline"`
}

type MyForms struct {
	FormID      int64  `json:"formID"`
	FormType    int64  `json:"formType"`
	FormName    string `json:"formName"`
	Status      string `json:"status"`
	DateAndTime string `json:"dateAndTime"`
	AreaID      int64  `json:"areaID"`
	AreaName    string `json:"areaName"`
	Comment     string `json:"comment"`
}

type MyFormsResponse struct {
	Status           string    `json:"status"`
	Forms            []MyForms `json:"myForm"`
	ExceptionMessage string    `json:"exceptionMessage"`
}

type UserEmailStruct struct {
	UserEmail string `json:"userEmail"`
}

type EditAccountRequest struct {
	//UserName  string `json:"userName"`
	//UserID    int64  `json:"userID"`
	UserEmail string `json:"userEmail"`
	Password  string `json:"password"`
	//FirstName      string `json:"firstName"`
	//MiddleName     string `json:"middleName"`
	//LastName       string `json:"lastName"`
	//Gender         string `json:"gender"`
	//Date           string `json:"date"`
	//Month          string `json:"month"`
	//Year           string `json:"year"`
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	Address3 string `json:"address3"`
	City     string `json:"city"`
	State    string `json:"state"`
	PinCode  string `json:"pinCode"`
	Mobile   string `json:"mobile"`
	Landline string `json:"landline"`
}

type EditAccountResponse struct {
	Status           string `json:"status"`
	ExceptionMessage string `json:"exceptionMessage"`
}

type MyDocsRequest struct {
	UserEmail      string `json:"userEmail"`
	AadharCard     string `json:"aadharCard"`
	PanCard        string `json:"panCard"`
	Passport       string `json:"passport"`
	DrivingLicense string `json:"drivingLicense"`
	Photo          string `json:"photo"`
	AddressProof   string `json:"addressProof"`
}

type MyDocsResponse struct {
	Status           string `json:"status"`
	ExceptionMessage string `json:"exceptionMessage"`
}

type DuplicateEmail struct {
	IsEmailValid string `json:"isEmailValid"`
}
