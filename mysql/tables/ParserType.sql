USE harvest;

CREATE TABLE IF NOT EXISTS ParserType (
    parserTypeId INT NOT NULL,
    name varchar(255) UNIQUE NOT NULL,
    PRIMARY KEY (parserTypeId)
) ENGINE=InnoDB;

INSERT INTO ParserType 
    (parserTypeId, name) 
VALUES
    (1, 'html'),
    (2, 'json'),
    (3, 'xml')
    AS new
ON DUPLICATE KEY UPDATE
    parserTypeId = new.parserTypeId,
    name = new.name;
