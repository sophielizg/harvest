USE harvest;

CREATE TABLE IF NOT EXISTS Result (
    runId INT NOT NULL,
    requestId INT UNSIGNED NOT NULL,
    parserId INT NOT NULL,
    scrapedTimestamp DATETIME,
    value TEXT,
    PRIMARY KEY (runId, requestId, parserId)
) ENGINE=InnoDB;
