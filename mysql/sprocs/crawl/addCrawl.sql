use harvest;

DROP PROCEDURE IF EXISTS addCrawl;

DELIMITER $$

CREATE PROCEDURE addCrawl(
    IN nameIn VARCHAR(255),
    IN configIn JSON
) BEGIN
    INSERT INTO Crawl
        (name, createdTimestamp, config)
    VALUES
        (nameIn, NOW(), configIn);
    SELECT LAST_INSERT_ID() AS crawlId;
END $$

DELIMITER ;