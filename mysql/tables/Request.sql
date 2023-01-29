USE harvest;

CREATE TABLE IF NOT EXISTS Request (
    requestId UNSIGNED INT NOT NULL,
    crawlRunId INT NOT NULL,
    visitedTimestamp DATETIME,
    request BLOB NOT NULL,
    createdByRequestId INT,
    PRIMARY KEY (requestId, crawlRunId)
) ENGINE=InnoDB;
