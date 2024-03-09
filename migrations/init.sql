CREATE TABLE IF NOT EXISTS projects (
    id SERIAL PRIMARY KEY,
    project_name VARCHAR(255) NULL,
    category VARCHAR(255) NOT NULL,
    project_type VARCHAR(255) NULL,
    age_category VARCHAR(255)  NULL,
    year VARCHAR(255) NULL,
    timing VARCHAR(255) NULL,
    keywords VARCHAR(255) NULL,
    summary VARCHAR(500) NULL,
    director VARCHAR(255) NULL,
    producer VARCHAR(255) NULL

);

CREATE TABLE IF NOT EXISTS admin_sessions (
    session VARCHAR(255) NOT NULL
);