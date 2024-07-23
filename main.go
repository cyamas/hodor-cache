package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/cyamas/hodor-cache/cache"
	"google.golang.org/grpc"
)

type myCacheServer struct {
	cache.UnimplementedCacheServer
	Cache *cache.Cache
}

func (cs *myCacheServer) Get(ctx context.Context, req *cache.GetRequest) (*cache.GetResponse, error) {
	key := req.Key
	log.Println("key: ", key)

	if val, ok := cs.Cache.Get(key); ok {
		switch v := val.(type) {
		case string:
			strVal := &cache.GetResponse_StrVal{StrVal: v}
			return &cache.GetResponse{Value: strVal}, nil
		case int64:
			intVal := &cache.GetResponse_IntVal{IntVal: v}
			return &cache.GetResponse{Value: intVal}, nil
		case float64:
			floatVal := &cache.GetResponse_FloatVal{FloatVal: v}
			return &cache.GetResponse{Value: floatVal}, nil
		case bool:
			boolVal := &cache.GetResponse_BoolVal{BoolVal: v}
			return &cache.GetResponse{Value: boolVal}, nil
		case *cache.StringArray:
			strArr := &cache.GetResponse_StrArr{StrArr: v}
			return &cache.GetResponse{Value: strArr}, nil
		case *cache.IntArray:
			intArr := &cache.GetResponse_IntArr{IntArr: v}
			return &cache.GetResponse{Value: intArr}, nil
		case *cache.FloatArray:
			floatArr := &cache.GetResponse_FloatArr{FloatArr: v}
			return &cache.GetResponse{Value: floatArr}, nil
		case *cache.BoolArray:
			boolArr := &cache.GetResponse_BoolArr{BoolArr: v}
			return &cache.GetResponse{Value: boolArr}, nil
		}
	}

	invalidKey := &cache.GetResponse_StrVal{StrVal: "Key does not exist"}
	return &cache.GetResponse{Value: invalidKey}, nil
}

func (cs *myCacheServer) Set(ctx context.Context, req *cache.SetRequest) (*cache.SetResponse, error) {
	key := req.Key
	var val interface{}

	switch v := req.Value.(type) {
	case *cache.SetRequest_StrVal:
		val = v.StrVal
	case *cache.SetRequest_IntVal:
		val = v.IntVal
	case *cache.SetRequest_FloatVal:
		val = v.FloatVal
	case *cache.SetRequest_BoolVal:
		val = v.BoolVal
	case *cache.SetRequest_StrArr:
		val = v.StrArr
	case *cache.SetRequest_IntArr:
		val = v.IntArr
	case *cache.SetRequest_FloatArr:
		val = v.FloatArr
	case *cache.SetRequest_BoolArr:
		val = v.BoolArr
	}

	cs.Cache.Set(key, val, time.Hour)
	cs.Cache.Display()
	return &cache.SetResponse{
		Stub: fmt.Sprintf("Set %s to %v", key, val),
	}, nil
}

func main() {
	c := cache.New()

	ln, err := net.Listen("tcp", ":3004")
	log.Println("Listening on port: 3004")
	if err != nil {
		log.Fatal("failed to listen", err)
	}
	serverRegistrar := grpc.NewServer()
	service := &myCacheServer{
		Cache: c,
	}

	cache.RegisterCacheServer(serverRegistrar, service)

	err = serverRegistrar.Serve(ln)
	if err != nil {
		log.Fatalf("could not serve")
	}
}
