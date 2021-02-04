package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"

	"github.com/go-playground/validator"
)

type User struct {
	Id           int    `json:"id"`
	Username     string `json:"username" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	FisrtName    string `json:"firstname"`
	Name         string `json:"name"`
	DateOfBirth  string `json:"dateofbirth"`
	Gender       string `json:"gender"`
	Address      string `json:"address"`
	Bio          string `json:"bio"`
	Achievements string `json:"achievements"`
}

// Users is a collection of User
type Users []*User

//Return every users
func GetUsers() Users {
	return userList
}

func AddUser(u *User) {
	u.Id = getNextID()
	userList = append(userList, u)
}

func UpdateUser(id int, u *User) error {
	_, pos, err := findUser(id)
	if err != nil {
		return err
	}

	u.Id = id
	userList[pos] = u

	return nil
}

var ErrUserNotFound = fmt.Errorf("User not found")

func findUser(id int) (*User, int, error) {
	for i, u := range userList {
		if u.Id == id {
			return u, i, nil
		}
	}
	return nil, -1, ErrUserNotFound
}

func getNextID() int {
	lu := userList[len(userList)-1]
	return lu.Id + 1
}

func (u *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

func (u *User) Validate() error {
	validate := validator.New()
	err := validate.RegisterValidation("email", validateEmail)
	if err != nil {
		//do something
	}
	return validate.Struct(u)
}

func validateEmail(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`([a-z0-9][-a-z0-9_\+\.]*[a-z0-9])@([a-z0-9][-a-z0-9\.]*[a-z0-9]\.(arpa|root|aero|biz|cat|com|coop|edu|gov|info|int|jobs|mil|mobi|museum|name|net|org|pro|tel|travel|ac|ad|ae|af|ag|ai|al|am|an|ao|aq|ar|as|at|au|aw|ax|az|ba|bb|bd|be|bf|bg|bh|bi|bj|bm|bn|bo|br|bs|bt|bv|bw|by|bz|ca|cc|cd|cf|cg|ch|ci|ck|cl|cm|cn|co|cr|cu|cv|cx|cy|cz|de|dj|dk|dm|do|dz|ec|ee|eg|er|es|et|eu|fi|fj|fk|fm|fo|fr|ga|gb|gd|ge|gf|gg|gh|gi|gl|gm|gn|gp|gq|gr|gs|gt|gu|gw|gy|hk|hm|hn|hr|ht|hu|id|ie|il|im|in|io|iq|ir|is|it|je|jm|jo|jp|ke|kg|kh|ki|km|kn|kr|kw|ky|kz|la|lb|lc|li|lk|lr|ls|lt|lu|lv|ly|ma|mc|md|mg|mh|mk|ml|mm|mn|mo|mp|mq|mr|ms|mt|mu|mv|mw|mx|my|mz|na|nc|ne|nf|ng|ni|nl|no|np|nr|nu|nz|om|pa|pe|pf|pg|ph|pk|pl|pm|pn|pr|ps|pt|pw|py|qa|re|ro|ru|rw|sa|sb|sc|sd|se|sg|sh|si|sj|sk|sl|sm|sn|so|sr|st|su|sv|sy|sz|tc|td|tf|tg|th|tj|tk|tl|tm|tn|to|tp|tr|tt|tv|tw|tz|ua|ug|uk|um|us|uy|uz|va|vc|ve|vg|vi|vn|vu|wf|ws|ye|yt|yu|za|zm|zw)|([0-9]{1,3}\.{3}[0-9]{1,3}))`)
	matches := re.FindAllString(fl.Field().String(), -1)

	return len(matches) == 1
}

var userList = []*User{
	{
		Id:       12345,
		Username: "sickboy",
	},
	{
		Id:       54321,
		Username: "Mark Renton",
	},
}
