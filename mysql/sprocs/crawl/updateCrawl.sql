use harvest;

DROP PROCEDURE IF EXISTS updateCrawl;

DELIMITER $$

CREATE PROCEDURE updateCrawl(
    IN crawlIdIn INT,
    IN nameIn VARCHAR(255),
    IN configIn JSON
) BEGIN
    UPDATE Crawl SET
        name = COALESCE(nameIn, name),
        config = COALESCE(configIn, config)
    WHERE crawlId = crawlIdIn;
END $$

DELIMITER ;