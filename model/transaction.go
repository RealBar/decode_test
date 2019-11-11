package model

func (db *DBWrapper) StartMysqlTx() error{
	tx, err := db.mysqlDB.DB().Begin()
	if err != nil{
		return err
	}
	db.lock.Lock()
	defer db.lock.Unlock()

	return nil
}
