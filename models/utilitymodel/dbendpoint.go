package utilitymodel

//DbEndpoints contains the IP and port to the db server
type DbEndpoints struct {
	IP   string `yaml:"IP"`
	Port string `yaml:"Port"`
}
