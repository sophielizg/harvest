USE harvest;

CREATE TABLE IF NOT EXISTS Error (
    runId INT NOT NULL,
    requestId INT NOT NULL,
    parserId INT NOT NULL,
    scrapedTimestamp DATETIME,
    statusCode INT,
    response BLOB,
    isMissngParseResult BOOL,
    errorMessage VARCHAR(4096),
    PRIMARY KEY (runId, requestId, parserId)
) ENGINE=InnoDB;
