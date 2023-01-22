USE harvest;

CREATE TABLE IF NOT EXISTS CrawlRun (
    crawlRunId INT AUTO_INCREMENT NOT NULL,
    crawlId INT NOT NULL,
    startTimestamp DATETIME,
    endTimestamp DATETIME,
    running BOOL,
    PRIMARY KEY (crawlRunId)
) ENGINE=InnoDB;
