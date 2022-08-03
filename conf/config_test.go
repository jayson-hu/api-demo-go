package conf_test

import (
	"os"
	"testing"

	"github.com/jayson-hu/api-demo-go/conf"
	"github.com/stretchr/testify/assert"
)

func TestConfigFromToml(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/demo.toml")
	if should.NoError(err) {
		should.Equal("demo1", conf.C().App.Name)
	}
}

func TestConfigFromEnv(t *testing.T) {
	os.Setenv("MYSQL_DATABASE", "unit_test")
	should := assert.New(t)
	err := conf.LoadConfigFromTEnv()
	if should.NoError(err) {
		should.Equal("unit_test", conf.C().MySQL.Database)
	}

}
func TestGetDB(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/demo.toml")
	if should.NoError(err) {
		conf.C().MySQL.GetDB()

	}

}