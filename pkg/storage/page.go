package storage

import (
	"os"
)

const PageSize = 4096

// CreateDB creates a new database file or truncates an existing one.
func CreateDB(path string) error {
	f,err := os.Create(path) //create a file or truncate it if it is already present
	if err != nil{
		return err
	}
	return f.Close()
}

// AllocatePage allocates a new page in the database file with owner read & write permissions,group read permissions,other read permissions.
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

// ReadPage reads a page from the database file using pageID.
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

// WritePage writes a page to the database file using pageID.
func WritePage(path string,pageID int64,data []byte) error {
	f,err := os.OpenFile(path,os.O_WRONLY,0644)
	if err != nil {
		return err
	}

	defer f.Close()

	_,err = f.WriteAt(data,pageID*PageSize)
	return err
}

