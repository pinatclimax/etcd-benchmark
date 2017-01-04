package etcd

import (
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
	"golang.org/x/net/context"
)

//Select pass the key and get value or values
func Select(key string) clientv3.GetResponse {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"10.0.1.11:2379", "10.0.1.12:2379", "10.0.1.13:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := cli.Get(ctx, key, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	cancel()

	if err != nil {
		log.Fatal(err)
	}

	return *resp

}

//Upsert represents insert or update
func Upsert(key string, value string) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"10.0.1.11:2379", "10.0.1.12:2379", "10.0.1.13:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_, err = cli.Put(ctx, key, value)
	cancel()

	if err != nil {
		log.Fatal(err)
	}
}

//Delete the key and value
func Delete(ctx context.Context, cli *clientv3.Client, key string) {

}
