package config

// Debug defines the available debug configuration.
type Debug struct {
	Addr   string `yaml:"addr" env:"OCM_DEBUG_ADDR" desc:"Bind address of the debug server, where metrics, health, config and debug endpoints will be exposed."`
	Token  string `yaml:"token" env:"OCM_DEBUG_TOKEN" desc:"Token to secure the metrics endpoint."`
	Pprof  bool   `yaml:"pprof" env:"OCM_DEBUG_PPROF" desc:"Enables pprof, which can be used for profiling."`
	Zpages bool   `yaml:"zpages" env:"OCM_DEBUG_ZPAGES" desc:"Enables zpages, which can be used for collecting and viewing in-memory traces."`
}
