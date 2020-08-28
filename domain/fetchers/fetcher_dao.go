package fetchers

import (
	"Kamil-Ambroziak/datasources/mysql"
	"Kamil-Ambroziak/logger"
	"Kamil-Ambroziak/utils"
	"errors"
)

const (
	queryInsertUser = "INSERT INTO fetchers(url, interv) VALUES(?, ?);"
	queryGetUser    = "SELECT id, url, interv FROM fetchers WHERE id=?;"
	queryUpdateUser = "UPDATE fetchers SET url=?, interv=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM fetchers WHERE id=?;"
	//todo change to select all
	queryFindAll = "SELECT * FROM fetchers;"
)

func (fetcher *Fetcher) SaveFetcher() utils.RestErr {
	stmt, err := mysql.Client.Prepare(queryInsertUser)
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

func (fetcher *Fetcher) UpdateFetcher() utils.RestErr {
	stmt, err := mysql.Client.Prepare(queryUpdateUser)
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

func (fetcher *Fetcher) DeleteFetcher() utils.RestErr {
	stmt, err := mysql.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete fetcher statement", err)
		return utils.NewInternalServerError("error when tying to update fetcher", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err = stmt.Exec(fetcher.Id); err != nil {
		logger.Error("error when trying to delete fetcher", err)
		return utils.NewInternalServerError("error when tying to save fetcher", errors.New("database error"))
	}
	return nil
}

func (fetcher *Fetcher) FindAllFetchers() ([]Fetcher, utils.RestErr) {
	stmt, err := mysql.Client.Prepare(queryFindAll)
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

	results := make([]Fetcher, 0)
	for rows.Next() {
		var fetcher Fetcher
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
func (fetcher *Fetcher) GetHistoryForFetcher() utils.RestErr {
	stmt, err := mysql.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get fetcher statement", err)
		return utils.NewInternalServerError("error when tying to get fetcher", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(fetcher.Id)

	if getErr := result.Scan(&fetcher.Id, &fetcher.Url, &fetcher.Interval); getErr != nil {
		logger.Error("error when trying to get fetcher by id", getErr)
		return utils.NewInternalServerError("error when tying to get fetcher", errors.New("database error"))
	}
	return nil
}