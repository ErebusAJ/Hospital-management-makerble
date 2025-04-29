-- name: RegisterPatient :exec
INSERT INTO patients ( name, email, phone, address, receptionist_id, doctor_id, password_hash) 
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: UpdatePatient :exec
UPDATE patients
SET
    name=$1,
    email=$2,
    phone=$3,
    phone=$4,
    address=$5,
    updated_at=CURRENT_TIMESTAMP
WHERE id=$6;  -- patient's UUID

-- name: GetPatientDetailByID :one
SELECT * FROM patients
WHERE id=$1;

-- name: DeletePatient :exec
DELETE FROM patients
WHERE id=$1;

