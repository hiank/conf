package conf_test

import (
	"github.com/hiank/conf"
	"testing"
)

type Sys struct {
	WsPort 		int64 		`json:"sys.wsPort"`
	K8sPort 	int64 		`json:"sys.k8sPort"`
}

type JsonConf struct {
	Wei 		string 		`json:"jsonConf.wei"`
	Shao 		string 		`json:"jsonConf.shao"`
}

var defaultSysConf = `
{
	"sys.wsPort": 30250,
	"sys.k8sPort": 30260
}
`

func TestJson(t *testing.T) {

	sysInfo, err := conf.NewInfoFromFile("conf_test.json", &Sys{})
	if err != nil {
		t.Error(err)
		return
	}

	cfg := conf.NewConf()
	cfg.Load(sysInfo)
	if sys, ok := cfg.Get("Sys"); ok {
		t.Logf("wsport : %d, k8sport : %d\n", sys.(*Sys).WsPort, sys.(*Sys).K8sPort)
	}

	jcInfo, err := conf.NewInfoFromFile("conf_test.json", &JsonConf{})
	if err != nil {
		t.Error(err)
		return
	}
	cfg.Load(jcInfo)
	if jsonConf, ok := cfg.Get("JsonConf"); ok {
		t.Log("name :", jsonConf.(*JsonConf).Wei, jsonConf.(*JsonConf).Shao)
	}

	bitsInfo := conf.NewInfoFromBits([]byte(defaultSysConf), &Sys{}, conf.TypeJSON)
	cfg.Load(bitsInfo)
	if bitsConf, ok := cfg.Get("Sys"); ok {
		t.Logf("bits wsport : %d, k8sport : %d\n", bitsConf.(*Sys).WsPort, bitsConf.(*Sys).K8sPort)
	}
}