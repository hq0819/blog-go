package configuration

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"sync"
)

var DB *gorm.DB
var once sync.Once

type Datasource struct {
	IP       string `yaml:"ip" json:"ip"`
	Port     string `yaml:"port" json:"port"`
	Database string `yaml:"database" json:"database"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}

func GetDbInstance() *gorm.DB {
	once.Do(func() {
		getenv := os.Getenv("profile")
		v := viper.New()
		v.AddConfigPath("configuration")
		v.SetConfigType("yml")
		v.SetConfigName("db-" + getenv)
		err := v.ReadInConfig()
		if err != nil {
			log.Print(err)
		}
		Datasource := &Datasource{}
		err = v.UnmarshalKey("Datasource", Datasource)
		if err != nil {
			log.Print(err)
		}
		dns := Datasource.Username + `:` + Datasource.Password +
			`@tcp` + `(` + Datasource.IP + `:` + Datasource.Port + `)` + `/` + Datasource.Database + `?parseTime=True`
		config := gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: false,
			},
		}
		DB, _ = gorm.Open(mysql.Open(dns), &config)
	})

	return DB

}
