package etcd

import (
	//"fmt"
	//"github.com/stretchr/testify/assert"
	"context"
	"testing"
)

func TestNewClient(t *testing.T) {
	client:= NewClient();
	defer client.Close()
	if client == nil{
		t.Error("client returned nil")
	}
}

func TestEtcdClient_PutValue(t *testing.T) {
	client:= NewClient();
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err := client.PutValue(ctx, "test_key", "test_value")
	cancel()
	if err != nil{
		t.Error("Put value returned error")
	}
}

func TestEtcdClient_DeleteKey(t *testing.T) {
	client:= NewClient();
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err := client.GetValue(ctx, "test_key")

	err = client.DeleteKey(ctx, "test_key")
	cancel()
	if err != nil{
		t.Error("error in deleting key")
	}
}

func TestEtcdClient_GetValue(t *testing.T) {
	client:= NewClient();
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err := client.PutValue(ctx, "test_key2", "test_value2")
	if err != nil{
		t.Error("error in get value");
	}
	_, err = client.GetValue(ctx, "test_key2")
	cancel()
	if err!= nil{
		t.Error("error in get value")
	}
}
func TestEtcdClient_GetValueWithRevision(t *testing.T) {
	client:= NewClient();
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)

	resp1, err := client.PutValue(ctx, "test_key2", "test_value2")
	if err != nil{
		t.Error("error in put value");
	}
	resp2, err := client.PutValue(ctx, "test_key2", "test_value3")
	if err != nil{
		t.Error("error in put value");
	}



	_, err = client.GetValueWithRevision(ctx, "test_key2", resp1)
	if err != nil{
		t.Error("error in get value");
	}

	_, err = client.GetValueWithRevision(ctx, "test_key2", resp2)
	if err != nil{
		t.Error("error in get value");
	}
	cancel()
}
