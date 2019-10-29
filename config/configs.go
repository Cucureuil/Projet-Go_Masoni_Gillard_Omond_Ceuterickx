// tempConf
package config

func GetTemp() *ConfigConnection {
	return &ConfigConnection{
		BrokerPort:    "8080",
		BrokerAddress: "tcp://localhost",
		QoSLevel:      0,
		ClientID:      "ST01",
	}
}

func GetWind() *ConfigConnection {
	return &ConfigConnection{
		BrokerPort:    "8080",
		BrokerAddress: "tcp://localhost",
		QoSLevel:      0,
		ClientID:      "SW01",
	}
}

func GetPress() *ConfigConnection {
	return &ConfigConnection{
		BrokerPort:    "8080",
		BrokerAddress: "tcp://localhost",
		QoSLevel:      0,
		ClientID:      "SP01",
	}
}

func GetAirportA() *ConfigConnection {
	return &ConfigConnection{
		BrokerPort:    "8080",
		BrokerAddress: "tcp://localhost",
		QoSLevel:      0,
		ClientID:      "AAA1",
	}
}

func GetAirportB() *ConfigConnection {
	return &ConfigConnection{
		BrokerPort:    "8080",
		BrokerAddress: "tcp://localhost",
		QoSLevel:      0,
		ClientID:      "BBB1",
	}
}
