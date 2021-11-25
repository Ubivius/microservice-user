package data

import (
	"regexp"

	"github.com/go-playground/validator"
)

func (user *User) Validate() error {
	validate := validator.New()
	err1 := validate.RegisterValidation("email", validateEmail)

	err2 := validate.RegisterValidation("dateofbirth", validateDateOfBirth)
	err3 := validate.RegisterValidation("isStatusType", validateIsStatusType)
	if err1 != nil || err2 != nil || err3 != nil {
		panic("Validator connexions failed")
	}
	return validate.Struct(user)
}

func validateEmail(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`([a-z0-9][-a-z0-9_\+\.]*[a-z0-9])@([a-z0-9][-a-z0-9\.]*[a-z0-9]\.(arpa|root|aero|biz|cat|com|coop|edu|gov|info|int|jobs|mil|mobi|museum|name|net|org|pro|tel|travel|ac|ad|ae|af|ag|ai|al|am|an|ao|aq|ar|as|at|au|aw|ax|az|ba|bb|bd|be|bf|bg|bh|bi|bj|bm|bn|bo|br|bs|bt|bv|bw|by|bz|ca|cc|cd|cf|cg|ch|ci|ck|cl|cm|cn|co|cr|cu|cv|cx|cy|cz|de|dj|dk|dm|do|dz|ec|ee|eg|er|es|et|eu|fi|fj|fk|fm|fo|fr|ga|gb|gd|ge|gf|gg|gh|gi|gl|gm|gn|gp|gq|gr|gs|gt|gu|gw|gy|hk|hm|hn|hr|ht|hu|id|ie|il|im|in|io|iq|ir|is|it|je|jm|jo|jp|ke|kg|kh|ki|km|kn|kr|kw|ky|kz|la|lb|lc|li|lk|lr|ls|lt|lu|lv|ly|ma|mc|md|mg|mh|mk|ml|mm|mn|mo|mp|mq|mr|ms|mt|mu|mv|mw|mx|my|mz|na|nc|ne|nf|ng|ni|nl|no|np|nr|nu|nz|om|pa|pe|pf|pg|ph|pk|pl|pm|pn|pr|ps|pt|pw|py|qa|re|ro|ru|rw|sa|sb|sc|sd|se|sg|sh|si|sj|sk|sl|sm|sn|so|sr|st|su|sv|sy|sz|tc|td|tf|tg|th|tj|tk|tl|tm|tn|to|tp|tr|tt|tv|tw|tz|ua|ug|uk|um|us|uy|uz|va|vc|ve|vg|vi|vn|vu|wf|ws|ye|yt|yu|za|zm|zw)|([0-9]{1,3}\.{3}[0-9]{1,3}))`)
	matches := re.FindAllString(fl.Field().String(), -1)

	return len(matches) == 1
}

func validateDateOfBirth(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[0-9]{1,2}/[0-9]{1,2}/[0-9]{4}`)
	matches := re.FindAllString(fl.Field().String(), -1)

	return len(matches) == 1
}

// validates the status type is valid
func validateIsStatusType(fieldLevel validator.FieldLevel) bool {
	statusType := fieldLevel.Field().String()

	switch statusType {
    case string(Online), string(Offline), string(InGame):
        return true
    }
	return false
}
