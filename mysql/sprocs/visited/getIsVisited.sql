use harvest;

DROP PROCEDURE IF EXISTS getIsVisited;

DELIMITER $$

CREATE PROCEDURE getIsVisited(
    IN runIdIn INT,
    IN requestHashIn BIGINT UNSIGNED
) BEGIN
    SELECT 1 FROM Visited
    WHERE requestHash = requestHashIn AND runId = runIdIn;
END $$

DELIMITER ;