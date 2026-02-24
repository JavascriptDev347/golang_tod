INSERT INTO users (
        name,
        email,
        password,
        role,
        created_at,
        updated_at
    )
VALUES (
        'Admin',
        'admin@gmail.com',
        '$2a$10$US0HTu.YtIPFspDujUN0SODhkSVFjIst1PjfkDaKvCLYw9QUAQvre',
        'admin',
        NOW(),
        NOW()
    )