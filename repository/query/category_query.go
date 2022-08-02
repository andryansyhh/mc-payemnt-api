package query

const (
	PostCategory = `
	INSERT INTO
		kategori(
			nama,
			created_at,
			updated_at		
		)
		VALUES(
			$1, now(), now()
		)
	`

	GetAllCategories = `
	SELECT
		id,
		nama,
		created_at,
		updated_at
	FROM kategori WHERE deleted_at is null
	`

	GetCategoryByID = `
	SELECT
		id,
		nama,
		created_at,
		updated_at
	FROM kategori 
	WHERE deleted_at is null 
	AND id = $1
	`

	UpdateCategoryByID = `
	UPDATE kategori SET
			nama = $2,
			updated_at	= now()	
	WHERE id = $1
	`

	DeleteCategoryByID = `
	UPDATE kategori SET deleted_at = now()
	WHERE id = $1
	`
)
