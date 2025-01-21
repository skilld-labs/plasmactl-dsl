package definition

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
	"regexp"
	"strconv"
	"strings"
)

// Generic validation function that accepts an enum type
func validateEnum(enum Enum) func(fl validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		return enum.Exists(value)
	}
}

func validateMilliCPU(fl validator.FieldLevel) bool {
	cpu := fl.Field().String()

	// Check if value is in millis
	if strings.HasSuffix(cpu, "m") {
		cpu = strings.TrimSuffix(cpu, "m")
		intNum, err := strconv.Atoi(cpu)
		if err != nil || intNum <= 0 {
			return false
		}
		if err == nil {
			return true
		}
	}

	// value cannot have precision lower 1 mCPU (0.001 CPU)
	// According to docs https://kubernetes.io/docs/tasks/configure-pod-container/assign-cpu-resource/#cpu-units
	floatNum, err := strconv.ParseFloat(cpu, 64)
	if err != nil || floatNum < 0.001 {
		return false
	}

	return true
}

func validateInfoUnits(fl validator.FieldLevel) bool {
	ram := fl.Field().String()

	// Regular expression to match valid Kubernetes memory formats like "128Mi", "2Gi", etc.
	// Excluding zero as the first character
	re := regexp.MustCompile(`^[1-9][0-9]{0,3}(Ei|Pi|Ti|Gi|Mi|Ki|EiB|PiB|TiB|GiB|MiB|KiB)$`)

	return re.MatchString(ram)
}

func validateTimeInterval(fl validator.FieldLevel) bool {
	interval := fl.Field().String()

	// Regular expression to match valid time intervals like "1min", "60s", "1h", "1d", "1w". Default is min if no time units specified
	// Not excluding zero as the first character, because some software may interpret it as a valid or none value
	re := regexp.MustCompile(`^[0-9]+(s|m|h|d|w)?$`)

	return re.MatchString(interval)
}

func validateTCPv4Port(fl validator.FieldLevel) bool {
	portNum := fl.Field().Uint()

	var maxPortNum uint64 = (2 << 15) - 1

	if portNum == 0 || portNum > maxPortNum {
		return false
	}
	return true
}

var DefaultValidatorFuncs = map[string]validator.Func{
	"isKind":              validateEnum(EnumKind),
	"isPcl":               validateEnum(EnumPcl),
	"isConcurrencyMode":   validateEnum(EnumConcurrencyMode),
	"isFlowType":          validateEnum(EnumFlowType),
	"isServiceScope":      validateEnum(EnumServiceScope),
	"isServiceBuilderCtx": validateEnum(EnumServiceBuilderContext),
	"isSkillStage":        validateEnum(EnumSkillStage),
	"isTCPv4Port":         validateTCPv4Port,
	"milliCPU":            validateMilliCPU,
	"infoUnits":           validateInfoUnits,
	"timeInterval":        validateTimeInterval,
}

type validationError struct {
	Namespace       string `yaml:"namespace"` // can differ when a custom TagNameFunc is registered or
	Field           string `yaml:"field"`     // by passing alt name to ReportError like below
	StructNamespace string `yaml:"structNamespace"`
	StructField     string `yaml:"structField"`
	Tag             string `yaml:"tag"`
	ActualTag       string `yaml:"actualTag"`
	Kind            string `yaml:"kind"`
	Type            string `yaml:"type"`
	Value           string `yaml:"value"`
	Param           string `yaml:"param"`
	Message         string `yaml:"message"`
}

type Validator[T Schema] struct {
	validator *validator.Validate
}

func NewValidator[T Schema]() (*Validator[T], error) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	// Register custom validation functions
	for tag, validatorFunc := range DefaultValidatorFuncs {
		if err := validate.RegisterValidation(tag, validatorFunc); err != nil {
			return nil, fmt.Errorf("error registration custom validation function %s: %s", tag, err.Error())
		}
	}

	return &Validator[T]{validator: validate}, nil
}

func (v *Validator[T]) Validate(schema *T) error {
	err := v.validator.Struct(schema)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err.(*validator.InvalidValidationError)
		}

		for _, err := range err.(validator.ValidationErrors) {
			e := validationError{
				Namespace:       err.Namespace(),
				Field:           err.Field(),
				StructNamespace: err.StructNamespace(),
				StructField:     err.StructField(),
				Tag:             err.Tag(),
				ActualTag:       err.ActualTag(),
				Kind:            fmt.Sprintf("%v", err.Kind()),
				Type:            fmt.Sprintf("%v", err.Type()),
				Value:           fmt.Sprintf("%v", err.Value()),
				Param:           err.Param(),
				Message:         err.Error(),
			}

			yamlOutput, err := yaml.Marshal(e)
			if err != nil {
				return fmt.Errorf("error while log yaml marshalling: %s", err.Error())
			}

			fmt.Println(string(yamlOutput))
		}
	}
	return err
}
