package tyip2region

import (
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
	log "github.com/thinkboy/log4go"
	"os"
)

type IpRegionServer struct {
	region *ip2region.Ip2Region
}

func init() {

}

func NewIpRegionServer() *IpRegionServer {
	return &IpRegionServer{}
}

func (irs *IpRegionServer) LoadData(db string) error {
	_, err := os.Stat(db)
	if os.IsNotExist(err) {
		log.Error(err)
		return err
	}
	irs.region, err = ip2region.New(db)

	return nil
}

func (irs *IpRegionServer) Close() {
	irs.region.Close()
}

func (irs *IpRegionServer) ConverIp(command string) string {
	ip, err := irs.region.BtreeSearch(command)
	if err != nil {
		log.Error(err)
		return ""
	}

	return ip.String()
}
