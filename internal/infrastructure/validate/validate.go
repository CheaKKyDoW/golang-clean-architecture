package validate

import (
	"errors"
	"fmt"
	"log/slog"
	"reflect"
	"regexp"
	"slices"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator(opt ...validator.Option) *Validator {
	val := validator.New(opt...)
	if err := val.RegisterValidation("digit", validateNumericDigit); err != nil {
		slog.Error("Error registering validation 'digit'", "error", err)
	}

	if err := val.RegisterValidation("nonumeric", validateNoNumeric); err != nil {
		slog.Error("Error registering validation 'nonumeric'", "error", err)
	}

	if err := val.RegisterValidation("validspecificcharacter", validateSpecificCharacter); err != nil {
		slog.Error("Error registering validation 'validspecificcharacter'", "error", err)
	}

	if err := val.RegisterValidation("validthaicharacter", validateThaiOnly); err != nil {
		slog.Error("Error registering validation 'validthaicharacter'", "error", err)
	}
	if err := val.RegisterValidation("prefixnamelabel", validatePrefixNameLabel); err != nil {
		slog.Error("Error registering prefixnamelabel ", "error", err)
	}
	return &Validator{
		validate: val,
	}
}

func (v Validator) Validate(data any) (err error) {
	validateErr := v.validate.Struct(data)
	var validateErrs validator.ValidationErrors
	if errors.As(validateErr, &validateErrs) {
		validationErrors := make([]error, 0, len(validateErrs))
		for _, filedErr := range validateErrs {
			validationErrors = append(validationErrors, filedErr)
		}
		err = errors.Join(validationErrors...)
	}
	return
}

func validateNumericDigit(fl validator.FieldLevel) bool {
	// Get the field value
	value := fl.Field().String()
	// Default regex for numeric values
	regexPattern := `^\d+$`
	if param := fl.Param(); param != "" {
		if length, err := strconv.Atoi(param); err == nil {
			regexPattern = fmt.Sprintf(`^\d{%d}$`, length)
		}
	}
	// re := regexp.MustCompile(fmt.Sprintf(`^\w{%s}$`, fl.Param()))
	re := regexp.MustCompile(regexPattern)
	return re.MatchString(value)
}

func validateSpecificCharacter(fl validator.FieldLevel) bool {
	characterRegex := regexp.MustCompile("^[0-9a-zA-Zก-๙/-]{0,}$")
	return characterRegex.MatchString(valueToString(fl.Field()))
}

func validateThaiOnly(fl validator.FieldLevel) bool {
	var thaiOnlyRegex = regexp.MustCompile("^[ก-๙]+$")
	return thaiOnlyRegex.MatchString(valueToString(fl.Field()))
}

func validateNoNumeric(fl validator.FieldLevel) bool {
	characterRegex := regexp.MustCompile("^[0-9]+$")
	return !characterRegex.MatchString(valueToString(fl.Field()))
}

func validatePrefixNameLabel(fl validator.FieldLevel) bool {
	prefixName := []string{"นาย", "นาง", "นางสาว"}
	return slices.Contains(prefixName, fl.Field().String())
}

func valueToString(rawV reflect.Value) string {
	fieldType := rawV.Type()
	switch fieldType.Name() {
	case "int64":
		return fmt.Sprintf("%d", rawV.Interface())
	case "int":
		return fmt.Sprintf("%d", rawV.Interface())
	case "float64":
		return fmt.Sprintf("%f", rawV.Interface())
	case "float32":
		return fmt.Sprintf("%f", rawV.Interface())
	case "string":
		return rawV.String()
	default:
		return ""
	}
}
