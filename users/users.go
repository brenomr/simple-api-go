package users

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	Mail      string `json:"mail"`
	Cellphone int    `json:"cellphone"`
}

var Users_example []User

func init() {
	Users_example = []User{
		{ID: "1", Name: "John Doe", Nickname: "Johny", Mail: "johdoe@john.com", Cellphone: 12345678901},
		{ID: "2", Name: "Pamela McCain", Nickname: "Pam", Mail: "pam@pam.com", Cellphone: 11255554444},
		{ID: "3", Name: "Jacky Niccolson", Nickname: "Jack", Mail: "jack@nick.net", Cellphone: 15588887456},
	}
}
