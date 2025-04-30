-- name: RegisterDoctor :exec
INSERT INTO doctors( name, email, phone, degree, specialization, password_hash) 
VALUES ($1, $2, $3, $4, $5, $6);

-- name: UpdateDoctor :exec
UPDATE doctors
SET
  name = $1,
  email = $2,
  phone = $3,
  degree = $4,
  specialization = $5,
  updated_at = CURRENT_TIMESTAMP
WHERE id = $6;

-- name: GetDoctorByID :one
SELECT * FROM doctors
WHERE id = $1;

-- name: GetDoctorByEmail :one
SELECT * FROM doctors
WHERE email = $1;

-- name: ListDoctors :many
SELECT * FROM doctors
ORDER BY created_at DESC;

-- name: DeleteDoctor :exec
DELETE FROM doctors
WHERE id = $1;
