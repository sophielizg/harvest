USE harvest;

CREATE TABLE IF NOT EXISTS Result (
    resutlId INT AUTO_INCREMENT NOT NULL,
    runId INT NOT NULL,
    requestId INT NOT NULL,
    parserId INT NOT NULL,
    scrapedTimestamp DATETIME,
    value TEXT,
    PRIMARY KEY (resutlId)
) ENGINE=InnoDB;
