package idgenerator

import (
	"fmt"
	"testing"
	"time"
)

func TestIdGenerator(t *testing.T) {

	start := time.Now()
	//id, _ := GetIdByModule(MODULE_TEST)
	id := 123
	fmt.Println(id, "===耗时===>", time.Since(start))

}
