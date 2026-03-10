package storage

import (
	"os"
)

const PageSize = 4096

func CreateDB(path string) error {
	f,err := os.Create(path) //create a file or truncate it if it is already present
	if err != nil{
		return err
	}
	return f.Close()
}

func AllocatePage(path string) (int64,error) {
	f,err := os.OpenFile(path,os.O_WRONLY|os.O_APPEND,0644)
	if err != nil {
		return 0,err
	}

	defer f.Close()

	blank := make([]byte,PageSize)
	_,err = f.Write(blank)
	if err != nil {
		return 0,err
	}

	info,err := f.Stat()
	if err != nil {
		return 0,err
	}

	return (info.Size() / PageSize) - 1,nil //return new page's id
}

func ReadPage (path string,pageID int64) ([]byte,error) {
	f,err := os.Open(path)
	if err != nil {
		return nil,err
	}

	defer f.Close()

	page := make([]byte,PageSize)
	_,err = f.ReadAt(page,pageID*PageSize)
	if err != nil {
		return nil,err
	}

	return page,nil
}

func WritePage(path string,pageID int64,data []byte) error {
	f,err := os.OpenFile(path,os.O_WRONLY,0644)
	if err != nil {
		return err
	}

	defer f.Close()

	_,err = f.WriteAt(data,pageID*PageSize)
	return err
}

