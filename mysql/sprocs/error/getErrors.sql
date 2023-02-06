use harvest;

DROP PROCEDURE IF EXISTS getErrors;

DELIMITER $$

CREATE PROCEDURE getErrors(
    IN scraperIdIn INT,
    IN runIdIn INT,
    IN tagsIn VARCHAR(1024)
) BEGIN
    SELECT e.*, GROUP_CONCAT(pt.name) AS tags FROM Result e
    INNER JOIN Run r ON e.runId = r.runId
    LEFT JOIN ParserTag pt ON e.parserId = pt.parserId
    WHERE (scraperIdIn IS NULL OR r.scraperId = scraperIdIn)
    AND (runIdIn IS NULL OR e.runId = runIdIn)
    AND (tagsIn IS NULL OR FIND_IN_SET(pt.name, tagsIn))
    GROUP BY e.errorId
    ORDER BY e.runId, e.requestId, e.elementIndex;
END $$

DELIMITER ;