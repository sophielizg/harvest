USE harvest;

CREATE TABLE IF NOT EXISTS Status (
    statusId INT AUTO_INCREMENT NOT NULL,
    runId INT NULL,
    runnerId INT NULL,
    lastUpdatedTimestamp DATETIME,
    queued INT,
    successes INT,
    errors INT,
    missing INT,
    PRIMARY KEY (statusId)
) ENGINE=InnoDB;
