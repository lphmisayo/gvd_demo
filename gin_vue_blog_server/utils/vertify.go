package utils

const (
	ParamError   = "ParamError"
	StructError  = "StructError"
	CaptchaError = "CaptchaError"
	NoError      = "NoError"
)

func (l *Login) LoginStVertify(isDefult bool) (string, bool) {
	if l != nil {
		if l.Username == "" {
			return ParamError, false
		}
		if l.Password == "" {
			return ParamError, false
		}
		if isDefult {
			if l.Captcha == "" || l.CaptchaId == "" {
				return CaptchaError, false
			}
		}
	} else {
		return StructError, false
	}
	return NoError, true
}

func (r *Register) RegistStVertify() (string, bool) {
	if r != nil {
		if r.Username == "" {
			return ParamError, false
		}
		if r.Password == "" {
			return ParamError, false
		}
	} else {
		return StructError, false
	}
	return NoError, true
}
