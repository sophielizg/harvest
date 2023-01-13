use harvest;

DROP PROCEDURE IF EXISTS getCrawlRuns;

DELIMITER $$

CREATE PROCEDURE getCrawlRuns(
    IN crawlIdIn INT
) BEGIN
    SELECT * FROM CrawlRun
    WHERE crawlId = crawlIdIn;
END $$

DELIMITER ;