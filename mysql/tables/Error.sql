USE harvest;

CREATE TABLE IF NOT EXISTS Error (
    requestId INT NOT NULL,
    parserId INT NOT NULL,
    scrapedTimestamp DATETIME,
    statusCode INT,
    isMissngParseResult BOOL,
    response TEXT,
    PRIMARY KEY (requestId, parserId)
) ENGINE=InnoDB;
