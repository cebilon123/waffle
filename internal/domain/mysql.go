package domain

import "database/sql"

type MysqlNameSystemProvider struct {
	db *sql.DB
}

func NewMysqlNameSystemProvider(db *sql.DB) *MysqlNameSystemProvider {
	return &MysqlNameSystemProvider{db: db}
}

func (m *MysqlNameSystemProvider) GetAddress(host string) (string, error) {
	//TODO implement me
	panic("implement me")
}
