CREATE TYPE todo_status AS ENUM ('pending', 'in_progress', 'completed');
CREATE TYPE todo_priority AS ENUM ('low', 'medium', 'high');
CREATE TABLE todos (
    id BIGSERIAl PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status todo_status NOT NULL DEFAULT 'pending',
    priority todo_priority NOT NULL DEFAULT 'medium',
    due_date TIMESTAMP DEFAULT NULL,
    user_id BIGINT NOT NULL REFERENCES users(id),
    category_id BIGINT REFERENCES categories(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
)