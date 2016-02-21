package models

type UserLoginRequest struct {
	UserEmail string `json:"userEmail"`
	Password  string `json:"password"`
}

type UserLoginResponse struct {
	Status           string `json:"status"`
	UserID           int64  `json:"userID"`
	UserName         string `json:"userName"`
	UserEmail        string `json:"userEmail"`
	ExceptionMessage string `json:"exceptionMessage"`
}

type UserRegistrationRequest struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	UserName       string `json:"userName"`
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
	AddressProof   string `json:"addressProof"`
	Photo          string `json:"photo"`
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

type TenantFormRequest struct {
	UserEmail string       `json:"userEmail"`
	Owner     OwnerData    `json:"owner"`
	Property  PropertyData `json:"property"`
	Tenant    TenantData   `json:"tenantData"`
	Agent     AgentData    `json:"agentData"`
}

type TenantFormResponse struct {
	Status           string `json:"status"`
	ExceptionMessage string `json:"exceptionMessage"`
}
