package main

import(
	"fmt"
	"os"
	"kv-store/internal/store"
)

func main() {
	kvstore := store.NewInMemoryKVStore()

	err := kvstore.Load()
	if err != nil {
		fmt.Println("Error loading store:", err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: kv-store <key> <value>")
		return 
	}

	cmd := os.Args[1]

	switch cmd {
	case "set":
		if len(os.Args) < 4 {
			fmt.Println("Usage: kv-store set [key] [value]")
			return
		}
		key := os.Args[2]
		value := os.Args[3]
		err := kvstore.Set(key, []byte(value))
		if err != nil {
			fmt.Println("Error setting value:", err)
			return
		} else {
			fmt.Println("Set key", key)
		}

		err = kvstore.Persist()
		if err != nil {
			fmt.Println("Error persisting store:", err)
			return
		}

	case "get":
		if len(os.Args) < 3 {
			fmt.Println("Usage: kv-store get [key]")
			return
		}
		key := os.Args[2]
		value, err := kvstore.Get(key)
		if err != nil {
			fmt.Println("Error getting value:", err)
			return
		} else {
			fmt.Println("Value:", string(value))
		}
	
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: kv-store delete [key]")
			return
		}
		key := os.Args[2]
		err := kvstore.Delete(key)
		if err != nil {
			fmt.Println("Error deleting key:", err)
			return
		} else {
			fmt.Println("Deleted key", key)
		}

	// case "persist":
	// 	err := store.Persist()
	// 	if err != nil {
	// 		fmt.Println("Error persisting store:", err)
			
	// 	}
	
	default:
		fmt.Println("Invalid command. Available commands: set, get, delete, persist")
	}

}