use harvest;

DROP PROCEDURE IF EXISTS getResults;

DELIMITER $$

CREATE PROCEDURE getResults(
    IN scraperIdIn INT,
    IN runIdIn INT,
    IN tagsIn VARCHAR(1024)
) BEGIN
    SELECT rs.* FROM Result rs
    INNER JOIN Runs r ON rs.runId = r.runId
    LEFT JOIN ParserTag pt ON rs.parserId = pt.parserId
    WHERE (scraperIdIn IS NULL OR r.scraperId = scraperIdIn)
    AND (runIdIn IS NULL OR rs.runId = runIdIn)
    AND (tagsIn IS NULL OR FIND_IN_SET(pt.name, tagsIn));
END $$

DELIMITER ;