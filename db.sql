tables :> 
    student:
        id uuid
        name 
        email
        created_at
        updated_at
        deleted_at
    course:
        id uuid
        title 
        price
        student_count
        created_at
        updated_at
        deleted_at

CRUD > student
CRUD > course

GET     student > courses
GET     course > students
POST    student to course