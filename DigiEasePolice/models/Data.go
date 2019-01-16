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
	ApplicationID string `json:"applicationID"`
	UserID        int64  `json:"userID"`
	FormID        int64  `json:"formID"`
	FormType      int64  `json:"formType"`
	Status        int64  `json:"status"`
	DateTime      string `json:"dateTime"`
	Area          int64  `json:"area"`
	FirstName     string `json:"firstName"`
	MiddleName    string `json:"middleName"`
	LastName      string `json:"lastName"`
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
	Age        int64  `json:"age"`
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
	Age              int64  `json:"age"`
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
	Age        int64  `json:"age"`
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
	Status        int64        `json:"status"`
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

type GetFormRequest struct {
	FormID   string `json:"formID"`
	FormType string `json:"formType"`
}

type UpdateStatusRequest struct {
	FormID         string `json:"formID"`
	Status         int64  `json:"status"`
	Discard        string `json:"discard"`
	Comment        string `json:"comment"`
	AreaCode       string `json:"areaCode"`
	AreaCodeStatus string `json:"areaCodeStatus"`
}

type UpdateStatusResponse struct {
	Status           string `json:"status"`
	ExceptionMessage string `json:"exceptionMessage"`
}
