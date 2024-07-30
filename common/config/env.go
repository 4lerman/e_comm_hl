package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	API_Port      string
	Users_Port    string
	Products_Port string
	Orders_Port   string
	Payments_Port string

	DBUser     string
	DBPassword string
	DBAddress  string
	DBPort     int64
	DBName     string

	Base_Url     string
	Users_Url    string
	Products_Url string
	Orders_Url  string
	Payments_Url string

	Token_Url        string
	Make_Payment_Url string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		API_Port:      getEnv("API_PORT", "8080"),
		Users_Port:    getEnv("USERS_PORT", "8081"),
		Products_Port: getEnv("PRODUCTS_PORT", "8082"),
		Orders_Port:   getEnv("ORDERS_PORT", "8083"),
		Payments_Port: getEnv("PAYMENTS_PORT", "8084"),

		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "mypassword"),
		DBAddress:  getEnv("DB_HOST", "127.0.0.1"),
		DBPort:     getEnvAsInt("DB_PORT", 5432),
		DBName:     getEnv("DB_NAME", "ecom"),

		Base_Url:     getEnv("BASE_URL", "http://localhost:8080/"),
		Users_Url:    getEnv("USERS_URL", "http://localhost:8081/"),
		Products_Url: getEnv("PRODUCTS_URL", "http://localhost:8082/"),
		Orders_Url:  getEnv("ORDERS_URL", "http://localhost:8083/"),
		Payments_Url: getEnv("PAYMENTS_URL", "http://localhost:8084/"),

		Token_Url:        getEnv("TOKEN_URL", "https://testoauth.homebank.kz/epay2/oauth2/token"),
		Make_Payment_Url: getEnv("MAKE_PAYMENT_URL", "https://testepay.homebank.kz/api/payment/cryptopay"),
	}
}

// Gets the env by key or fallbacks
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
