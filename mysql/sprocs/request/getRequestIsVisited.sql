use harvest;

DROP PROCEDURE IF EXISTS getRequestIsVisited;

DELIMITER $$

CREATE PROCEDURE getRequestIsVisited(
    IN runIdIn INT,
    IN requestIdIn INT UNSIGNED
) BEGIN
    SELECT 1 FROM Request
    WHERE requestId = requestIdIn AND runId = runIdIn;
END $$

DELIMITER ;