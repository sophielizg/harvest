use harvest;

DROP PROCEDURE IF EXISTS getRequestIsVisited;

DELIMITER $$

CREATE PROCEDURE getRequestIsVisited(
    IN crawlRunIdIn INT,
    IN requestHashIn VARCHAR(16)
) BEGIN
    SELECT 1 FROM Request r
    INNER JOIN Scrape s ON r.scrapeId = s.scrapeId
    WHERE r.requestHash = requestHashIn AND s.crawlRunId = crawlRunIdIn;
END $$

DELIMITER ;