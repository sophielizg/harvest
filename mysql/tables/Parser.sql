USE harvest;

CREATE TABLE IF NOT EXISTS Parser (
    parserId INT AUTO_INCREMENT NOT NULL,
    scraperId INT NOT NULL,
    createdTimestamp DATETIME,
    parserTypeId INT NOT NULL,
    selector VARCHAR(255),
    attr VARCHAR(255),
    xpath VARCHAR(255),
    enqueueScraperId INT,
    autoIncrementRules JSON,
    PRIMARY KEY (parserId)
) ENGINE=InnoDB;
