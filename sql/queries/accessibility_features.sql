-- AddAccessibilityFeatures :exec
INSERT INTO accessibility_features (review_id, feature)
VALUES (?, ?);
-- name: DeleteAccessibilityFeature :exec
DELETE FROM accessibility_features
WHERE review_id = ?
  AND feature = ?;
-- name: FindAccessibilityFeaturesByReviewId :many
SELECT feature
FROM accessibility_features
WHERE review_id = ?;
-- name: UpdateAccessibilityFeatureByReviewIdAndFeature :exec
UPDATE accessibility_features
SET feature = ?
WHERE review_id = ?
  AND feature = ?