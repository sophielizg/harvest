USE harvest;

CREATE TABLE IF NOT EXISTS Crawl (
    crawlId INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(255) UNIQUE,
    createdTimestamp DATETIME,
    config JSON,
    PRIMARY KEY (crawlId)
) ENGINE=InnoDB;
