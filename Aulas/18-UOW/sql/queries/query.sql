-- name: CreateCategory :exec
insert into categories (id, name)
VALUES (?,?);


-- name: CreateCourse :exec
insert into courses (id, name, category_id)
VALUES (?,?,?);


