use harvest;

DROP PROCEDURE IF EXISTS unpauseCrawl;

DELIMITER $$

CREATE PROCEDURE unpauseCrawl(
    IN crawlIdIn INT
) BEGIN
    UPDATE Crawl SET
        running = 1
    WHERE crawlId = crawlIdIn;
END $$

DELIMITER ;