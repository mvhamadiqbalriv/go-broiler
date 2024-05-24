package helper

import "gorm.io/gorm"

//commit or rollback gorm
func CommitOrRollback(tx *gorm.DB) {
    if r := recover(); r != nil {
        // Rollback the transaction in case of panic
        tx.Rollback()
        panic(r)
    } else if tx.Error != nil {
        // Rollback the transaction if an error occurred
        tx.Rollback()
        panic(tx.Error)
    } else {
        // Commit the transaction if no error or panic occurred
        err := tx.Commit().Error
        if err != nil {
            tx.Rollback()
            panic(err)
        }
    }
}
