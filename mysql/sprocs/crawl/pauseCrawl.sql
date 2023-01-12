use harvest;

DROP PROCEDURE IF EXISTS pauseCrawl;

DELIMITER $$

CREATE PROCEDURE pauseCrawl(
    IN crawlIdIn INT
) BEGIN
    UPDATE Crawl SET
        running = 0
    WHERE crawlId = crawlIdIn;
END $$

DELIMITER ;