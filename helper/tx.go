package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {

	err := recover()

	if err != nil {
		errRollback := tx.Rollback()
		FatalIfError(errRollback)
		panic(err)
	} else {
		errCommit := tx.Commit()
		FatalIfError(errCommit)
	}
}
