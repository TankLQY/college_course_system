package model

import (
	"ccs/db"
	"log"
)

type StudentCourse struct {
	SId int64 `json:"sid"`
	CId string `json:"cid"`
	Grade int32 `json:"grade"`
	GPA float32 `json:"gpa"`
}

func (sc *StudentCourse) Add() error {
	tx, err := db.DB.Begin()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	sql := "insert into student_course values(?,?,?,?)"
	_, err = tx.Exec(sql, sc.SId, sc.CId, sc.Grade, sc.GPA)
	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
		return err
	}

	sql = "update course set stu_num = stu_num + 1 where id=?"
	_, err = tx.Exec(sql, sc.CId)
	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}


