package storage

import (
	"errors"

	fetchers "github.com/Kamzs/Kamil-Ambroziak"
	"github.com/Kamzs/Kamil-Ambroziak/logger"
	"github.com/Kamzs/Kamil-Ambroziak/utils"
)

func (db *MySQL) SaveHistoryForFetcher(historyEl *fetchers.HistoryElement) utils.RestErr {

	stmt, err := db.client.Prepare(queryInsertHistoryElement)
	if err != nil {
		logger.Error("error when trying to prepare save fetcher statement", err)
		return utils.NewInternalServerError("error when tying to save fetcher", errors.New("database error"))
	}
	defer stmt.Close()

	_, saveErr := stmt.Exec(historyEl.Id, historyEl.Response, historyEl.Duration, historyEl.CreatedAt)
	if saveErr != nil {
		logger.Error("error when trying to save fetcher", saveErr)
		return utils.NewInternalServerError("error when tying to save fetcher", errors.New("database error"))
	}
	return nil
}

func (db *MySQL) GetHistoryForFetcher(fetcherId int64) ([]fetchers.HistoryElement, utils.RestErr) {
	stmt, err := db.client.Prepare(queryGetHistory)
	if err != nil {
		logger.Error("error when trying to prepare get fetcher statement", err)
		return nil, utils.NewInternalServerError("error when tying to get history for fetcher", errors.New("database error"))
	}
	defer stmt.Close()

	rows, err := stmt.Query(fetcherId)
	if err != nil {
		logger.Error("error when trying to find users by status", err)
		return nil, utils.NewInternalServerError("error when tying to get history for fetcher", errors.New("database error"))
	}

	defer rows.Close()

	var results []fetchers.HistoryElement
	for rows.Next() {
		var historyEl fetchers.HistoryElement
		if err := rows.Scan(&historyEl.Id, &historyEl.Response, &historyEl.Duration, &historyEl.CreatedAt); err != nil {
			logger.Error("error when scan historyEl row into historyElementResponse struct", err)
			return nil, utils.NewInternalServerError("error when tying to get history for fetcher", errors.New("database error"))
		}
		results = append(results, historyEl)
	}
	if len(results) == 0 {
		return nil, utils.NewNotFoundError("no history found")
	}
	return results, nil
}
