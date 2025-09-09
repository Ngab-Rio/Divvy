package domain

type users struct{
	ID string `db:"id"`
	username string `db:"username"`
	email string `db:"email"`
}