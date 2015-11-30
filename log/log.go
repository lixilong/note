package log
import (
	"fmt"
	"os"
	"strings"
)

var (
	file *File
)
func init() {
	//file, err:= OpenFile("log.txt",0777)
	file, err := os.Create(userFile)

	if err != nil {
		fmt.Printf("Open log fil.txt failed\n")
		WriteFile("fetal.txt", "Open log.xml file failed\n", 0777) error
		return
	}
	defer file.Close()
}
func Append(logStr string, arg...interface{}) {
	item := Sprintf(logStr,arg)
	//Sprintf(format string, a ...interface{}) string
	//err := WriteFile("log.txt", item, 0777)
	//err := file.Append(item)
	file.WriteString(item)
	//if (err != nil) {
	//	fmt.Printf("append log failed in log.txt\n")
	//}
}