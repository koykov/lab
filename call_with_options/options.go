package call_with_options

import "github.com/koykov/inspector"

type OptionsInterface interface {
	GetOption(string) interface{}
}

type Options struct {
	BaseOptions
	ins inspector.Inspector
}

func (o *Options) WithInspector(ins inspector.Inspector) OptionsInterface {
	o.ins = ins
	return o
}

func (o Options) GetOption(name string) interface{} {
	if name == "inspector" {
		return o.ins
	}
	return o.BaseOptions.GetOption(name)
}

func WithInspector(ins inspector.Inspector) OptionsInterface {
	return Options{ins: ins}
}
