package models

type UserLoginRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Status           string `json:"status"`
	AreaCode         int64  `json:"areaCode"`
	ExceptionMessage string `json:"exceptionMessage"`
}

type FormThumbnail struct {
	FormType   int64  `json:"formType"`
	FormID     int64  `json:"formID"`
	Status     int64  `json:"status"`
	UserID     int64  `json:"userID"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
	DateTime   string `json:"dateTime"`
}

type FormThumbnailsRequest struct {
	AreaCode string `json:"areaCode"`
}

type FormThumbnailsResponse struct {
	FormThumbnails   []FormThumbnail `json:"formThumbnails"`
	ExceptionMessage string          `json:"exceptionMessage"`
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

type TenantFormResponse struct {
	Owner            OwnerData    `json:"owner"`
	Property         PropertyData `json:"property"`
	Tenant           TenantData   `json:"tenantData"`
	Agent            AgentData    `json:"agentData"`
	ExceptionMessage string       `json:"exceptionMessage"`
}

type TenantFormRequest struct {
	FormID   int64 `json:"formID"`
	FormType int64 `json:"formType"`
}
