package test

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestIdGenerator(t *testing.T) {

	start := time.Now()
	//id, _ := GetIdByModule(MODULE_TEST)
	id := 123
	fmt.Println(id, "===耗时===>", time.Since(start))

}

func TestUrl(t *testing.T) {

	ossUrl := "/oss-storage/test/test/asdasfasf.mp4"

	urlArr := strings.Split(ossUrl, "/")

	fmt.Println(urlArr)

	bucketName := urlArr[1]

	objectName := urlArr[2]

	urlArr = urlArr[3:]

	for i := 0; i < len(urlArr); i++ {
		objectName = objectName + "/" + urlArr[i]
	}

	fmt.Println("路径" + bucketName)
	fmt.Println("路径" + objectName)

}
