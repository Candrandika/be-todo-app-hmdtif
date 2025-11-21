# Penugasan Akhir Stacil PIT Backend HMDTIF

## Deskripsi Tugas
```
Membuat Rest API untuk mengelola to-do list. Sistem hanya memerlukan 1 tabel, yaitu tabel tasks. Sistem tidak memerlukan autentikasi, jadi hanya terdapat endpoint untuk CRUD task saja.
```

## Endpoint
| Endpoint                 | Kegunaan         |
|--------------------------|------------------|
| POST /api/v1/tasks       | Candra           |
| GET /api/v1/tasks        | Candra           |
| GET /api/v1/tasks/:id    | Arin             |
| PUT /api/v1/tasks/:id    | Arin             |
| DELETE /api/v1/tasks/:id | Arin             |

## Pembagian Tugas
| Tugas / Endpoint  | Yang Mengerjakan |
|-------------------|------------------|
| Set up awal       | Candra           |
| POST /tasks       | Candra           |
| GET /tasks        | Candra           |
| GET /tasks/:id    | Arin             |
| PUT /tasks/:id    | Arin             |
| DELETE /tasks/:id | Arin             |

## Cara Menjalankan Project

- Clone repository dengan menjalankan perintah

```shell
git clone https://github.com/Candrandika/be-todo-app-hmdtif.git
```
- Jalankan perintah untuk menjalankan container docker
```shell
docker compose up -d
```

- Endpoint dapat dicek menggunakan postman atau tools sejenis