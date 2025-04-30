-- name: RegisterReceptionist :exec
INSERT INTO receptionist ( name, email, phone, address, password_hash) 
VALUES ($1, $2, $3, $4, $5);

-- name: UpdateReceptionist :exec
UPDATE receptionist
SET
  name = $1,
  email = $2,
  phone = $3,
  address = $4,
  updated_at = CURRENT_TIMESTAMP
WHERE id = $5;

-- name: GetReceptionistByID :one
SELECT * FROM receptionist
WHERE id = $1;

-- name: GetReceptionistByEmail :one
SELECT * FROM receptionist
WHERE email = $1;

-- name: ListReceptionists :many
SELECT * FROM receptionist
ORDER BY created_at DESC;

-- name: DeleteReceptionist :exec
DELETE FROM receptionist
WHERE id = $1;


