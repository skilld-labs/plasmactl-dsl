package definition

import (
	"github.com/go-playground/validator/v10"
	"log"
	"testing"
)

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())

	for tag, validatorFunc := range DefaultValidatorFuncs {
		if err := validate.RegisterValidation(tag, validatorFunc); err != nil {
			log.Fatalf("error registration custom validation function '%s': '%s'", tag, err.Error())
		}
	}
}

func TestValidateMilliCPU(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{"Valid CPU - 64", "64", true},
		{"Valid CPU - 128", "128", true},
		{"Valid CPU - 512m", "512m", true},
		{"Valid CPU - 1m", "1m", true},
		{"Valid CPU - 0.001", "0.001", true},
		{"Invalid CPU - 0", "0", false},
		{"Invalid CPU - m0", "m0", false},
		{"Invalid CPU - Too small", "0.0001", false},
		{"Invalid CPU - negative cpu", "-1", false},
		{"Invalid CPU - negative mCPU", "-1m", false},
		{"Invalid CPU - fractional CPU", "-0.0001", false},
		{"Invalid CPU - negative fractional CPU", "-0.0001", false},
		{"Invalid CPU - wrong case", "M1", false},
		{"Invalid CPU - wrong formatting", "mi1", false},
	}

	for _, tt := range tests {
		err := validate.Var(tt.value, "milliCPU")
		if (err == nil) != tt.want {
			t.Errorf("%s: expected %v, got %v", tt.name, tt.want, err == nil)
		}
	}
}

func TestValidateInfoUnits(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{"Valid RAM - 128Mi", "128Mi", true},
		{"Valid RAM - 2Gi", "2Gi", true},
		{"Invalid RAM - 128", "128", false},
		{"Invalid RAM - random string", "abc", false},
		{"Valid Storage - 1Ti", "1Ti", true},
		{"Invalid Storage - -100Gi", "-100Gi", false},
		{"Invalid Storage - 0Gi", "0Gi", false},
		{"Invalid Storage - 0", "0", false},
		{"Invalid Storage - 01", "01", false},
		{"Invalid Storage - 01Gi", "01Gi", false},
		{"Invalid Storage - 1i", "1i", false},
	}

	for _, tt := range tests {
		err := validate.Var(tt.value, "infoUnits")
		if (err == nil) != tt.want {
			t.Errorf("%s: expected %v, got %v", tt.name, tt.want, err == nil)
		}
	}
}

func TestValidateTimeInterval(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{"Valid Interval - 60s", "60s", true},
		{"Valid Interval - 1m", "1m", true},
		{"Valid Interval - 1h", "1h", true},
		{"Invalid Interval - 1min", "1min", false},
		{"Invalid Interval - 10seconds", "10seconds", false},
		{"Invalid Interval - random string", "abc", false},
	}

	for _, tt := range tests {
		err := validate.Var(tt.value, "timeInterval")
		if (err == nil) != tt.want {
			t.Errorf("%s: expected %v, got %v", tt.name, tt.want, err == nil)
		}
	}
}

func TestValidateTCPv4Port(t *testing.T) {
	tests := []struct {
		name  string
		value uint64
		want  bool
	}{
		{"Valid TCPv4Port - 1", 1, true},
		{"Valid TCPv4Port - 22", 22, true},
		{"Valid TCPv4Port - 65535", 65535, true},
		{"Invalid TCPv4Port - 0", 0, false},
		{"Invalid TCPv4Port - 65536", 65536, false},
	}

	for _, tt := range tests {
		err := validate.Var(tt.value, "isTCPv4Port")
		if (err == nil) != tt.want {
			t.Errorf("%s: expected %v, got %v", tt.name, tt.want, err == nil)
		}
	}
}

func TestValidateServiceBuilderContextType(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{"Valid ServiceBuilderContextType - cluster", "cluster", true},
		{"Valid ServiceBuilderContextType - image", "image", true},
		{"Invalid ServiceBuilderContextType - random", "random", false},
	}

	for _, tt := range tests {
		err := validate.Var(tt.value, "isServiceBuilderCtx")
		if (err == nil) != tt.want {
			t.Errorf("%s: expected %v, got %v", tt.name, tt.want, err == nil)
		}
	}
}

func TestValidateServiceScope(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{"Valid ServiceScope - local", "local", true},
		{"Valid ServiceScope - cluster", "cluster", true},
		{"Invalid ServiceScope - random", "random", false},
	}

	for _, tt := range tests {
		err := validate.Var(tt.value, "isServiceScope")
		if (err == nil) != tt.want {
			t.Errorf("%s: expected %v, got %v", tt.name, tt.want, err == nil)
		}
	}
}

func TestValidateFlowType(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{"Valid FlowType - lake", "lake", true},
		{"Valid FlowType - data", "data", true},
		{"Invalid FlowType - random", "random", false},
	}

	for _, tt := range tests {
		err := validate.Var(tt.value, "isFlowType")
		if (err == nil) != tt.want {
			t.Errorf("%s: expected %v, got %v", tt.name, tt.want, err == nil)
		}
	}
}
