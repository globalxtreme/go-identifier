package data

type MyGXIdentifierData struct {
	ID           int              `json:"id"`
	UUID         string           `json:"uuid"`
	Number       string           `json:"number"`
	Branches     []mygxBranch     `json:"branches"`
	FullName     string           `json:"fullName"`
	Emails       []mygxEmail      `json:"emails"`
	Phones       []mygxPhone      `json:"phones"`
	BirthDate    string           `json:"birthDate"`
	Priority     mygxPriority     `json:"priority"`
	SocialStatus mygxSocialStatus `json:"socialStatus"`
	Address      string           `json:"address"`
	Nationality  mygxNationality  `json:"nationality"`
	User         mygxUser         `json:"user"`
}

type mygxBranch struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type mygxEmail struct {
	ID      int    `json:"id"`
	Email   string `json:"email"`
	Primary bool   `json:"primary"`
}

type mygxPhone struct {
	ID        int    `json:"id"`
	PhoneCode string `json:"phoneCode"`
	Phone     string `json:"phone"`
	Primary   bool   `json:"primary"`
}
type mygxPriority struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Alias *string `json:"code"`
}
type mygxSocialStatus struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type mygxNationality struct {
	ID   int     `json:"id"`
	Name string  `json:"name"`
	Code *string `json:"code"`
}
type mygxUser struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Email    *string `json:"email"`
	Primary  *bool   `json:"primary"`
}
