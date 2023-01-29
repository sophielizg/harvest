use harvest;

DROP PROCEDURE IF EXISTS getErrors;

DELIMITER $$

CREATE PROCEDURE getErrors(
    IN scraperIdIn INT,
    IN runIdIn INT,
    IN tagsIn VARCHAR(1024)
) BEGIN
    SELECT e.* FROM Result e
    INNER JOIN Runs r ON e.runId = r.runId
    LEFT JOIN ParserTag pt ON e.parserId = pt.parserId
    WHERE (scraperIdIn IS NULL OR r.scraperId = scraperIdIn)
    AND (runIdIn IS NULL OR e.runId = runIdIn)
    AND (tagsIn IS NULL OR FIND_IN_SET(pt.name, tagsIn));
END $$

DELIMITER ;