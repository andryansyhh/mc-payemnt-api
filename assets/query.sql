-- number 5
select
	*
from
	buku b
join penulis p on
	b.penulis_id = p.id
where
	p.nama = 'helmi'

-- number 6
select
	*
from
	buku b
join penulis p on
	b.penulis_id = p.id
join kategori k on
    b.kategori_id = k.id
where
	k.nama = 'horor'
