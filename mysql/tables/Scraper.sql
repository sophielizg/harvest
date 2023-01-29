USE harvest;

CREATE TABLE IF NOT EXISTS Scraper (
    scraperId INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(255) UNIQUE,
    createdTimestamp DATETIME,
    config JSON,
    PRIMARY KEY (scraperId)
) ENGINE=InnoDB;
