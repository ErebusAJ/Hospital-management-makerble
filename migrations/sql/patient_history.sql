-- name: CreatePatientHistory :exec
INSERT INTO patient_history (
    patient_id,
    doctor_id,
    symptoms,
    diagnosis,
    prescription,
    notes,
    tests_recommended,
    follow_up_date
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: GetPatientHistoryByID :one
SELECT * FROM patient_history
WHERE id = $1;

-- name: ListPatientHistoryByPatientID :many
SELECT * FROM patient_history
WHERE patient_id = $1
ORDER BY visit_date DESC;

-- name: UpdatePatientHistory :exec
UPDATE patient_history
SET
    symptoms = $1,
    diagnosis = $2,
    prescription = $3,
    notes = $4,
    tests_recommended = $5,
    follow_up_date = $6,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $7;

-- name: DeletePatientHistory :exec
DELETE FROM patient_history
WHERE id = $1;

