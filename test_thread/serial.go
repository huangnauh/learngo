package main

import (
	"os"
	"fmt"
	"sort"
	"crypto/md5"
	"path/filepath"
	"io/ioutil"
	"sync"
	"errors"
)

func MD5ALL(root string)(map[string] [md5.Size]byte, error){
	m := make(map[string] [md5.Size]byte)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error{
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		m[path] = md5.Sum(data)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return m, nil
}

type result struct {
	path string
	sum [md5.Size]byte
	err error
}

func sumFiles(done <- chan struct{}, root string)(<-chan result, <-chan error){
	c := make(chan result)
	errc := make(chan error)

	go func() {
		var wg sync.WaitGroup
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error)error{
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			wg.Add(1)
			go func() {
				data, err := ioutil.ReadFile(path)
				select {
				case c <- result{path, md5.Sum(data), err}:
				case <- done:
				}
				wg.Done()
			}()

			select {
			case <- done:
				return errors.New("walk canceled")
			default:
				return nil
			}
		})

		go func() {
			wg.Wait()
			close(c)
		}()

		errc <- err
	}()

	return c, errc
}

func MD5ALLP(root string) (map[string] [md5.Size]byte, error){
	done := make(chan struct{})
	defer close(done)

	c, errc := sumFiles(done, root)
	m := make(map[string][md5.Size]byte)
	for r := range c {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sum
	}
	if err := <- errc; err != nil {
		return nil, err
	}
	return m, nil
}


func main() {
	m, err := MD5ALL(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	var paths [] string
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x %s\n", m[path], path)
	}
}
