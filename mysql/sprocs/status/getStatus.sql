use harvest;

DROP PROCEDURE IF EXISTS getStatus;

DELIMITER $$

CREATE PROCEDURE getStatus(
    IN scraperIdIn INT,
    IN runIdIn INT
) BEGIN
    DECLARE runIdForStatus INT;

    IF runIdIn IS NULL THEN
        SELECT runId INTO runIdForStatus FROM Run
        WHERE scraperId = scraperIdIn 
        ORDER BY startTimestamp DESC
        LIMIT 1;
    ELSE
        SET runIdForStatus := runIdIn;
    END IF;

    SELECT
        SUM(queued) AS queued,
        SUM(successes) AS successes, 
        SUM(errors) AS errors, 
        SUM(missing) AS missing,
        MAX(lastUpdatedTimestamp) AS lastUpdatedTimestamp
    FROM Status
    WHERE runId = runIdForStatus;
END $$

DELIMITER ;