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
SELECT p.id as patient_id, p.name as patient_name, p.email, p.phone, p.address, d.name as doctor_name, d.id as doctor_id FROM patients p
INNER JOIN doctors d ON p.doctor_id = d.id
ORDER BY p.created_at DESC;

-- name: DeletePatient :exec
DELETE FROM patients
WHERE id=$1;

-- name: GetPatientsByDoctor :many
SELECT p.id, p.name, p.email, p.phone, d.name as doctor_name, d.specialization, r.name as receptionist_name, p.receptionist_id FROM patients p
INNER JOIN doctors d ON p.doctor_id = d.id
INNER JOIN receptionist r ON r.id = p.receptionist_id
WHERE doctor_id=$1;
