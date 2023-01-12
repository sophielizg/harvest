USE harvest;

CREATE TABLE IF NOT EXISTS Request (
    requestId INT AUTO_INCREMENT NOT NULL,
    scrapeId INT NOT NULL,
    requestHash VARCHAR(16) NOT NULL,
    visitedTimestamp DATETIME,
    request BLOB NOT NULL,
    createdByRequestId INT,
    PRIMARY KEY (requestId)
) ENGINE=InnoDB;
