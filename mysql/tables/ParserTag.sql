USE harvest;

CREATE TABLE IF NOT EXISTS ParserTag (
    tagId INT AUTO_INCREMENT NOT NULL,
    parserId INT NOT NULL,
    createdTimestamp DATETIME,
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (tagId)
) ENGINE=InnoDB;
