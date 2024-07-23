package main

import (
	"fmt"

	"github.com/cyamas/hodor-cache/client"
	"github.com/cyamas/hodor-cache/server"
)

func main() {
	greeting()
	go server.Start()
	client.Start()
}

func greeting() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("          ***** Welcome to hodor-cli! *****          ")
	fmt.Println("")
	fmt.Println("     ***** Instructions For Use *****     ")
	fmt.Println("")
	fmt.Println("*** To SET a key-value pair ***")
	fmt.Println("     Enter: set <key> <value type> <value>")
	fmt.Println("     Example:	set name str Jerry	(sets key: name to str value: Jerry)")
	fmt.Println("")
	fmt.Println("*** To GET a value by key ***")
	fmt.Println("     Enter: get <key>")
	fmt.Println("     Example:	get name	(should print Jerry)")
	fmt.Println("")
	fmt.Println("*** To DELETE a value by key ***")
	fmt.Println("     Enter: del <key>")
	fmt.Println("     Example:	del name	(should delete the key-value pair name: Jerry)")
	fmt.Println("")
	fmt.Println("*** Supported Value Types ***")
	fmt.Println("")
	fmt.Println("'str': string")
	fmt.Println("")
	fmt.Println("'int': int64")
	fmt.Println("")
	fmt.Println("'float': float64")
	fmt.Println("")
	fmt.Println("'bool': boolean")
	fmt.Println("")
	fmt.Println("'arr(str)': Array of strings")
	fmt.Println("")
	fmt.Println("'arr(int)': Array of int64s")
	fmt.Println("")
	fmt.Println("'arr(float)': Array of float64s")
	fmt.Println("")
	fmt.Println("'arr(bool)': Array of booleans")
	fmt.Println("")
	fmt.Println("*** SET an ARRAY VALUE ***")
	fmt.Println("     Enter: (<item1>, <item2>, ..., <itemN>)")
	fmt.Println("     SET Example: set firstFibos arr(int) (1, 1, 2, 3, 5, 8, 13, 21)")
	fmt.Println("")
	fmt.Println("")
}
