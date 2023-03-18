package iop

type Config struct {
	BindAddress string
	Contracts   ContractsConfig
	Ethereum    EthereumConfig
	IOTA        IOTAConfig
}

type ContractsConfig struct {
	RegistryContractAddress string
	OracleContractAddress   string
}

type EthereumConfig struct {
	TargetAddress string
	SourceAddress string
	PrivateKey    string
	ChainID       int64
}

type IOTAConfig struct {
	Rest  string
	Mqtt  string
	Topic string
}
