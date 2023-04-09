package config_model

type Server struct {
	AppMode   string `yaml:"AppMode"`
	BackPort  string `yaml:"BackPort"`
	FrontPort string `yaml:"FrontPort"`
}
