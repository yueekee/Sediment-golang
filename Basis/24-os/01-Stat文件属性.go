package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("my1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size())    // 13
	fmt.Println(fileInfo.Name())    // my.txt
	fmt.Println(fileInfo.IsDir())   // false
	fmt.Println(fileInfo.Mode())    // -rw-r--r-- 权限
	fmt.Println(fileInfo.ModTime()) // 2020-12-28 14:44:33.676664597 +0800 CST 创建时间
	fmt.Println(fileInfo.Sys())     // &{16777221 33188 1 12862575 501 20 0 [0 0 0 0] {1609137875 38290913} {1609137873 676664597} {1609137873 676664597} {1609137764 140087197} 13 8 4096 0 0 0 [0 0]}
	fmt.Printf("%+v", fileInfo.Sys())
	// &{Dev:16777221 Mode:33188 Nlink:1 Ino:12862575 Uid:501 Gid:20 Rdev:0 Pad_cgo_0:[0 0 0 0]
	//Atimespec:{Sec:1609137875 Nsec:38290913} Mtimespec:{Sec:1609137873 Nsec:676664597}
	//Ctimespec:{Sec:1609137873 Nsec:676664597} Birthtimespec:{Sec:1609137764 Nsec:140087197}
	//Size:13 Blocks:8 Blksize:4096 Flags:0 Gen:0 Lspare:0 Qspare:[0 0]}
}
