package storage

import (
	fetchers "Kamil-Ambroziak"
	"Kamil-Ambroziak/logger"
	"Kamil-Ambroziak/utils"

	"errors"

)

const (
	queryInsertFetcher = "INSERT INTO fetchers(url, interv) VALUES(?, ?);"
	queryGetFetcher    = "SELECT id, url, interv FROM fetchers WHERE id=?;"
	queryGetHistory    = ""

	queryUpdateUser    = "UPDATE fetchers SET url=?, interv=? WHERE id=?;"
	queryDeleteUser    = "DELETE FROM fetchers WHERE id=?;"
	//todo change to select all
	queryFindAll = "SELECT * FROM fetchers;"
)


func (db *MySQL) SaveFetcher(fetcher *fetchers.Fetcher) utils.RestErr {
	stmt, err := db.client.Prepare(queryInsertFetcher)
	if err != nil {
		logger.Error("error when trying to prepare save fetcher statement", err)
		return utils.NewInternalServerError("error when tying to save fetcher", errors.New("database error"))
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(fetcher.Url, fetcher.Interval)
	if saveErr != nil {
		logger.Error("error when trying to save fetcher", saveErr)
		return utils.NewInternalServerError("error when tying to save fetcher", errors.New("database error"))
	}

	fetcherId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new fetcher", err)
		return utils.NewInternalServerError("error when tying to save fetcher", errors.New("database error"))
	}
	fetcher.Id = fetcherId

	return nil
}
func (db *MySQL) UpdateFetcher(fetcher *fetchers.Fetcher) utils.RestErr {
	stmt, err := db.client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update fetcher statement", err)
		return utils.NewInternalServerError("error when tying to update fetcher", errors.New("database error"))
	}
	defer stmt.Close()

	_, err = stmt.Exec(fetcher.Url, fetcher.Interval, fetcher.Id)
	if err != nil {
		logger.Error("error when trying to update fetcher", err)
		return utils.NewInternalServerError("error when tying to update fetcher", errors.New("database error"))
	}
	return nil
}
func (db *MySQL) DeleteFetcher(fetcherId int64) utils.RestErr {
	stmt, err := db.client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete fetcher statement", err)
		return utils.NewInternalServerError("error when tying to update fetcher", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err = stmt.Exec(fetcherId); err != nil {
		logger.Error("error when trying to delete fetcher", err)
		return utils.NewInternalServerError("error when tying to save fetcher", errors.New("database error"))
	}
	return nil
}

func (db *MySQL) FindAllFetchers() ([]fetchers.Fetcher, utils.RestErr) {
	stmt, err := db.client.Prepare(queryFindAll)
	if err != nil {
		logger.Error("error when trying to prepare find users by status statement", err)
		return nil, utils.NewInternalServerError("error when tying to get fetcher", errors.New("database error"))
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		logger.Error("error when trying to find users by status", err)
		return nil, utils.NewInternalServerError("error when tying to get fetcher", errors.New("database error"))
	}
	defer rows.Close()

	results := make([]fetchers.Fetcher, 0)
	for rows.Next() {
		var fetcher fetchers.Fetcher
		if err := rows.Scan(&fetcher.Id, &fetcher.Url, &fetcher.Interval); err != nil {
			logger.Error("error when scan fetcher row into fetcher struct", err)
			return nil, utils.NewInternalServerError("error when tying to gett fetcher", errors.New("database error"))
		}
		results = append(results, fetcher)
	}
	if len(results) == 0 {
		return nil, utils.NewNotFoundError("no fetchers found")
	}
	return results, nil
}
func (db *MySQL) GetHistoryForFetcher(fetcherId int64 ) (*fetchers.Fetcher, utils.RestErr) {
	stmt, err := db.client.Prepare(queryGetFetcher)
	if err != nil {
		logger.Error("error when trying to prepare get fetcher statement", err)
		return nil, utils.NewInternalServerError("error when tying to get fetcher", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(fetcherId)
	fetcher:= fetchers.Fetcher{}
	if getErr := result.Scan(&fetcher.Id, &fetcher.Url, &fetcher.Interval); getErr != nil {
		logger.Error("error when trying to get fetcher by id", getErr)
		return nil, utils.NewInternalServerError("error when tying to get fetcher", errors.New("database error"))
	}
	return &fetcher,nil
}