package etcd

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"context"
	"testing"

	"github.com/coreos/etcd/clientv3"
)

func TestNewClient(t *testing.T) {
	db, err := clientv3.New(clientv3.Config{
		DialTimeout: dialTimeout,
		Endpoints:   []string{"localhost:2379"},
	})
	defer db.Close()
	assert.NoError(t, err)
}

func TestEtcdClient_PutValue(t *testing.T) {
	db, err := clientv3.New(clientv3.Config{
		DialTimeout: dialTimeout,
		Endpoints:   []string{"localhost:2379"},
	})
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err = db.Put(ctx, "test_key", "test_value")
	cancel()
	assert.NoError(t, err)
}

func TestEtcdClient_DeleteKey(t *testing.T) {
	db, _ := clientv3.New(clientv3.Config{
		DialTimeout: dialTimeout,
		Endpoints:   []string{"localhost:2379"},
	})
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	gresp, err := db.Get(ctx, "test_key", clientv3.WithPrefix())
	assert.NoError(t, err)
	fmt.Println("Finding keys: ", gresp.Kvs)
	resp, err := db.Delete(ctx, "test_key", clientv3.WithPrefix())
	cancel()
	assert.NoError(t, err)
	fmt.Println(resp)
}

func TestEtcdClient_GetValue(t *testing.T) {
	db, _ := clientv3.New(clientv3.Config{
		DialTimeout: dialTimeout,
		Endpoints:   []string{"localhost:2379"},
	})
	defer db.Close()

	_, err := db.Put(context.TODO(), "test_key2", "test_value2")
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := db.Get(ctx, "test_key2")
	assert.NoError(t, err)
	cancel()

	fmt.Println("Get response: ", resp.Kvs)
}

func TestEtcdClient_GetValueWithRevision(t *testing.T) {
	db, _ := clientv3.New(clientv3.Config{
		DialTimeout: dialTimeout,
		Endpoints:   []string{"localhost:2379"},
	})
	defer db.Close()
	resp1, err := db.Put(context.TODO(), "test_key2", "test_value2")
	assert.NoError(t, err)
	resp2, err := db.Put(context.TODO(), "test_key2", "test_value3")
	assert.NoError(t, err)


	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)

	resp, err := db.Get(ctx, "test_key2", clientv3.WithRev(resp1.Header.Revision))
	assert.NoError(t, err)
	fmt.Println("Get response v1: ", resp.Kvs)

	resp, err = db.Get(ctx, "test_key2", clientv3.WithRev(resp2.Header.Revision))
	assert.NoError(t, err)
	fmt.Println("Get response v2: ", resp.Kvs)

	cancel()
}
