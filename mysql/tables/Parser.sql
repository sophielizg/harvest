USE harvest;

CREATE TABLE IF NOT EXISTS Parser (
    parserId INT AUTO_INCREMENT NOT NULL,
    crawlId INT NOT NULL,
    createdTimestamp DATETIME,
    typeId INT NOT NULL,
    selector VARCHAR(255),
    attr VARCHAR(255),
    xpath VARCHAR(255),
    jsonPath VARCHAR(255),
    enqueueCrawlId INT,
    autoIncrementRules JSON,
    PRIMARY KEY (parserId)
) ENGINE=InnoDB;
