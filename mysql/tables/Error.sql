USE harvest;

CREATE TABLE IF NOT EXISTS Error (
    requestId INT NOT NULL,
    parserId INT NOT NULL,
    scrapedTimestamp DATETIME,
    statusCode INT,
    isMissngParseResult BOOL,
    errorMessage VARCHAR(4096),
    response BLOB,
    PRIMARY KEY (requestId, parserId)
) ENGINE=InnoDB;
