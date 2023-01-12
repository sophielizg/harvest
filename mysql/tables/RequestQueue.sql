USE harvest;

CREATE TABLE IF NOT EXISTS RequestQueue (
    requestQueueId INT AUTO_INCREMENT NOT NULL,
    scrapeId INT,
    createdTimestamp DATETIME,
    request BLOB NOT NULL,
    createdByRequestId INT,
    isInitialRequest BOOL,
    PRIMARY KEY (requestQueueId)
) ENGINE=InnoDB;
