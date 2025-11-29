-- name: ListCategories :many
select * from categories;

-- name: GetCategory :one
select * from categories 
where id = ?;

-- name: CreateCategory :exec
insert into categories (id, name, description)
VALUES (?,?,?);

-- name: UpdateCategories :exec
update categories
set name = ?, description = ? 
where id = ?;

-- name: DeleteCategories :exec
delete from categories where id = ?;


-- name: CreateCourse :exec
insert into courses (id, name, description, price, category_id)
VALUES (?,?,?,?,?);


-- name: ListCourses :many
select c.*, ca.name as category_name from courses as c
inner join categories as ca on c.category_id = ca.id