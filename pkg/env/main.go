package env

import(
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func Get(name, def string) string {

	if e := os.Getenv(name); e != "" {
		return e
	}

	return def
}