package models

import   (
	"data/db_mysql"
)

/*
 *上传文件的记录
 */
type Upload struct {
	Id int
	UserId int
	FileName string
	FileSize int64
	FileCert string
	FileTitle string
	CertTime int64
}

/**
 *把一条认证数据保存到数据库中
 */

func (u Upload) SaveRecord() (int64, error) {
	rs, err := db_mysql.Db.Exec("insert into data_buchong(user_id, file_name, file_size, file_cert, file_title, cert_time) " +
		"values(?,?,?,?,?) ")
	if err != nil {
		return -1, err
	}
	id, err := rs.RowsAffected()
	if err != nil {
		return -1, err
	}
	return id, nil
}

/**
 *根据用户Id查询符合条件的认证数据记录
 */
func QueryRecordsByUserId(userId int) ([]Upload, error) {
	rs, err := db_mysql.Db.Query("select id, user_id, file_name, file_size, file_cert, file_title, cert_time) where user_id = ?", userId)
	if err != nil {
		return nil, err
	}
	//从rs中读取查询到的数据，返回
	records := make([]Upload, 0) //容器
	for rs.Next() {
		var record Upload
		err := rs.Scan(&record.Id, &record.UserId, &record.FileName, &record.FileSize, &record.FileCert, &record.FileTitle, &record.CertTime)
	    if err != nil {
	    	return nil, err
		}
		records = append(records,record)
	}
	return records, nil
}