# mc-payment-api

## Run/Test Project

## env sample

- APP_PORT=8080
- POSTGRES_ADDRESS=localhost
- POSTGRES_DB=local
- POSTGRES_PASSWORD=myPassword
- POSTGRES_USER=postgres
- POSTGRES_PORT=5432

# API DOCUMENTATION

## BOOK

#### GET ALL BOOKS

```http
  GET /book
```

example response

```
[
  {
    "id": 5,
    "judul": "mencari naga",
    "tahun_terbit": "2019",
    "jumlah_halaman": 201,
    "kategori_id": 1,
    "penulis_id": 1,
    "created_at": "2022-08-02T18:55:07.211447Z",
    "updated_at": "2022-08-02T23:14:08.163737Z"
  }
]
```

#### GET BOOK BY ID

```http
  GET /book/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of book to fetch |

example response

```
  {
    "id": 5,
    "judul": "mencari naga",
    "tahun_terbit": "2019",
    "jumlah_halaman": 201,
    "kategori_id": 1,
    "penulis_id": 1,
    "created_at": "2022-08-02T18:55:07.211447Z",
    "updated_at": "2022-08-02T23:14:08.163737Z"
  }
```

#### POST BOOK

```http
  POST /book
```

example request

```
{
	"judul" :"mencari naga",
	"tahun_terbit": "2019",
	"jumlah_halaman": 201,
	"kategori_id" : 1,
	"penulis_id" : 10
}
```

#### UPDATE BOOK BY ID

```http
  PUT /book/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of book to fetch |

example request

```
{
	"judul" : "hantu penasaran banget"
}
```

#### DELETE BOOK BY ID

```http
  DELETE /book/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of book to fetch |

## CATEGORY

#### GET ALL CATEGORIES

```http
  GET /category
```

example response

```
[
  {
    "id": 1,
    "nama": "horor",
    "created_at": "2022-08-02T18:52:53.12556Z",
    "updated_at": "2022-08-02T18:52:53.12556Z"
  }
]
```

#### GET CATEGORY BY ID

```http
  GET /category/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of book to fetch |

example response

```
{
  "id": 1,
  "nama": "horor",
  "created_at": "2022-08-02T18:52:53.12556Z",
  "updated_at": "2022-08-02T18:52:53.12556Z"
}
```

#### POST CATEGORY

```http
  POST /category
```

example request

```
{
"nama":"adults"
}
```

#### UPDATE CATEGORY BY ID

```http
  PUT /category/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of book to fetch |

example request

```
{
	"nama" : "fiction"
}
```

#### DELETE CATEGORY BY ID

```http
  DELETE /category/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of book to fetch |

## AUTHOR

#### GET ALL AUTHORS

```http
  GET /author
```

example response

```
[
  {
    "id": 1,
    "nama": "rahmat",
    "tanggal_lahir": "12-09-1990",
    "created_at": "2022-08-02T18:53:39.512335Z",
    "updated_at": "2022-08-02T18:53:39.512335Z"
  }
]
```

#### GET AUTHOR BY ID

```http
  GET /author/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of book to fetch |

example response

```
  {
    "id": 1,
    "nama": "rahmat",
    "tanggal_lahir": "12-09-1990",
    "created_at": "2022-08-02T18:53:39.512335Z",
    "updated_at": "2022-08-02T18:53:39.512335Z"
  }
```

#### POST AUTHOR

```http
  POST /author
```

example request

```
{
    "nama":"edo",
	"tanggal_lahir":"02-01-1990"
}   
```

#### UPDATE AUTHOR BY ID

```http
  PUT /author/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of book to fetch |

example request

```
{
	"nama" : "helmi"
}
```

#### DELETE AUTHOR BY ID

```http
  DELETE /author/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of book to fetch |