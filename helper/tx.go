package helper

import "database/sql"

func CommitOrRollbakc(tx *sql.Tx) {
	err := recover()
	defer func() {
		if err != nil {
			errorRolback := tx.Rollback()
			PanicIfError(errorRolback)
		} else {
			errorCommit := tx.Commit()
			PanicIfError(errorCommit)
		}
	}()
}
