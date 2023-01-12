use harvest;

DROP PROCEDURE IF EXISTS getCrawlByCrawlId;

DELIMITER $$

CREATE PROCEDURE getCrawlByCrawlId(
    IN crawlIdIn INT
) BEGIN
    SELECT * FROM Crawl WHERE crawlId = crawlIdIn;
END $$

DELIMITER ;