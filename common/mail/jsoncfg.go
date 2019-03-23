package tymail

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//{
//"From": "2241172930@qq.com",
//"Pass": "ltdxvumslixddjea",
//"To": [
//"373102227@qq.com"
//],
//"host": "smtp.qq.com",
//"port": 465
//}

type MailConfig struct {
	From string   `json:"Form"`
	To   []string `json:"To"`
	Pass string   `json:"pass"`
	Host string   `json:"host"`
	Port int64    `json:"port"`
}

func init() {

}

func LoadConfig(file string) (*MailConfig, error) {
	cfg := &MailConfig{}
	err := cfg.Load(file)
	return cfg, err
}

func (c *MailConfig) Load(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(data, c); err != nil {
		return err
	}

	return nil
}

func (c MailConfig) String() string {
	return fmt.Sprintf("{From:%s To:%v Pass:%s Host:%s Port:%d}", c.From, c.To, c.Pass, c.Host, c.Port)
}
