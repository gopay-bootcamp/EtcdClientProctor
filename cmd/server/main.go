package main

import (
	"context"
	"fmt"
	"time"
	"proctorNew/internal/app/service/infra/db/etcd"
)

var (
	dialTimeout    = 2 * time.Second
	requestTimeout = 10 * time.Second
)

func main() {

	ctx,_:= context.WithTimeout(context.Background(), requestTimeout)

	client := etcd.NewClient()
	defer client.Close()


	_ = client.DeleteKey(ctx,"key")

	pr,_:=client.PutValue(ctx,"key","newvalue")
	_,_ =client.GetValueWithRevision(ctx,"key",pr)

	fmt.Println(pr.Header.Revision)

	res,_ :=  client.GetValue(ctx,"key")
	for _,data := range res.Kvs{
		fmt.Println(string(data.Key),string(data.Value))
	}
}
