package conf

import (
	"errors"
	"strconv"
	"encoding/json"
	"gopkg.in/yaml.v2"
)

//Conf 用于解析配置数据
type Conf int 

//
const (
	Undefined 	= iota		//NOTE: 未定义的配置格式
	JSON
	YAML
)

//Unmarshal 解析配置到结构体中
func (c *Conf) Unmarshal(in []byte, out interface{}) (err error) {

	switch *c {
	case JSON: err = c.unmarshalJSON(in, out)
	case YAML: err = c.unmarshalYAML(in, out)
	case Undefined: fallthrough
	default:
		err = errors.New("not support type : " + strconv.Itoa(int(*c)))
	}
	return 
}

//UnmarshalToMap 解析配置文件到map中 key - value
func (c *Conf) UnmarshalToMap(in []byte) (out map[string]interface{}, err error) {

	out = make(map[string]interface{})
	err = c.Unmarshal(in, &out)
	if err != nil {
		out = nil
	}
	return
}

func (c *Conf) unmarshalJSON(in []byte, out interface{}) error {

	return json.Unmarshal(in, out)
}

func (c *Conf) unmarshalYAML(in []byte, out interface{}) error {

	return yaml.Unmarshal(in, out)
}
