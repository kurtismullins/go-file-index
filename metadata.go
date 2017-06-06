package main

import "os"
import "log"
import "crypto/md5"
import "io"
import "fmt"
import "path/filepath"
import "github.com/rwcarlsen/goexif/exif"
import "os/exec"
import "strings"

type FileRecord struct {
	MD5 string
	Path string
	//DateCreated os.time
	//DateModified os.time
	Size int64
	Extension string
	Name string
	Exif exif.Exif
	Mime string
}

func GetMetaData(path string, info os.FileInfo) FileRecord {

	//record := FileRecord{MD5Sum: GetSum(path), Path: path, FileSize: GetSize(info)}
	record := FileRecord{}
	record.MD5 = GetSum(path)
	record.Path = path
	record.Size = GetSize(info)
	record.Extension = GetExtension(path)
	record.Name = info.Name() //GetFileName(path)
	record.Exif = GetExifData(path)
	record.Mime = GetFileType(path)
	return record
	
}

func GetExifData(path string) exif.Exif {
	f, err := os.Open(path)
	if err != nil {
		//log.Fatal(err)
		//return ""
		return exif.Exif{}
	}
	x, err := exif.Decode(f)
	if err != nil {
		//log.Fatal(err)
		//return ""
		return exif.Exif{}
	}
	return *x
	
}

func GetFileType(path string) string {
	out, err := exec.Command("file", "--brief", "--mime-type", path).Output()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(out))
}

func GetExtension(path string) string {
	return filepath.Ext(path)
}

func GetSize(info os.FileInfo) int64 {
	return info.Size()
}

//func GetExifData(path string) 

func GetSum(path string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil)) // TODO Optimize this, if possible
}
