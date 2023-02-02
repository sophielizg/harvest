use harvest;

DROP PROCEDURE IF EXISTS getQueueSize;

DELIMITER $$

CREATE PROCEDURE getQueueSize(
    IN runIdIn INT
) BEGIN
    SELECT COUNT(*) AS queueSize FROM RequestQueue
    WHERE runId = runIdIn AND runnerId IS NULL;
END $$

DELIMITER ;