package main

import "regexp"

var (
	asciiRegexp = regexp.MustCompile(`^[0x21-0x7e\s]+$`)
)

type Modifier interface {
	Modify(sa SayArgs) (SayArgs, error)
}

type ModifierFunc func(sa SayArgs) (SayArgs, error)

func (f ModifierFunc) Modify(sa SayArgs) (SayArgs, error) {
	return f(sa)
}

type VoiceLanguageModifier struct {
	Config SayConfig
}

func (vlm VoiceLanguageModifier) Modify(sa SayArgs) (SayArgs, error) {
	if asciiRegexp.MatchString(sa.Text) {
		sa.Voice = vlm.Config.Voice.En
	}
	return sa, nil
}
