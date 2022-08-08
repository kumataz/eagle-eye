CREATE SCHEMA IF NOT EXISTS ironfish;

DROP TABLE IF EXISTS miner_ironfish;

CREATE TABLE miner_ironfish (
    machine_id        INT(11) NOT NULL,
    ip_address        VARCHAR(256) NOT NULL,
    machine_tag       TEXT(65535) NOT NULL,
    work_status       TEXT(65535) NOT NULL,
    cpu_usage         TEXT(65535) NOT NULL,
    mem_usage         TEXT(65535) NOT NULL,
    update_time       TEXT(65535) NOT NULL,
    node_graffi       TEXT(65535),
    node_height       TEXT(65535),
    node_version      TEXT(65535) NOT NULL,
    node_count        TEXT(65535),
    miner_count       TEXT(65535),
    PRIMARY KEY (machine_id)
);



