CREATE TABLE
IF NOT EXISTS nodes
(
    node_id SERIAL PRIMARY KEY,  
    node_name VARCHAR
(50) NOT NULL,
    ipaddr VARCHAR
(50) NOT NULL,
    port VARCHAR
(50) NOT NULL,
    job_count INT NOT NULL DEFAULT 0
);

CREATE TABLE
IF NOT EXISTS jobs
(
    job_id SERIAL PRIMARY KEY,  
    job_name VARCHAR
(50) NOT NULL,
    state VARCHAR
(50) NOT NULL,
    owner VARCHAR
(50) NOT NULL,
    create_time TIME NOT NULL,
    last_start_time TIME NOT NULL,
    last_finsh_time TIME NOT NULL,
);