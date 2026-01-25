-- name: GetCourses :many
SELECT DISTINCT course FROM collections
WHERE creator_id = @user_id;

-- name: GetCourseCollections :many
SELECT * FROM collections
WHERE course = @course_id
AND creator_id = @user_id;