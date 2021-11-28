CREATE TABLE
IF NOT EXISTS nodes
(
    id bigserial PRIMARY KEY,  
    created_at timestamp
(0)
with time zone NOT NULL DEFAULT NOW
(),
    name text NOT NULL,
    ipaddr text NOT NULL,
    port text NOT NULL,
    jobcount integer NOT NULL,
    version integer NOT NULL DEFAULT 1
);