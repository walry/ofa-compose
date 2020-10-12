package _import

import (
	"bufio"
	"fmt"
	"io"
	"ofa/services/logic"
	"os"
	"path/filepath"
	"strings"
)

func ImportIndustryData(){
	path,_ := filepath.Abs(filepath.Join("import","data.txt"))
	file,err := os.Open(path)
	fmt.Println(err)
	defer file.Close()
	rd := bufio.NewReader(file)
	for {
		line,err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		l := line[:len(line) - 2]
		var d []string
		for _,item := range strings.Split(l,"\t") {
			d = append(d,item)
		}
		fmt.Println(d)
		logic.BatchSave(d)
	}
}

func init()  {
	ImportIndustryData()
}


