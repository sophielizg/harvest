USE harvest;

CREATE TABLE IF NOT EXISTS Run (
    runId INT AUTO_INCREMENT NOT NULL,
    scraperId INT NOT NULL,
    startTimestamp DATETIME,
    endTimestamp DATETIME,
    isRunning BOOL,
    PRIMARY KEY (runId)
) ENGINE=InnoDB;
