package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-crud/helper"
	"go-crud/model/domain"
)

type CategoryRepositoryImpl struct {
	DB *sql.DB
}

func NewCategoryRepostory(db *sql.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		DB: db,
	}
}

func (repository CategoryRepositoryImpl) Save(ctx context.Context, category domain.Category) domain.Category {
	tx, err := repository.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "insert into category(name) value(?)"
	r, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)

	id, err := r.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category

}

func (repository CategoryRepositoryImpl) Update(ctx context.Context, category domain.Category) domain.Category {
	tx, err := repository.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "update category set name = ? where id = ?"
	_, err = tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)
	return category
}

func (repository CategoryRepositoryImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := repository.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "delete from category where id = ?"
	_, err = tx.ExecContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
}

func (repository CategoryRepositoryImpl) FindById(ctx context.Context, categoryId int) (domain.Category, error) {
	tx, err := repository.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select id, name from category where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category not found")
	}
}

func (repository CategoryRepositoryImpl) FindAll(ctx context.Context) []domain.Category {
	tx, err := repository.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select id, name from category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}
