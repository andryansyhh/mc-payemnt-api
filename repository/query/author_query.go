package query

const (
	PostAuthor = `
	INSERT INTO
		penulis(
			nama,
			tanggal_lahir,
			created_at,
			updated_at		
		)
		VALUES(
			$1, $2, now(), now()
		)
	`

	GetAllAuthors = `
	SELECT
		id,
		nama,
		tanggal_lahir,
		created_at,
		updated_at
	FROM penulis WHERE deleted_at is null
	`

	GetAuthorByID = `
	SELECT
		id,
		nama,
		tanggal_lahir,
		created_at,
		updated_at
	FROM penulis 
	WHERE deleted_at is null 
	AND id = $1
	`

	UpdateAuthorByID = `
	UPDATE penulis SET
			nama = $2,
			tanggal_lahir = $3,
			updated_at	= now()	
	WHERE id = $1
	`

	DeleteAuthorByID = `
	UPDATE penulis SET deleted_at = now()
	WHERE id = $1
	`
)
