package models

import (
	"github.com/RosaSinaga1/go_students_crud_mysql/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

//deklarasi cetakan Student
type Student struct {
	gorm.Model
	NIM      string `gorm:""json:"nim"`
	Name     string `json:"name"`
	IPK      string `json:"ipk"`
	Jurusan  string `json:"jurusan"`
	Angkatan string `json:"angkatan"`
}

//set fungsi init()
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Student{})
}

//fungsi CreateStudent()
func (b *Student) CreateStudent() *Student {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

//set fungsi GetAllStudents()
func GetAllStudents() []Student {
	//set objek cetakan Student
	var Students []Student
	db.Find(&Students)
	return Students
}

//set fungsi GetStudentById()
func GetStudentById(nim int64) (*Student, *gorm.DB) {
	//set objek cetakan Student
	var getStudent Student
	db := db.Where("NIM=?", nim).Find(&getStudent)
	return &getStudent, db
}

//set fungsi DeleteStudent()
func DeleteStudent(nim int64) Student {
	//set objek cetakan Student
	var student Student
	db.Where("NIM=?", nim).Delete(student)
	return student
}
