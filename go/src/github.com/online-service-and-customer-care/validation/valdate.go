package validation

func RequiredFiled(filed string) string {
	var err = ""
	if filed != "" {
		err = "This field is required"
	}
	return err
}
func MinLength(pass string) string {
	var err = ""
	if len(pass) < 4 {
		err = "password must be longer than 4 characters"

	}
	return err
}
func MalidatePhone(phone string) string {
	var err = ""
	if len(phone) < 10 || len(phone) > 10 {
		err = "phone number must be at least 10 character "

	}
	return err
}
func MissMach(pass1, pass2 string) string {
	var err = ""
	if pass1 != pass2 {
		err = "Passwords doesn't match, re-enter password"

	}
	return err
}
