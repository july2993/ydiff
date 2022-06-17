package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func main() {
	objs := make(map[string]*unstructured.Unstructured)

	var data []byte

	handle := func() {
		err := feed(objs, data)
		if err != nil {
			fmt.Println(err)
		}
		data = data[:0]
	}

	lines := readLines(os.Stdin)
	for {
		select {
		case line, ok := <-lines:
			if !ok {
				return
			}
			// fmt.Println("debug: ", line)

			if line == "---" {
				handle()
			} else {
				data = append(data, []byte(line)...)
				data = append(data, '\n')
			}
			// Assume we get the whole yaml document if can not read more data in 1 second
		case <-time.After(time.Second):
			handle()
		}
	}

}

func readLines(r io.Reader) chan string {
	s := make(chan string, 1)

	go func() {
		defer close(s)
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			s <- scanner.Text()
		}
	}()

	return s
}

func key(obj *unstructured.Unstructured) string {
	return obj.GetNamespace() + obj.GetName()
}

func feed(objs map[string]*unstructured.Unstructured, data []byte) error {
	if len(data) == 0 {
		return nil
	}

	// fmt.Println("debug: ", string(data))

	obj := new(unstructured.Unstructured)
	err := yaml.Unmarshal(data, obj)
	if err != nil {
		return err
	}

	key := key(obj)
	pre, ok := objs[key]
	objs[key] = obj

	if !ok {
		fmt.Println("---")
		fmt.Println(string(data))
		return nil
	}

	diff := cmp.Diff(pre, obj)
	fmt.Printf("%s/%s diff:\n", obj.GetNamespace(), obj.GetName())
	fmt.Println(diff)

	return nil
}
