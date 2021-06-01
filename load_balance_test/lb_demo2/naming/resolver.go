package naming

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"google.golang.org/grpc/resolver"
)

var schema string
var cli *clientv3.Client

type etcdResolver struct {
	rawAddr string
	schema  string
	cc      resolver.ClientConn
}

func NewResolver(etcdAddr, schema string) resolver.Builder {
	return &etcdResolver{rawAddr: etcdAddr, schema: schema}
}

func (r *etcdResolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	//fmt.Println("target:", target)
	var err error
	if cli == nil {
		cli, err = clientv3.New(clientv3.Config{
			Endpoints:   strings.Split(r.rawAddr, ";"),
			DialTimeout: 15 * time.Second,
		})
		if err != nil {
			return nil, err
		}
	}

	r.cc = cc

	go r.watch("/" + target.Scheme + "/" + target.Endpoint + "/")

	return r, nil
}

func (r etcdResolver) Scheme() string {
	return r.schema
}

func (r etcdResolver) ResolveNow(rn resolver.ResolveNowOptions) {
	//log.Println("ResolveNow")
}

func (r etcdResolver) Close() {
	//log.Println("Close")
}

func (r *etcdResolver) watch(keyPrefix string) {
	var addrList []resolver.Address

	getResp, err := cli.Get(context.Background(), keyPrefix, clientv3.WithPrefix())
	if err != nil {
		log.Println(err)
	} else {
		for i := range getResp.Kvs {

			weight, err := strconv.ParseInt(string(getResp.Kvs[i].Value), 10, 64)
			if err != nil {
				continue
			}

			addr := strings.TrimPrefix(string(getResp.Kvs[i].Key), keyPrefix)
			fmt.Println(weight, addr)
			for j := 0; j < int(weight); j++ {
				addrList = append(addrList, resolver.Address{Addr: addr, ServerName: strconv.FormatInt(int64(j), 10)})
			}
			//addrList = append(addrList, resolver.Address{Addr: strings.TrimPrefix(string(getResp.Kvs[i].Key), keyPrefix)})
		}
	}
	fmt.Println(addrList)
	// 新版本etcd去除了NewAddress方法 以UpdateState代替
	r.cc.UpdateState(resolver.State{Addresses: addrList})

	rch := cli.Watch(context.Background(), keyPrefix, clientv3.WithPrefix())
	for n := range rch {
		for _, ev := range n.Events {
			addr := strings.TrimPrefix(string(ev.Kv.Key), keyPrefix)
			weight, err := strconv.ParseInt(string(ev.Kv.Value), 10, 64)
			fmt.Println(weight)
			if err != nil {
				continue
			}
			switch ev.Type {
			case mvccpb.PUT:
				if !exist(addrList, addr, int(weight)) {
					for i := 0; i < int(weight); i++ {
						addrList = append(addrList, resolver.Address{Addr: addr})
					}
					//addrList = append(addrList, resolver.Address{Addr: addr})

					r.cc.UpdateState(resolver.State{Addresses: addrList})
				}
			case mvccpb.DELETE:
				if s, ok := remove(addrList, addr, int(weight)); ok {
					addrList = s
					r.cc.UpdateState(resolver.State{Addresses: addrList})
				}
			}
			log.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}

func exist(l []resolver.Address, addr string, weight int) bool {
	for i := range l {
		if l[i].Addr == addr {
			return true
		}
	}
	return false

	//count := 0
	//for i := range l {
	//	if l[i].Addr == addr {
	//		count += 1
	//		if count == weight {
	//			return true
	//		}
	//	}
	//}
	//return false
}

func remove(s []resolver.Address, addr string, weight int) ([]resolver.Address, bool) {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i].Addr == addr {
			count += 1
			s[i] = s[len(s)-count]
			if count == weight {
				return s[:len(s)-count], true
			}
		}
		i--
	}
	return nil, false
	//for i := range s {
	//	if s[i].Addr == addr {
	//		s[i] = s[len(s)-1]
	//		return s[:len(s)-1], true
	//	}
	//}
	//return nil, false
}
