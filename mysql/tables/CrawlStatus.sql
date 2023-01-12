USE harvest;

CREATE TABLE IF NOT EXISTS CrawlStatus (
    crawlStatusId INT AUTO_INCREMENT NOT NULL,
    crawlId INT NOT NULL,
    scrapeId INT,
    queued INT,
    successes INT,
    errors INT,
    missing INT,
    PRIMARY KEY (crawlStatusId)
) ENGINE=InnoDB;
