USE harvest;

CREATE TABLE IF NOT EXISTS RequestQueue (
    requestQueueId INT AUTO_INCREMENT NOT NULL,
    scraperId INT NOT NULL,
    runId INT NULL,
    runnerId INT NULL,
    createdTimestamp DATETIME,
    requestBlob BLOB NOT NULL,
    isInitialRequest BOOL,
    PRIMARY KEY (requestQueueId)
) ENGINE=InnoDB;
