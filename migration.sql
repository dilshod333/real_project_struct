CREATE TABLE student (
    id UUID PRIMARY KEY,
    name VARCHAR(64),
    email VARCHAR(64),
    created_at TIMESTAMP DEFAULT timezone('Asia/Tashkent', CURRENT_TIMESTAMP) NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT timezone('Asia/Tashkent', NULL),
    deleted_at TIMESTAMPTZ DEFAULT timezone('Asia/Tashkent', NULL)
);

CREATE TABLE course (
    id UUID PRIMARY KEY,
    title VARCHAR(64),
    price DECIMAL(9, 2),
    student_count INT,
    created_at TIMESTAMP DEFAULT timezone('Asia/Tashkent', CURRENT_TIMESTAMP) NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT timezone('Asia/Tashkent', NULL),
    deleted_at TIMESTAMPTZ DEFAULT timezone('Asia/Tashkent', NULL)
);

CREATE TABLE student_course (
    student_id UUID REFERENCES student(id),
    course_id UUID REFERENCES course(id)
);