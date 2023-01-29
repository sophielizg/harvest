USE harvest;

CREATE TABLE IF NOT EXISTS Request (
    requestId INT UNSIGNED NOT NULL,
    runId INT NOT NULL,
    visitedTimestamp DATETIME,
    requestBlob BLOB NOT NULL,
    parentRequestId INT UNSIGNED,
    originatorRequestId INT UNSIGNED,
    PRIMARY KEY (requestId, runId)
) ENGINE=InnoDB;
