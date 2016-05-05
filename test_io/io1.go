package main

import (
	"io"
	"time"
	"fmt"
	"errors"
	"strings"
	"bytes"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"reflect"
	"unsafe"
)

func PipeWriter(pipeWriter *io.PipeWriter) {
	var (
		i int = 0
		err error
		n int
	)

	data := []byte("这是一个测试")
	for _, err = pipeWriter.Write(data); err == nil; n, err = pipeWriter.Write(data){
		i++
		if i == 3 {
			pipeWriter.CloseWithError(errors.New("done..."))
		}
	}
	fmt.Println("writer close 字节数：", n, " error：",  err)
}

func PipeReader(pipeReader *io.PipeReader) {
	var (
		err error
		n int
	)

	data := make([]byte, 512)
	for n, err = pipeReader.Read(data); err == nil; n, err = pipeReader.Read(data){
		fmt.Printf("%s\n", data[:n])
		time.Sleep(time.Second)
	}
	fmt.Println("read err：", err)
}

func Pipe() {
	pipeReader, pipeWriter := io.Pipe()
	go PipeReader(pipeReader)
	go PipeWriter(pipeWriter)
	time.Sleep(time.Second *5)
}

func MutiReader() {
	readers  := []io.Reader{
		strings.NewReader("from strings"),
		bytes.NewBufferString("from bytes"),
	}
	reader := io.MultiReader(readers...)
	data := make([]byte, 0, 1024)
	var (
		err error
		n int
	)

	for err != io.EOF {
		tmp := make([]byte, 10)
		n, err = reader.Read(tmp)
		if err == nil{
			data = append(data, tmp[:n]...)
		} else {
			if err != io.EOF {
				panic(err)
			}
		}
	}
	fmt.Printf("%s\n", data)
}

func Tree(dirname string, curlevel int, levelmap map[int]bool) error {
	dir_abs, err := filepath.Abs(dirname)
	if err != nil {
		return err
	}
	file_infos, err := ioutil.ReadDir(dir_abs)
	if err != nil {
		return err
	}

	file_num := len(file_infos)
	for i, file_info := range file_infos {
		for j := 1; j<curlevel; j++ {
			if levelmap[j] {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
			fmt.Print(strings.Repeat(" ", 3))
		}

		tmp_map := map[int]bool{}
		for k,v := range levelmap {
			tmp_map[k] = v
		}

		if i+1 == file_num {
			fmt.Print("`")
			delete(tmp_map, curlevel)
		} else {
			fmt.Print("|")
			tmp_map[curlevel] = true
		}
		fmt.Print("-- ")
		fmt.Println(file_info.Name())
		if file_info.IsDir() {
			Tree(filepath.Join(dir_abs,file_info.Name()), curlevel+1,
				levelmap,
				//tmp_map,
			)
		}
	}
	return nil
}

type Person struct {
	Name string
	Age int
}

func (this *Person) String() string {
	buffer := bytes.NewBufferString("this is ")
	buffer.WriteString(this.Name + ", he is ")
	buffer.WriteString(strconv.Itoa(this.Age))
	buffer.WriteString(" years old.")
	return buffer.String()
}

func (this *Person) Format(f fmt.State, c rune) {
	if c == 'L' {
		f.Write([]byte(this.String()))
		f.Write([]byte(" Person has two fields."))
	} else {
		f.Write([]byte(fmt.Sprintln(this.String())))
	}
}

func (this *Person) GoString() string{
	return "&Person{Name is " + this.Name + ", Age is " + strconv.Itoa(this.Age) +"}"
}

func main() {
	//Pipe()

	//a := bytes.NewBufferString("this is ")
	//if b, ok := a.(*bytes.Buffer); ok{
	//	fmt.Println(b.Bytes())
	//}

	var val interface{} = int64(58)
    fmt.Println(reflect.TypeOf(val))
    val = 50
    fmt.Println(reflect.TypeOf(val))

	type MyInt int
	var xx MyInt = 7
	tt := reflect.ValueOf(&xx)
	fmt.Println("type:", tt.Type())
	fmt.Println("kind:", tt.Kind())
	//fmt.Println("value:", tt.Int())
	yy, ok := tt.Interface().(MyInt)
	fmt.Println("123:", yy, ok)
	fmt.Println(tt.Interface())

	fmt.Println("settability of v:" , tt.Elem().CanSet())
	tt.Elem().SetInt(20)
	fmt.Println(tt.Elem().Interface())
	fmt.Println("change the value of x:", xx)

	checkit := func (v interface{}) (int, bool){
    	s,ok:=v.(int)
    	fmt.Printf("%v:%v\n",s,ok)
		return s, ok
	}
	a := 10
	if b, ok := checkit(a); ok{
		fmt.Println(b)
	}
	cur_map := map[int]bool{}
	Tree(".", 1, cur_map)
	p := &Person{"huang", 27}
	fmt.Println(p)
	fmt.Printf("%L\n", p)
	fmt.Printf("#v\n", p)
	//bufio.NewReaderSize()
}
