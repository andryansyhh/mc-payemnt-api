package query

const (
	PostBook = `
	INSERT INTO
		buku(
			judul,
			tahun_terbit,
			jumlah_halaman,
			kategori_id,
			penulis_id,
			created_at,
			updated_at		
		)
		VALUES(
			$1, $2, $3, $4, $5, now(), now()
		)
	`

	GetAllBooks = `
	SELECT
		id,
		judul,
		tahun_terbit,
		jumlah_halaman,
		kategori_id,
		penulis_id,
		created_at,
		updated_at
	FROM buku WHERE deleted_at is null
	`

	GetBookByID = `
	SELECT
		id,
		judul,
		tahun_terbit,
		jumlah_halaman,
		kategori_id,
		penulis_id,
		created_at,
		updated_at
	FROM buku 
	WHERE deleted_at is null 
	AND id = $1
	`

	UpdateBookByID = `
	UPDATE buku SET
			judul = $2,
			tahun_terbit = $3,
			jumlah_halaman = $4,
			kategori_id = $5,
			penulis_id = $6,
			updated_at	= now()	
	WHERE id = $1
	`

	DeleteBookByID = `
	UPDATE buku SET deleted_at = now()
	WHERE id = $1
	`
)
