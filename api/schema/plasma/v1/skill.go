package v1

type Skill struct {
	Core          `yaml:",inline" validate:"required"`
	SkillSettings `yaml:"settings" validate:"required"`
}

type SkillSettings struct {
	SkillValidationRequired bool   `yaml:"skillValidationRequired"`
	SkillStage              string `yaml:"skillStage" validate:"isSkillStage"`
	//		"\(pcsn)_log_verbosity" `yaml:"logVerbosity""`
}

func (s Skill) GetCore() *Core {
	return &s.Core
}
