USE harvest;

CREATE TABLE IF NOT EXISTS Result (
    requestId INT NOT NULL,
    parserId INT NOT NULL,
    scrapedTimestamp DATETIME,
    value TEXT,
    PRIMARY KEY (requestId, parserId)
) ENGINE=InnoDB;
