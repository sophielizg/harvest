use harvest;

DROP PROCEDURE IF EXISTS getResults;

DELIMITER $$

CREATE PROCEDURE getResults(
    IN scraperIdIn INT,
    IN runIdIn INT,
    IN tagsIn VARCHAR(1024)
) BEGIN
    SELECT rs.*, GROUP_CONCAT(pt.name) AS tags FROM Result rs
    INNER JOIN Run r ON rs.runId = r.runId
    LEFT JOIN ParserTag pt ON rs.parserId = pt.parserId
    WHERE (scraperIdIn IS NULL OR r.scraperId = scraperIdIn)
    AND (runIdIn IS NULL OR rs.runId = runIdIn)
    AND (tagsIn IS NULL OR FIND_IN_SET(pt.name, tagsIn))
    GROUP BY rs.resultId
    ORDER BY rs.runId, rs.requestId, rs.elementIndex;
END $$

DELIMITER ;