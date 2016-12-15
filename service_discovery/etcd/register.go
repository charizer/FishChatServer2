package etcd

import (
	"fmt"
	etcd "github.com/coreos/etcd/clientv3"
	"github.com/golang/glog"
	"github.com/oikomi/FishChatServer2/common/xtime"
	"golang.org/x/net/context"
	"log"
	"strings"
	"time"
)

var rgClient *etcd.Client
var serviceKey string

// Register is the helper function to self-register service into Etcd/Consul server
// should call Unregister when pocess stop
// name - service name
// host - service host
// port - service port
// target - etcd dial address, for example: "http://127.0.0.1:2379;http://127.0.0.1:12379"
// interval - interval of self-register to etcd
// ttl - ttl of the register information
func Register(name string, rpcServerAddr string, target string, interval xtime.Duration, ttl xtime.Duration) (err error) {
	// get endpoints for register dial address
	endpoints := strings.Split(target, ",")
	conf := etcd.Config{
		Endpoints:   endpoints,
		DialTimeout: time.Second,
	}
	rgClient, err = etcd.New(conf)
	if err != nil {
		glog.Error(err)
		return
	}
	serviceID := fmt.Sprintf("%s-%s", name, rpcServerAddr)
	serviceKey = fmt.Sprintf("/%s/%s/%s", Prefix, name, serviceID)
	addrKey := fmt.Sprintf("/%s/%s/%s/addr", Prefix, name, serviceID)
	go func() {
		// invoke self-register with ticker
		ticker := time.NewTicker(time.Duration(interval))
		// should get first, if not exist, set it
		for {
			<-ticker.C
			_, err := rgClient.Get(context.Background(), serviceKey)
			if err != nil {
				if _, err = rgClient.Put(context.Background(), addrKey, rpcServerAddr); err != nil {
					glog.Error(err)
				}
				resp, err := rgClient.Grant(context.Background(), int64(time.Duration(ttl)/time.Second))
				if err != nil {
					glog.Error(err)
				}
				if _, err = rgClient.Put(context.Background(), serviceKey, "", etcd.WithLease(resp.ID)); err != nil {
					glog.Error(err)
				}
			} else {
				resp, err := rgClient.Grant(context.Background(), int64(time.Duration(ttl)/time.Second))
				if err != nil {
					glog.Error(err)
				}
				_, err = rgClient.Put(context.Background(), serviceKey, "", etcd.WithLease(resp.ID))
				if err != nil {
					glog.Error(err)
				}
			}
		}
	}()
	// initial register
	if _, err = rgClient.Put(context.Background(), addrKey, rpcServerAddr); err != nil {
		glog.Error(err)
		return
	}
	resp, err := rgClient.Grant(context.Background(), int64(time.Duration(ttl)/time.Second))
	if err != nil {
		glog.Error(err)
	}
	if _, err = rgClient.Put(context.Background(), serviceKey, "", etcd.WithLease(resp.ID)); err != nil {
		glog.Error(err)
		return
	}
	return
}

// Unregister delete service from etcd
func Unregister() (err error) {
	_, err = rgClient.Delete(context.Background(), serviceKey)
	if err != nil {
		glog.Error(err)
	}
	return
}
