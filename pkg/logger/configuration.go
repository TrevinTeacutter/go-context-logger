package logger

type Configuration struct {
	Format       string `mapstructure:"format" json:"format" yaml:"format"`
	Verbose      bool   `mapstructure:"verbose" json:"verbose" yaml:"verbose"`
	OmitMetadata bool   `mapstructure:"omitMetadata" json:"omitMetadata" yaml:"omitMetadata"`
	Flatten      bool   `mapstructure:"flatten" json:"flatten" yaml:"flatten"`
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Separator    string `mapstructure:"separator" json:"separator" yaml:"separator"`
	MessageKey   string `mapstructure:"messageKey" json:"messageKey" yaml:"messageKey"`
}
