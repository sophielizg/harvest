USE harvest;

CREATE TABLE IF NOT EXISTS RunnerQueue (
    runnerId INT AUTO_INCREMENT NOT NULL,
    runId INT NOT NULL,
    startTimestamp DATETIME,
    endTimestamp DATETIME,
    PRIMARY KEY (runnerId)
) ENGINE=InnoDB;
