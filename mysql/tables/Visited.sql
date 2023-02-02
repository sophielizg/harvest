USE harvest;

CREATE TABLE IF NOT EXISTS Visited (
    requestHash INT UNSIGNED NOT NULL,
    runId INT NOT NULL,
    PRIMARY KEY (requestHash, runId)
) ENGINE=InnoDB;
