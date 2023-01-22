package repository

const createTodoQuery = `INSERT INTO todos(id, text, created_at, updated_at, deleted) VALUES($1, $2, $3, $4, $5)`

const updateTodoQuery = `UPDATE todos SET text = $1, updated_at = $2 WHERE id = $3`

const deleteTodoQuery = `UPDATE todos SET deleted = true WHERE id = $1`

const getTodoQuery = `SELECT id, text FROM todos WHERE id = $1 AND deleted = false`
