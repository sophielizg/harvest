USE harvest;

CREATE TABLE IF NOT EXISTS Scrape (
    scrapeId INT AUTO_INCREMENT NOT NULL,
    crawlRunId INT NOT NULL,
    startTimestamp DATETIME,
    endTimestamp DATETIME,
    PRIMARY KEY (scrapeId)
) ENGINE=InnoDB;
