package yoha

import (
	"database/sql"
	"log"
	"strings"
)

type Session struct {
	db      *sql.DB
	sql     strings.Builder
	sqlVars []interface{}
}

type Core struct {
	db *sql.DB
}

func New(db *sql.DB) *Session {
	return &Session{db: db}
}

func (s *Session) DB() *sql.DB {
	return s.db
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
}

func Open(driver, source string) (Session, error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		return Session{}, err
	}
	if err := db.Ping(); err != nil {
		log.Println(err.Error())
		return Session{}, err
	}

	return Session{db: db}, nil
}
