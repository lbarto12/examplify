-- name: CreateCourse :exec
INSERT INTO courses (name, creator_id)
VALUES (@name, @creator_id);

-- name: GetCourses :many
SELECT name FROM courses
WHERE creator_id = @user_id
ORDER BY created_at DESC;

-- name: CourseExists :one
SELECT EXISTS(
    SELECT 1 FROM courses
    WHERE name = @name AND creator_id = @user_id
) AS exists;

-- name: GetCourseCollections :many
SELECT * FROM collections
WHERE course = @course_id
AND creator_id = @user_id;