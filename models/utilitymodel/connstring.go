package utilitymodel

//ConnectionString Object mongo Database
type ConnectionString struct {
	UserName     string        `yaml:"Username"`
	Password     string        `yaml:"Password"`
	Endpoints    []DbEndpoints `yaml:"DbEndpoints"`
	DatabaseName string        `yaml:"DatabaseName"`
}
