package str_conn

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

func BufferStrConn()  {
	bufStr := bytes.Buffer{}
	t1 := time.Now()
	for i := 0; i < 3000000; i++ {
		//for i := 0; i < 10; i++ {
		bufStr.WriteString("a")
		//bufStr.String()
	}
	fmt.Println("BufferStrConn", time.Now().Sub(t1))
}

func StringBuilderConn()  {
	//t1 := time.Now()
	strBuf := strings.Builder{}
	for j := 0; j < 3000000; j++ {
		strBuf.WriteString("a")
	}
	//fmt.Println("StringBuilderConn", time.Now().Sub(t1))
}

