CREATE TABLE IF NOT EXISTS nodes (
    node_id SERIAL PRIMARY KEY,  
    node_name VARCHAR (50) NOT NULL,
    ipaddr VARCHAR (50) NOT NULL,
    port VARCHAR (50) NOT NULL,
    job_count INT NOT NULL DEFAULT 0
);
