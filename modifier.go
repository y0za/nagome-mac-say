package main

import "regexp"

var (
	asciiRegexp = regexp.MustCompile(`^[\x21-\x7e\s]+$`)
	urlRegexp   = regexp.MustCompile(`https?:\/\/[^\s　]+`)
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

type OverflowModifier struct {
	Limit int
}

func (om OverflowModifier) Modify(sa SayArgs) (SayArgs, error) {
	r := []rune(sa.Text)
	if len(r) > om.Limit {
		sa.Text = string(r[0:om.Limit])
	}
	return sa, nil
}

func OmitURL(sa SayArgs) (SayArgs, error) {
	sa.Text = urlRegexp.ReplaceAllString(sa.Text, "URL省略")
	return sa, nil
}
