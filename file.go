package lhktools

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//t = 0 以字节为单位来读取，相对底层一些，代码量也较大
//t =1  ReadAll接收一个io.Reader的参数 返回字节切片
//t =2  ioutil.ReadFile(文件名字)
func GetFileContent(fileType int,file_path string) (string,error) {
	file, err := os.Open(file_path)
	if err!=nil{
		return "", err
	}
	defer file.Close()

	if fileType ==1{
		// ReadAll接收一个io.Reader的参数 返回字节切片
		bytes, err := ioutil.ReadAll(file)
		if err!=nil {
			return "", err
		}
		return  string(bytes),nil
	}

	if fileType == 2{
		bytes, err := ioutil.ReadFile(file_path)
		if err!=nil {
			return "", err
		}
		return  string(bytes),nil
	}

	// 字节切片缓存 存放每次读取的字节
	buf := make([]byte, 1024)

	// 该字节切片用于存放文件所有字节
	var bytes []byte

	for {
		// 返回本次读取的字节数
		count, err := file.Read(buf)


		// 检测是否到了文件末尾
		if err == io.EOF {
			break
		}
		// 取出本次读取的数据
		currBytes := buf[:count]

		// 将读取到的数据 追加到字节切片中
		bytes = append(bytes, currBytes...)
	}
	// 将字节切片转为字符串 最后打印出来文件内容
	return string(bytes),nil

}

func WriteByteTofile(datas []byte,filename string) error {
	return  writeFile(datas,"",filename)
}

func WriteStrTofile(data string,filename string) error {
     return writeFile(nil,data,filename)
}
func writeFile(databyte []byte,dataStr,filename string) error {
	file, err := os.OpenFile(filename, os.O_RDWR | os.O_APPEND | os.O_CREATE, 0777)
	defer file.Close()
	if err!=nil{
		return err
	}
	if len(databyte) !=0 {
		// 写入字节
		_,err=file.Write(databyte)
	}
	if len(dataStr)!=0 {
		// 写入字符串
		_,err=file.WriteString(dataStr)
	}
	if err!=nil{
		return err
	}
	// 确保写入到磁盘
	file.Sync()
	return nil
}

func xx(){
	folderPath:="demoxxxx/ttt6"
	if !FileOrDirIsExist(folderPath) {
		// 必须分成两步：先创建文件夹、再修改权限
		//os.Mkdir(folderPath, 0777) //0777也可以os.ModePerm
		os.MkdirAll(folderPath,os.ModePerm)
		os.Chmod(folderPath, 0777)
	}else{
		fmt.Print("==isEXIT")
	}
	str:="1.txt"
	f, err1 := os.Create(str) //创建文件
	os.Chmod(str, 0777)
	fmt.Println(f,err1)
}

//判断文件夹或者文件是否存在
func FileOrDirIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}