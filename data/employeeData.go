package data

type EmployeeIdentifierData struct {
	ID               string        `json:"id"`
	FullName         string        `json:"fullName"`
	EmployeeNo       string        `json:"employeeNo"`
	Email            string        `json:"email"`
	Superadmin       bool          `json:"superadmin"`
	Department       *idName       `json:"department"`
	Division         *idName       `json:"division"`
	User             user          `json:"user"`
	CompanyOffice    companyOffice `json:"companyOffice"`
	CompanyOfficeIds []int         `json:"companyOfficeIds"`
}

type idName struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type user struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type companyOffice struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Company  idName `json:"company"`
	Location idName `json:"location"`
}
