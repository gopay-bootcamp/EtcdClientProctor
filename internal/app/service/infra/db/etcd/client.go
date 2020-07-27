package etcd

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

type EtcdClient interface {
	DeleteKey(ctx context.Context, key string) error
	PutValue(ctx context.Context, key string, value string) (*clientv3.PutResponse, error)
	GetValue(ctx context.Context, key string) (*clientv3.GetResponse, error)
	GetValueWithRevision(ctx context.Context, key string, pr *clientv3.PutResponse) (*clientv3.GetResponse, error)
	Close()
	SetWatchOnKey(ctx context.Context, key string) clientv3.WatchChan
}

type etcdClient struct {
	db *clientv3.Client
}

var (
	dialTimeout    = 2 * time.Second
	requestTimeout = 10 * time.Second
)

func NewClient() EtcdClient {

	db, _ := clientv3.New(clientv3.Config{
		DialTimeout: dialTimeout,
		Endpoints:   []string{"localhost:2379"},
	})
	return &etcdClient{
		db:db,
	}
}

func (client *etcdClient) DeleteKey(ctx context.Context, key string) error {
	_, err := client.db.Delete(ctx, key)
	if err != nil {
		return err
	}
	return nil
}

func (client *etcdClient) PutValue(ctx context.Context, key string, value string) (*clientv3.PutResponse, error) {
	pr, err := client.db.Put(ctx, key, value)
	if err != nil {
		return pr.OpResponse().Put(), err
	}
	return pr.OpResponse().Put(), nil
}

func (client *etcdClient) GetValue(ctx context.Context, key string) (*clientv3.GetResponse, error) {
	res, err := client.db.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	return res.OpResponse().Get(), nil
}

func (client *etcdClient) GetValueWithRevision(ctx context.Context, key string, pr *clientv3.PutResponse) (*clientv3.GetResponse, error) {
	res, err := client.db.Get(ctx, key, clientv3.WithLastRev()...)
	if err != nil {
		return nil, err
	}
	return res.OpResponse().Get(), nil
}

func (client *etcdClient) SetWatchOnKey(ctx context.Context, key string) clientv3.WatchChan {
	watchChan := client.db.Watch(ctx, key)
	fmt.Println("set WATCH on " + key)
	return watchChan

}

func (client *etcdClient) Close() {
	fmt.Println("Closing connections to db")
	defer client.db.Close()
}
