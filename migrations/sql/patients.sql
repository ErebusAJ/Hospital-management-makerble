-- name: RegisterPatient :exec
INSERT INTO patients ( name, email, phone, address, receptionist_id, doctor_id, password_hash) 
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: UpdatePatient :exec
UPDATE patients
SET
    name=$1,
    email=$2,
    phone=$3,
    address=$4,
    doctor_id=$5,
    updated_at=CURRENT_TIMESTAMP
WHERE id=$6;  -- patient's UUID

-- name: GetPatientByID :one
SELECT * FROM patients
WHERE id=$1;

-- name: GetPatientByEmail :one
SELECT * FROM patients
WHERE email = $1;

-- name: ListPatients :many
SELECT * FROM patients
ORDER BY created_at DESC;

-- name: DeletePatient :exec
DELETE FROM patients
WHERE id=$1;

-- name: GetPatientsByDoctor
SELECT * FROM patients
INNER JOIN doctors ON patients.doctor_id = doctors.id 
WHERE doctor_id=$1  
