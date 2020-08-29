package storage

const (
	//fetcher
	queryInsertFetcher = "INSERT INTO fetchers(url, inter, job_id) VALUES(?, ?, ?);"
	queryGetFetcher    = "SELECT url, inter, job_id FROM fetchers WHERE id=?;"
	queryUpdateFetcher = "UPDATE fetchers SET inter=?, url=?, job_id=? WHERE id=?;"
	queryDeleteUser    = "DELETE FROM fetchers WHERE id=?;"
	queryFindAll = "SELECT url, inter, job_id FROM fetchers;"
	//history
	queryInsertHistoryElement = "INSERT INTO history(id, response, duration, created_at) VALUES(?, ?, ?, ?);"
	queryGetHistory    = "SELECT id, response, duration, created_at FROM fetchers WHERE id=?;"

)
