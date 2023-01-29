USE harvest;

CREATE TABLE IF NOT EXISTS Cookies (
    runId INT NOT NULL,
    host VARCHAR(255) NOT NULL,
    value TEXT,
    PRIMARY KEY (runId, host)
) ENGINE=InnoDB;
