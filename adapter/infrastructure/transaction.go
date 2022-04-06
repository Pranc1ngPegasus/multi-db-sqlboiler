package infrastructure

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

var (
	// ErrBeginTx はトランザクションの開始に失敗したことを示す。
	ErrBeginTx = errors.New("failed to begin transaction")
	// ErrCommitTx はトランザクションのCommitに失敗したことを示す。
	ErrCommitTx = errors.New("failed to commit transaction")
)

type transactionKey struct{}

// ContextWithTransaction はコンテキストに tx を関連付けたコンテキストオブジェクトのコピーを返す。
func ContextWithTransaction(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, transactionKey{}, tx)
}

// TransactionFromContext はコンテキストに関連付けられたトランザクションインタフェースを返す。
func TransactionFromContext(ctx context.Context) *sql.Tx {
	transaction, _ := ctx.Value(transactionKey{}).(*sql.Tx)
	return transaction
}

// RunInTransaction はコンテキストに関連付けられたトランザクション内で inner 関数を実行し、
// エラーを返さないときは Commit を行います。エラーが返った場合は Rollback を行います。
// コンテキストにトランザクションがない場合はトランザクションを開始し、コンテキストへの関連付けも行います。
func RunInTransaction(ctx context.Context, dbConn *sql.DB, inner func(ctx context.Context) error) (err error) {
	// TODO: Datadog traceとの紐付け
	//       gormは https://github.com/DataDog/dd-trace-go/tree/v1/contrib/gopkg.in/jinzhu/gorm.v1 があり
	//       	dbConn = gormtrace.WithContext(ctx, dbConn)
	//       で紐づけていたが、sql.DBの場合の紐付け方がよくわからないので調べる
	tx := TransactionFromContext(ctx)
	// ctxにトランザクションがない場合はトランザクションを開始してctxに関連付ける
	if tx == nil {
		tx, err = dbConn.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			return fmt.Errorf("fail RunInTransaction.Begin original error=%s: %w", err, ErrBeginTx)
		}
		ctx = ContextWithTransaction(ctx, tx)
	}

	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			transactionErr := EndTransaction(ctx, fmt.Errorf("%v", recoverErr))
			panic(transactionErr)
		} else {
			err = EndTransaction(ctx, err)
		}
	}()

	// inner()に渡るctxには必ずトランザクションが存在する
	err = inner(ctx)
	return err
}

// EndTransaction はコンテキストに関連付けられたトランザクションを終了させます
// 第二引数のerrorがnilなら Commit を行い、エラーの値が存在する場合はRollbackを行います
func EndTransaction(ctx context.Context, err error) error {
	tx := TransactionFromContext(ctx)
	// Transactionが存在しない
	if tx == nil {
		return err
	}

	if err != nil {
		rollbackError := tx.Rollback()
		if rollbackError != nil {
			dbErr := fmt.Errorf("%s: %w", rollbackError.Error(), ErrCommitTx)
			return fmt.Errorf("%s: %w", err.Error(), dbErr)
		}
		return err
	}

	committedError := tx.Commit()
	if committedError != nil {
		return fmt.Errorf("fail EndTransaction original error=%s: %w", committedError.Error(), ErrCommitTx)
	}

	return nil
}

// IsTransactionErr はエラーがTransactionに関わるエラーであったかを判定する
func IsTransactionErr(err error) bool {
	return errors.Is(err, ErrCommitTx)
}
