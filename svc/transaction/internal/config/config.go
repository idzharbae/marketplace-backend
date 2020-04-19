package config

type Config struct {
	Grpc struct {
		Port string
	}
	Db struct {
		Debug    bool     `json:"debug"`
		DBEngine string   `json:"db_engine"`
		Master   DbParams `json:"master"`
		Slave    DbParams `json:"slave"`
	}
	Gateways struct {
		Catalog struct {
			Grpc struct {
				Port string
			}
		}
	}
}

type DbParams struct {
	Address  string `json:"address"`
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"db_name"`
	SSLMode  string `json:"ssl_mode"`
	Port     int    `json:"port"`
}