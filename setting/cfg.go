package setting

type AppCfg struct {
	*Evn `mapstructure:"evn" json:"evn" yaml:"evn"`
}
type Evn struct {
	Test string `mapstructure:"test" json:"test" yaml:"test"`
	Cur  string `mapstructure:"cur" json:"cur" yaml:"cur"`
}
