package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Onai23/go_students_crud_mysql/pkg/models"
	"github.com/Onai23/go_students_crud_mysql/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//set objek cetakan Student
var NewStudent models.Student

//set fungsi GetStudent()
func GetStudent(w http.ResponseWriter, r *http.Request) {
	//memanggil fungsi GetAllStudents() package models
	newStudents := models.GetAllStudents()
	//mengubah data ke dalam bentuk JSON
	res, _ := json.Marshal(newStudents)
	//set content-type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//set fungsi GetStudentById()
func GetStudentById(w http.ResponseWriter, r *http.Request) {
	//return the route variable for the current request
	vars := mux.Vars(r)
	studentId := vars["studentId"]
	NIM, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	//memanggil method GetStudentById
	studentDetails, _ := models.GetStudentById(NIM)
	//mengubah data ke dalam bentuk json
	res, _ := json.Marshal(studentDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//set fungsi CreateStudent()
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	//set objek cetakan Student dari package models
	CreateStudent := &models.Student{}
	//memanggil method ParseBody()
	utils.ParseBody(r, CreateStudent)
	//memanggil method CreateStudent()
	b := CreateStudent.CreateStudent()
	//mengubah tipe data ke dalam bentuk json
	res, _ := json.Marshal(b)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//set fungsi DeleteStudent()
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	//return route for the current request
	vars := mux.Vars(r)
	//set endpoint id
	studentId := vars["studentId"]
	NIM, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	//memanggil method DeleteStudent dari package models
	student := models.DeleteStudent(NIM)
	//mengubah tipe data ke dalam bentuk json
	res, _ := json.Marshal(student)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//set fungsi UpdateStudent()
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	//set objek cetakanStudent dengan alamat Cetakan Student
	var updateStudent = &models.Student{}
	//memanggil method ParseBOdy()
	utils.ParseBody(r, updateStudent)
	//set route dari request saat ini
	vars := mux.Vars(r)
	//set endpointe id
	studentId := vars["studentId"]
	NIM, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	//memanggil method GetStudentById()
	studentDetails, db := models.GetStudentById(NIM)
	if updateStudent.Name != "" {
		studentDetails.Name = updateStudent.Name
	}
	if updateStudent.IPK != "" {
		studentDetails.IPK = updateStudent.IPK
	}
	if updateStudent.Jurusan != "" {
		studentDetails.Jurusan = updateStudent.Jurusan
	}
	if updateStudent.Angkatan != "" {
		studentDetails.Angkatan = updateStudent.Angkatan
	}

	//simpan data yg baru
	db.Save(&studentDetails)
	//mengubah data ke dalam bentuk json
	res, _ := json.Marshal(studentDetails)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
