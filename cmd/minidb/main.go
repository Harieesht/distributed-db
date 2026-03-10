package main

import (
	"fmt"

	"github.com/Harieesht/distributed-db/pkg/storage"
)

func main() {

	path := "test.db"

	storage.CreateDB(path)

	pageID,_ := storage.AllocatePage(path)

	page,_ := storage.ReadPage(path,pageID)
	
	fmt.Println(len(page))
	fmt.Println(page[0])

	page[0] = 100

	storage.WritePage(path,pageID,page)

	page2,_ := storage.ReadPage(path,pageID)
	fmt.Println(page2[0])


}