package data

type MyGXIdentifierData struct {
	ID               int                   `json:"id"`
	UUID             string                `json:"uuid"`
	CompanyOffice    mygxCompanyOffice     `json:"companyOffice"`
	AccountNumber    string                `json:"accountNumber"`
	HolderName       string                `json:"holderName"`
	Status           idName                `json:"status"`
	Type             idName                `json:"type"`
	User             mygxUser              `json:"user"`
	ContactInfo      mygxContactInfo       `json:"contactInfo"`
	ServiceLocations []mygxServiceLocation `json:"serviceLocations"`
}

type mygxUser struct {
	ID       string  `json:"id"`
	Username string  `json:"username"`
	Email    *string `json:"email"`
}

type mygxCompanyOffice struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type mygxContactInfo struct {
	Email       *string `json:"email"`
	HomePhone   *string `json:"homePhone"`
	MobilePhone *string `json:"mobilePhone"`
	Gender      *string `json:"gender"`
	Nationality *string `json:"nationality"`
	Address     *string `json:"address"`
	City        *string `json:"city"`
	Area        *string `json:"area"`
}

type mygxServiceLocation struct {
	ID         int         `json:"id"`
	UUID       string      `json:"uuid"`
	Status     idName      `json:"status"`
	Latitude   *string     `json:"latitude"`
	Longitude  *string     `json:"longitude"`
	StreetName *string     `json:"streetName"`
	LocationId string      `json:"locationId"`
	Nickname   *string     `json:"nickname"`
	Package    mygxPackage `json:"package"`
}

type mygxPackage struct {
	ID             int     `json:"id"`
	UUID           string  `json:"uuid"`
	Name           string  `json:"name"`
	IdentifierName *string `json:"identifierName"`
	Alias          *string `json:"alias"`
}
