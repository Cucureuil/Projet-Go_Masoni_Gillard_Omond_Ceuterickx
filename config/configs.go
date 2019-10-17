// tempConf
package config

func GetTemp() *Config {
	return &Config{
		BrokerPort:    "1883",
		BrokerAddress: "tcp://localhost",
		QoSLevel:      0,
		ClientID:      "ST01",
	}
}

func GetWind() *Config {
	return &Config{
		BrokerPort:    "1883",
		BrokerAddress: "tcp://localhost",
		QoSLevel:      0,
		ClientID:      "SW01",
	}
}

func GetPress() *Config {
	return &Config{
		BrokerPort:    "1883",
		BrokerAddress: "tcp://localhost",
		QoSLevel:      0,
		ClientID:      "SP01",
	}
}
