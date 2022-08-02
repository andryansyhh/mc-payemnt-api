DROP TABLE IF EXISTS buku;
DROP TABLE IF EXISTS kategori;
DROP TABLE IF EXISTS penulis;

CREATE TABLE buku (
	id serial4 NOT NULL,
	judul text NULL,
	tahun_terbit varchar NULL,
	jumlah_halaman int4 NULL,
	kategori_id int4 NULL,
	penulis_id int4 NULL,
	created_at timestamp NULL DEFAULT now(),
	updated_at timestamp NULL DEFAULT now(),
	deleted_at timestamp NULL,
	CONSTRAINT buku_pkey PRIMARY KEY (id)
);

CREATE TABLE kategori (
	id serial4 NOT NULL,
	nama varchar NULL,
	created_at timestamp NULL DEFAULT now(),
	updated_at timestamp NULL DEFAULT now(),
	deleted_at timestamp NULL,
	CONSTRAINT kategori_pkey PRIMARY KEY (id)
);

CREATE TABLE public.penulis (
	id serial4 NOT NULL,
	nama varchar NULL,
	tanggal_lahir varchar NULL,
	created_at timestamp NULL DEFAULT now(),
	updated_at timestamp NULL DEFAULT now(),
	deleted_at timestamp NULL,
	CONSTRAINT penulis_pkey PRIMARY KEY (id)
);