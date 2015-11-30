package models
import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	"time"
	"fmt"
	"encoding/json"
	"log"
)
var (
	Db orm.Ormer
	storagePath string
)
type Note struct {
	Id uint64 `orm:"column(id);pk;auto"`
	Title string `orm:"column(id);"` //not longer than 255
	Memo string `orm:"column(memo); type(longtext)"`
	Description string `orm:"column(description); type(longtext)"`
	CreateAt time.Time `orm:"column(create_at);aut_now_add;type(datetime)"`
	UpdateAt	time.Time	`orm:"column(updated_at);auto_now;type(datetime)"`
	ExecutAt    time.Time `orm: "column(execut_at); type(datetime)"`
	Type int `orm: "column(type)"`  //memo:1 or arrange:2
	MemoType int `orm: "column(memo_type)"` //text,image radio or vedio
	PathAt string `orm:"column(path_at)"`  //if if not text should store in the path
	Repeat  int `orm: "column(repeat)"` //weekly or monthly
	User	User `orm: "rel(fk)"`
}
func (n *Note) TableName() string {
	return "notes"
}
type User struct {
	Id string `orm:"column(id); pk"`
	Nick string `orm:"column(nick)"`
	Pass string `orm:"column(pass)"`
	Sex  int    `orm: "column(sex)"`
	Age  int    `orm: "column(age)"`
	Mobile string `orm: "column(mobile)"`
	Email  string  `orm: "column(email)"`
	IsUse  int     `orm: column(is_use)`
	Geo  string `orm:column(geo)`  //geo location,may useful
	NotePad []*Note `orm:"column(note_pad);reverse(many)"`
}
func (n *User) TableName() string {
	return "users"
}

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	dbUser := beego.AppConfig.String("user")
	dbPass := beego.AppConfig.String("sqlpass")
	dbHost := beego.AppConfig.String("host")
	dbName := beego.AppConfig.String("db")
	orm.RegisterDataBase("default", "mysql", dbUser + ":" + dbPass + "@" + dbHost + "/" + dbName+"?charset=utf8")
	orm.RegisterModel(new(Note), new(User))
	Db = orm.NewOrm()
}
func copyNoteToDb(dst *Note ,src *typecho.Note) {
	note.Title = pn.Title
	note.Memo = pn.Memo
	note.Description = pn.Description
	note.ExecutAt = pn.ExecutAt
	note.Type = pn.Type
	note.MemoType = pn.MemoType
	note.PathAt = pn.PathAt
	note.Repeat = pn.Repeat
	note.User = pn.User
}

func UpdateNote(pn *typecho.Note) {
	if IsUserExisted(pn.User) {
		if pn.Id != "" {
			note := Note{Id:pn.Id}
			err = Db.Read(&note)
			copyNoteToDb(&note, pn)
			if num, err := Db.Update(&note); err != nil {
				log.Append("Update Fail %s \n",pn)
				return err
			}
		} else {
			return AddNewNote(pn)
		}
	}
}

func AddNewNote(pn *typecho.Note) string {
	note := new(Note)
	CopyNoteToDb(note, pn)
	id, err := Db.Insert(note)
	if err != nil {
		fmt.Printf("Insert Fail %s \n", pn)
		log.Append("Insert Fail %s \n",pn)
	} else {
		fmt.Printf("Insert OK %s \n", pn)
	}
	return err
}
func DelNote(id uint64) {
	if num, err:= db.Delete(&Note{Id:id}); err != nil {
		fmt.Printf("Delete Fail %d \n", id)
		log.Append("Delete Fail %s \n",id)
		return 0
	}
	return num
}
func GetNotesByDay(userId string, day *time.Time) []*Note{
	// need add repeat event like weekly days 
	var notes []*Note
	num, err:= db.QueryTable("Note").Filter("ExecutAt__icontains", day).All(&Note)
	if err != nil {
		fmt.Printf("get note date period %s failed \n", day)
		log.Append("get note date period %s failed \n", day)
	}
	return Note
}
func GetNotesByPeriod(userId string, begin *time.Time, end *time.Time) []*Note{
	var notes []*Note
	num, err:= db.QueryTable("Note").Filter(ExecutAt__gt, begin).Filter(ExecutAt__lt, end).All(&Note)
	if err != nil {
		fmt.Printf("get note date period %s failed \n", begin)
		log.Append("get note date period %s failed \n", begin)
	}
	// need add repeat event like weekly days 
	return Note
}

func GetNoteById(noteId uint64) *Note {
	note := Note{Id: noteId}
	err = db.Read(&note)
	if err ！= nil {
		return &note
	} else {
		fmt.Printf("get note Id %d failed \n", noteId)
		log.Append("get note Id %d failed \n", noteId)
		return nil
	}
}