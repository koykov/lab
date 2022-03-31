package call_with_options

type BaseOptions struct {
	Flag bool
}

func (o BaseOptions) GetOption(name string) interface{} {
	if name == "flag" {
		return o.Flag
	}
	return nil
}

func (o *BaseOptions) WithFlag(flag bool) OptionsInterface {
	return BaseOptions{Flag: flag}
}

func WithFlag(flag bool) OptionsInterface {
	return BaseOptions{Flag: flag}
}
