package conf

import (
	"gotest.tools/v3/assert"
	"testing"
)

var jsonIn = `
	{"ip": "192.168.1.1", "port": 1024}
`

func TestJsonConf(t *testing.T) {

	c := Conf(JSON)
	m, err := c.UnmarshalToMap([]byte(jsonIn))
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, int(m["port"].(float64)), 1024)
	assert.Equal(t, m["ip"], "192.168.1.1")
}

var yamlIn = `
ip: 192.168.1.1
port: 1024
`

func TestYamlConf(t *testing.T) {

	c := Conf(YAML)
	m, err := c.UnmarshalToMap([]byte(yamlIn))
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, m["port"], 1024)
	assert.Equal(t, m["ip"], "192.168.1.1")
}

type yamlConf struct {
	IP   string `yaml:"ip"`
	Port int
}

func TestYamlConfToStruct(t *testing.T) {

	c := Conf(YAML)
	var val yamlConf
	err := c.Unmarshal([]byte(yamlIn), &val)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, val.Port, 1024)
	assert.Equal(t, val.IP, "192.168.1.1")
}
