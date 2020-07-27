package main

import (
	"context"
	"fmt"
	"time"
	"proctorNew/etcdPoc/internal/app/service/infra/db/etcd"
)

var (
	dialTimeout    = 2 * time.Second
	requestTimeout = 10 * time.Second
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), requestTimeout)

	client := etcd.NewClient()

	//_ = client.DeleteKey(ctx,"key")

	pr, _ := client.PutValue(ctx,"key","value")

	res,_ :=  client.GetValue(ctx,"key", pr)

	pr,_ = client.PutValue(ctx,"key","newValue")

	res,_ =  client.GetValue(ctx,"key",pr)
	for _,data := range res.Kvs{
		fmt.Println(string(data.Key),string(data.Value))
	}
	_,_ = client.GetRevisionValue(ctx,"key");




}
