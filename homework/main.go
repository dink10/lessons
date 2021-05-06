package main

import (
	"fmt"
	"strings"
	"sync"
	"unicode"
)

const LOREM = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut venenatis ipsum justo, quis commodo mi dapibus quis." +
	" Aenean scelerisque volutpat massa, sit amet dictum lorem facilisis vitae. " +
	"Sed molestie ipsum tempor bibendum viverra. Aenean vel dui lobortis, vulputate leo a, interdum justo." +
	" Maecenas eget consectetur lectus. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae;" +
	" Integer semper commodo nulla, in ullamcorper felis. Suspendisse posuere venenatis purus, ut euismod purus euismod id." +
	" Nullam ante purus, pretium eu congue a, faucibus vulputate neque. Quisque ac justo non sapien semper condimentum sit amet sed orci." +
	" Vestibulum quis leo nunc Sed nec pretium massa. Vestibulum placerat luctus augue, nec tincidunt ligula tincidunt at." +
	" Etiam tristique lectus leo, id ornare risus feugiat eu. Cras fringilla lacus odio, sed dignissim magna blandit eu." +
	" Proin vitae lobortis orci, quis porta leo. Pellentesque interdum ex quis tincidunt porttitor. Ut tempor, lectus sed tempus tempor, urna tortor fermentum augue, eu dignissim nisl ipsum id tellus." +
	" Suspendisse molestie ultrices lorem eget rutrum. Curabitur non velit sagittis, rhoncus mauris at, lacinia mi." +
	" Pellentesque ac pellentesque lacus. Suspendisse hendrerit sapien sed dictum imperdiet. Etiam tempus facilisis tellus, vel cursus nibh fermentum eu." +
	"Cras sapien leo, blandit id aliquam vitae, aliquam eu tortor. Praesent pharetra lacinia felis, eu convallis orci porta vel. " +
	"Etiam turpis eros, porta ut libero vel, condimentum rutrum dui. Praesent sed lorem vulputate, volutpat elit in, suscipit arcu."

func CountWords(lorem string, sm *sync.Map, wg *sync.WaitGroup) {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	words := strings.FieldsFunc(lorem, f)
	for _, word := range words {
		if len(word) > 1 {
			value, ok := sm.LoadOrStore(word, 1)
			if ok {
				sm.Store(word, value.(int)+1)
			}
		}
	}
	wg.Done()
}

func main() {
	n := 3
	wg := sync.WaitGroup{}
	sm := sync.Map{}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(arg string, sm *sync.Map, wg *sync.WaitGroup) {
			CountWords(arg, sm, wg)
		}(LOREM[i*(len(LOREM)/n):(i+1)*(len(LOREM)/n)], &sm, &wg)
	}

	wg.Wait()
	sm.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
