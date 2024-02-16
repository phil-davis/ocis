package config

// Log defines the available log configuration.
type Log struct {
	Level  string `mapstructure:"level" env:"OCIS_LOG_LEVEL;THUMBNAILS_LOG_LEVEL" desc:"The log level. Valid values are: 'panic', 'fatal', 'error', 'warn', 'info', 'debug', 'trace'." introductionVersion:"pre5.0"`
	Pretty bool   `mapstructure:"pretty" env:"OCIS_LOG_PRETTY;THUMBNAILS_LOG_PRETTY" desc:"Activates pretty log output." introductionVersion:"pre5.0"`
	Color  bool   `mapstructure:"color" env:"OCIS_LOG_COLOR;THUMBNAILS_LOG_COLOR" desc:"Activates colorized log output." introductionVersion:"pre5.0"`
	File   string `mapstructure:"file" env:"OCIS_LOG_FILE;THUMBNAILS_LOG_FILE" desc:"The path to the log file. Activates logging to this file if set." introductionVersion:"pre5.0"`
}
