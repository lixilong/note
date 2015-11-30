package typecho

import (
	"regexp"
	"io/ioutil"
	"strings"
	"os"
)
type Note struct {
	Id uint64 
	Title string //not longer than 255
	Memo string 
	Description string 
	ExecutAt    time.Time
	Type int
	MemoType int  //text,image radio or vedio
	PathAt string   //if if not text should store in the path
	Repeat  int //weekly or monthly
	User	string  //must be user ID
}
type User struct {
	Id string  //user's name
	Nick string 
	Pass string
	Sex  int   //1 :male, 2: female
	Age  int 
	Mobile string 
	Email  string 
	Geo  string   //geo location,may useful
}

