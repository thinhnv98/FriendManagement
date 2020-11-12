package utils

import "regexp"

const EmailValidationRegex = "[_A-Za-z0-9-\\+]+(\\.[_A-Za-z0-9-]+)*@[A-Za-z0-9-]+(\\.[A-Za-z0-9]+)*(\\.[A-Za-z]{2,})"

func IsValidEmail(email string) (bool, error) {
	re, err := regexp.Compile(EmailValidationRegex)
	if err != nil {
		return false, err
	}
	if re.MatchString(email) {
		return true, nil
	}
	return false, nil
}

func FindEmailFromText(text string) []string {
	regex := regexp.MustCompile(`[_A-Za-z0-9-\+]+(\.[_A-Za-z0-9-]+)*@[A-Za-z0-9-]+(\.[A-Za-z0-9]+)*(\.[A-Za-z]{2,})`)

	emailChain := regex.FindAllString(text, -1)
	email := make([]string, len(emailChain))
	for index, emailCharacter := range emailChain {
		email[index] = emailCharacter
	}
	return email
}
