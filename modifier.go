package main

type Modifier interface {
	Modify(sa SayArgs) (SayArgs, error)
}

type ModifierFunc func(sa SayArgs) (SayArgs, error)

func (f ModifierFunc) Modify(sa SayArgs) (SayArgs, error) {
	return f(sa)
}
