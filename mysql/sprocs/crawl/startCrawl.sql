use harvest;

DROP PROCEDURE IF EXISTS startCrawl;

DELIMITER $$

CREATE PROCEDURE startCrawl(
    IN crawlIdIn INT
) BEGIN
    UPDATE Crawl SET currentRunStartTimestamp = NOW()
    WHERE crawlId = crawlIdIn AND currentRunStartTimestamp = NULL;
END $$

DELIMITER ;