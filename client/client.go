package client

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/cyamas/hodor-cache/cache"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Start() {
	conn, err := grpc.NewClient(":3004", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := cache.NewCacheClient(conn)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("hodor-cli >> ")
		input, _ := reader.ReadString('\n')
		tokens := tokenize(input)
		if len(tokens) == 0 {
			fmt.Println("Invalid entry. Please enter a command")
			continue
		}

		switch tokens[0] {
		case "get":
			handleGet(c, tokens)
		case "set":
			handleSet(c, tokens)
		case "del":
			handleDel(c, tokens)
		default:
			fmt.Println("Invalid entry. Hodor supports 'get', 'set', and 'del' commands")
		}
	}
}

func tokenize(input string) []string {
	input = strings.TrimSpace(input)
	tokens := strings.SplitN(input, " ", 4)

	return tokens
}

func handleGet(c cache.CacheClient, tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("Invalid GET request")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &cache.GetRequest{Key: tokens[1]}
	resp, err := c.Get(ctx, req)
	if err != nil {
		log.Println("GET error: ", err)
		return
	}

	switch v := resp.Value.(type) {
	case *cache.GetResponse_StrVal:
		fmt.Println(v.StrVal)
	case *cache.GetResponse_IntVal:
		fmt.Println(v.IntVal)
	case *cache.GetResponse_FloatVal:
		fmt.Println(v.FloatVal)
	case *cache.GetResponse_BoolVal:
		fmt.Println(v.BoolVal)
	case *cache.GetResponse_StrArr:
		fmt.Println(v.StrArr.Val)
	case *cache.GetResponse_IntArr:
		fmt.Println(v.IntArr.Val)
	case *cache.GetResponse_FloatArr:
		fmt.Println(v.FloatArr.Val)
	case *cache.GetResponse_BoolArr:
		fmt.Println(v.BoolArr.Val)
	}
}

func handleSet(c cache.CacheClient, tokens []string) {
	if len(tokens) != 4 {
		fmt.Printf("Invalid SET request. Should have %d tokens. Got %d", 4, len(tokens))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	valType := tokens[2]
	val := tokens[3]
	req := &cache.SetRequest{Key: tokens[1]}
	switch valType {
	case "str":
		req.Value = &cache.SetRequest_StrVal{StrVal: val}
	case "int":
		intVal, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("Could not parse value as int")
			return
		}
		req.Value = &cache.SetRequest_IntVal{IntVal: int64(intVal)}
	case "float":
		floatVal, err := strconv.ParseFloat(val, 64)
		if err != nil {
			fmt.Println("Could not parse value as float")
			return
		}
		req.Value = &cache.SetRequest_FloatVal{FloatVal: floatVal}
	case "bool":
		boolReq := &cache.SetRequest_BoolVal{}
		switch val {
		case "true":
			boolReq.BoolVal = true
		case "false":
			boolReq.BoolVal = false
		default:
			fmt.Println("Invalid entry. Value should be 'true' or 'false'")
			return
		}
		req.Value = boolReq
	case "arr(str)":
		if !isValidArray(val) {
			fmt.Println("Invalid entry. Array values should be surrounded by '()' and should not be empty")
			return
		}

		val = strings.Trim(val, "()")
		items := strings.Split(val, ",")
		for _, item := range items {
			item = strings.TrimSpace(item)
		}
		strArr := &cache.SetRequest_StrArr{StrArr: &cache.StringArray{Val: items}}
		req.Value = strArr
	case "arr(int)":
		if !isValidArray(val) {
			fmt.Println("Invalid entry. Array values should be surrounded by '()' and should not be empty")
			return
		}

		arr := valToIntArray(val)
		if len(arr) > 0 {
			intArr := &cache.SetRequest_IntArr{IntArr: &cache.IntArray{Val: arr}}
			req.Value = intArr
		}
	case "arr(float)":
		if !isValidArray(val) {
			fmt.Println("Invalid entry. Array values should be surrounded by '()' and should not be empty")
			return
		}

		arr := valToFloatArray(val)
		if len(arr) > 0 {
			floatArr := &cache.SetRequest_FloatArr{FloatArr: &cache.FloatArray{Val: arr}}
			req.Value = floatArr
		}
	case "arr(bool)":
		if !isValidArray(val) {
			fmt.Println("Invalid entry. Array values should be surrounded by '()' and should not be empty")
			return
		}

		arr := valToBoolArray(val)
		if len(arr) > 0 {
			boolArr := &cache.SetRequest_BoolArr{BoolArr: &cache.BoolArray{Val: arr}}
			req.Value = boolArr
		}
	default:
		fmt.Println("Invalid type. Hodor does not support this type.")
	}
	resp, err := c.Set(ctx, req)
	if err != nil {
		fmt.Println("Error Setting key-value pair: ", err)
	}
	fmt.Println(resp.Stub)
}

func isValidArray(val string) bool {
	return val[0] == '(' && val[len(val)-1] == ')' && val != "()"
}

func valToIntArray(val string) []int64 {
	val = strings.Trim(val, "()")
	items := strings.Split(val, ",")
	var arr []int64
	for _, item := range items {
		item = strings.TrimSpace(item)
		intItem, err := strconv.Atoi(item)
		if err != nil {
			log.Println("Could not parse value as int")
			return []int64{}
		}
		arr = append(arr, int64(intItem))
	}
	return arr
}

func valToFloatArray(val string) []float64 {
	val = strings.Trim(val, "()")
	items := strings.Split(val, ",")
	var arr []float64
	for _, item := range items {
		item = strings.TrimSpace(item)
		floatItem, err := strconv.ParseFloat(item, 64)
		if err != nil {
			log.Println("Could not parse value as float")
			return []float64{}
		}
		arr = append(arr, floatItem)
	}
	return arr
}

func valToBoolArray(val string) []bool {
	val = strings.Trim(val, "()")
	items := strings.Split(val, ",")
	var arr []bool
	for _, item := range items {
		item = strings.TrimSpace(item)
		boolItem, err := strconv.ParseBool(item)
		if err != nil {
			log.Println("Could not parse value as bool")
			return []bool{}
		}
		arr = append(arr, boolItem)
	}
	return arr
}

func handleDel(c cache.CacheClient, tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("Invalid DEL request")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &cache.DelRequest{Key: tokens[1]}
	resp, err := c.Del(ctx, req)
	if err != nil {
		log.Println("GET error: ", err)
		return
	}
	fmt.Println(resp.Stub)
}
