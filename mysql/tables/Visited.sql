USE harvest;

CREATE TABLE IF NOT EXISTS Visited (
    requestHash BIGINT UNSIGNED NOT NULL,
    runId INT NOT NULL,
    PRIMARY KEY (requestHash, runId)
) ENGINE=InnoDB;
