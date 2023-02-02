USE harvest;

CREATE TABLE IF NOT EXISTS Request (
    requestId INT AUTO_INCREMENT NOT NULL,
    runId INT NOT NULL,
    visitedTimestamp DATETIME,
    requestBlob BLOB NOT NULL,
    parentRequestId INT,
    originatorRequestId INT,
    PRIMARY KEY (requestId)
) ENGINE=InnoDB;
