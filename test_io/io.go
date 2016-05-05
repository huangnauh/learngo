package main

import (
	"io"
	"os"
	"fmt"
	"path/filepath"
	"path"
	"strings"
	"bytes"
	"bufio"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	fmt.Println(n)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}



func GetProjectRoot() string {
	binDir, err := executableDir()
	if err != nil {
		return ""
	}
	return path.Dir(binDir)
}

func executableDir() (string, error) {
	pathAbs, err := filepath.Abs(os.Args[0])
	if err != nil {
		return "", err
	}
	return filepath.Dir(pathAbs), nil
}

func main() {
	data, err := ReadFrom(os.Stdin, 11)
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}else{
		fmt.Printf("ok: %s\n", data)
	}


	file, err := os.Open(GetProjectRoot() + "/test_io/1.txt")
	if err != nil {
		fmt.Printf("open error\n")
	}

	data, err = ReadFrom(file, 11)
	file.Close()
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}else{
		fmt.Printf("ok: %s\n", data)
	}

	data, err = ReadFrom(strings.NewReader("string"), 6)

	if err != nil {
		fmt.Printf("err: %s\n", err)
	}else{
		fmt.Printf("ok: %s\n", data)
	}

	buffer := new(bytes.Buffer)
	err = buffer.WriteByte('1')
	if err == nil{
		ch,_ := buffer.ReadByte()
		fmt.Printf("%c\n",ch)
	}
	str := "一个"
	for _,s := range str{
		fmt.Println(string(s))
	}

	r := []rune(str)
	fmt.Printf("%s %s\n",string(r[0]), string(r[1]))

	reader := strings.NewReader("这个是测试")
	reader.Seek(-6, os.SEEK_END)
	var rc rune
	rc,_,_ = reader.ReadRune()
	fmt.Printf("now... %c\n",rc)
	p := make([]byte, 6)
	_, err = reader.ReadAt(p, 3)
	if err != nil {
	    panic(err)
	}
	fmt.Printf("%s\n",p)

	file, err = os.Create("abc.txt")
	if err != nil {
		panic(err)
	}

	file.WriteString("这是一个写操作")
	n, err := file.WriteAt([]byte("两个"), 6)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)

	file.Close()

	file, err = os.Open(GetProjectRoot() + "/test_io/1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(os.Stdout)
	writer.ReadFrom(file)
	writer.Flush()
}
