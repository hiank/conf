package conf

import (
	"strconv"
	"errors"
	"strings"
	"reflect"
	"encoding/json"
	"io/ioutil"
)

//Conf is json format setting
type Conf struct {
	
	hub 	map[string]interface{} 		//NOTE: map[key]ItemParser 
}

//NewConf 创建一个新的Conf
func NewConf() *Conf {

	return &Conf{
		hub : make(map[string]interface{}),
	}
}

//Get 根据配置数据结构获得相关解析器
func (conf *Conf) Get(ns string) (val interface{}, ok bool) {

	val, ok = conf.hub[ns]
	return
}

//Load 根据配置加载设置
func (conf *Conf) Load(ci *Info) (err error) {
	
	t := reflect.TypeOf(ci.val)
	if t.Kind() == reflect.Ptr {

		t = t.Elem()
	}
	ns := t.Name()
	val, ok := conf.Get(ns)
	if !ok {

		val = ci.val
		conf.hub[ns] = val
	}
	switch ci.format {
	case TypeJSON:
		err = json.Unmarshal(ci.buf, val)
	case TypeYAML:
		err = errors.New("not support yaml now")
	default:
		err = errors.New("cann't parse config file typed " + strconv.Itoa(ci.format))
	}
	return
}

//Type 配置文件类型
const (
	TypeJSON	= iota				//NOTE: json类型，默认类型
	TypeYAML						//NOTE: yaml类型
)

//Info 用于构成配置文件信息
type Info struct {

	buf 		[]byte
	val 		interface{}
	format 		int
}

//NewInfoFromFile 新建一个ConfInfo
func NewInfoFromFile(fPath string, val interface{}) (info *Info, err error) {

	idx := strings.LastIndexByte(fPath, '.')
	if idx == -1 {
		err = errors.New("failed file suffix : cann't find filename suffix")
		return
	}
	var format int
	switch strings.ToLower(fPath[idx+1:]) {
	case "json": format = TypeJSON
	case "yaml": format = TypeYAML
	default:
		err = errors.New("failed file suffix : cann't parse filename suffixed " + fPath[idx+1:])
		return
	}

	var bits []byte
	if bits, err = ioutil.ReadFile(fPath); err == nil {

		info = NewInfoFromBits(bits, val, format)
	}
	return
}

//NewInfoFromBits 创建配置信息
func NewInfoFromBits(bits []byte, val interface{}, format int) *Info {

	return &Info{bits, val, format}
}


// var _conf *conf
// var mtx sync.RWMutex
// //Init 初始化k8s conn pool
// func Init(cis ...*Info) {

// 	mtx.Lock()
// 	if _conf == nil {

// 		_conf = &conf{make(map[string]interface{})}
// 	}
// 	for _, ci := range cis {

// 		_conf.load(ci)
// 	}
// 	mtx.Unlock()
// }

// // Get 获得全局Conf, ns必须是struct 的完整名称，大小写一致
// func Get(ns string) (interface{}, bool) {

// 	mtx.RLock()

// 	val, ok := _conf.get(ns)
	
// 	mtx.RUnlock()
// 	return val, ok
// }
