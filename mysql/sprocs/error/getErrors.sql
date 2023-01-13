use harvest;

DROP PROCEDURE IF EXISTS getErrors;

DELIMITER $$

CREATE PROCEDURE getErrors(
    IN crawlIdIn INT,
    IN crawlRunIdIn INT,
    IN tagsIn VARCHAR(1024)
) BEGIN
    SELECT e.* FROM Result e
    INNER JOIN Request rq ON e.requestId = rq.requestId
    INNER JOIN Scrape s ON rq.scrapeId = s.scrapeId
    LEFT JOIN ParserTag pt ON e.parserId = pt.parserId
    WHERE (crawlIdIn IS NULL OR s.crawlId = crawlIdIn)
    AND (crawlRunIdIn IS NULL OR s.crawlRunId = crawlRunIdIn)
    AND (tagsIn IS NULL OR FIND_IN_SET(pt.name, tagsIn));
END $$

DELIMITER ;