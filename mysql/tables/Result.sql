USE harvest;

CREATE TABLE IF NOT EXISTS Result (
    resultId INT AUTO_INCREMENT NOT NULL,
    runId INT NOT NULL,
    requestId INT NOT NULL,
    parserId INT NOT NULL,
    elementIndex INT,
    scrapedTimestamp DATETIME,
    value TEXT,
    PRIMARY KEY (resultId)
) ENGINE=InnoDB;
