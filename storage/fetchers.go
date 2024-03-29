package storage

import (
	"errors"
	"fmt"

	fetchers "github.com/Kamzs/Kamil-Ambroziak"
	"github.com/Kamzs/Kamil-Ambroziak/logger"
	"github.com/Kamzs/Kamil-Ambroziak/utils"
	sq "github.com/Masterminds/squirrel"
)

func prepareInsertFetcher(fetcher *fetchers.Fetcher) (string, []interface{}) {

	sql, args, err := sq.
		Insert("fetchers").
		Columns("url", "inter", "job_id").
		Values(fetcher.Url, fetcher.Interval, fetcher.JobID).
		ToSql()

	if err != nil {
		logger.Error("could not parse query", err)
		return "", nil
	}
	return sql, args
}

func prepareGetFetcher(fetcherId int64) (string, []interface{}) {
	//"SELECT url, inter, job_id FROM fetchers WHERE id=?;"
	users := sq.
		Select("url", "inter", "job_id").
		From("fetchers")

	active := users.Where(sq.Eq{"id": fetcherId})

	sql, args, err := active.ToSql()
	if err != nil {
		logger.Error("could not parse query", err)
		return "", nil
	}
	return sql, args
}

func (db *MySQL) SaveFetcher(fetcher *fetchers.Fetcher) utils.RestErr {
	q, p := prepareInsertFetcher(fetcher)
	if q == "" {
		return utils.NewInternalServerError("error when tying to save fetcher", errors.New("query error"))
	}
	stmt, err := db.client.Prepare(q)
	if err != nil {
		logger.Error("error when trying to prepare save fetcher statement", err)
		return utils.NewInternalServerError("error when tying to save fetcher", errors.New("database error"))
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(p...)
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

	stmt, err := db.client.Prepare(queryUpdateFetcher)
	if err != nil {
		logger.Error("error when trying to prepare update fetcher statement", err)
		return utils.NewInternalServerError("error when tying to update fetcher", errors.New("database error"))
	}
	defer stmt.Close()
	_, err = stmt.Exec(fetcher.Interval, fetcher.Url, fetcher.JobID, fetcher.Id)
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
		return utils.NewInternalServerError("error when tying to delete fetcher", errors.New("database error"))
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
		if err := rows.Scan(&fetcher.Id, &fetcher.Url, &fetcher.Interval, &fetcher.JobID); err != nil {
			logger.Error("error when scan fetcher row into fetcher struct", err)
			return nil, utils.NewInternalServerError("error when tying to get fetcher", errors.New("database error"))
		}
		results = append(results, fetcher)
	}
	if len(results) == 0 {
		return nil, utils.NewNotFoundError("no fetchers found")
	}
	return results, nil
}

func (db *MySQL) GetFetcher(fetcherId int64) (*fetchers.Fetcher, utils.RestErr) {
	q, p := prepareGetFetcher(fetcherId)
	if q == "" {
		return nil, utils.NewInternalServerError("error when tying to save fetcher", errors.New("query error"))
	}
	stmt, err := db.client.Prepare(q)
	if err != nil {
		logger.Error("error when trying to prepare get fetcher statement", err)
		return nil, utils.NewInternalServerError("error when tying to get fetcher", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(p...)
	fetcher := fetchers.Fetcher{}
	if getErr := result.Scan(&fetcher.Url, &fetcher.Interval, &fetcher.JobID); getErr != nil {
		logger.Error("error when trying to get fetcher by id", getErr)
		if fmt.Sprint(getErr) == ErrorNoRows {
			return nil, utils.NewNotFoundError(ErrorNoRows)
		}
		return nil, utils.NewInternalServerError("error when tying to get fetcher", getErr)
	}
	return &fetcher, nil
}
