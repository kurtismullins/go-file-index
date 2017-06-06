package main

import "fmt"
import "path/filepath"
import "os"
//import "time"
import "encoding/json"

func main() {

	//root := "/Users/kurtismullins/Dropbox/Photos"
	root := "/Users/kurtismullins/Dropbox/Camera Uploads"

	records := []FileRecord{} // Storage for all file records

	// Walk through all files, determine their metadata, and print it.
	filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {

			// Ignore directories
			if info.IsDir() { return nil }
			
			data := GetMetaData(path, info)
			
			//fmt.Printf("%+v\n", data)
			records = append(records, data)
			return nil
		})

	//fmt.Println(records)
	jsonData, _ := json.MarshalIndent(records, "", "  ")
	jsonString := string(jsonData)
	fmt.Println(jsonString)
	
}
