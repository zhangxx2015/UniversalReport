package Conf

import (
	"encoding/json"
	"os"
)

// lims presto://192.168.25.99:10010/ilab
// Data Source=.;Initial Catalog=ODFS;User ID=ODFS;Password=Qur#COdOLBzu$clL;Pooling=true; MAX Pool Size=512;Min Pool Size=50;Connection Lifetime=30;
type Setting struct {
	Driver   string // "mysql"
	Username string // "root"
	Password string // "123456"
	Server   string // "localhost"
	Database string // ""
	Test     string
	Debug    bool
}

func Config(conf string) []Setting {
	file, _ := os.Open(conf)
	defer file.Close()
	decoder := json.NewDecoder(file)
	sett := []Setting{}
	err := decoder.Decode(&sett)
	if err != nil {
		panic(err.Error())
	}
	return sett
}
