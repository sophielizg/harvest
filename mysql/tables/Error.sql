USE harvest;

CREATE TABLE IF NOT EXISTS Error (
    errorId INT AUTO_INCREMENT NOT NULL,
    runId INT NOT NULL,
    requestId INT NOT NULL,
    parserId INT NOT NULL,
    elementIndex INT,
    scrapedTimestamp DATETIME,
    statusCode INT,
    response TEXT,
    isMissngParseResult BOOL,
    errorMessage VARCHAR(4096),
    PRIMARY KEY (errorId)
) ENGINE=InnoDB;
