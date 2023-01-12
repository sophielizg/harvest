use harvest;

DROP PROCEDURE IF EXISTS resetCrawl;

DELIMITER $$

CREATE PROCEDURE resetCrawl(
    IN crawlIdIn INT
) BEGIN
    UPDATE Crawl SET currentRunStartTimestamp = NULL 
    WHERE crawlId = crawlIdIn;
END $$

DELIMITER ;