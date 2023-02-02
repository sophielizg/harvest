use harvest;

DROP PROCEDURE IF EXISTS getIsVisited;

DELIMITER $$

CREATE PROCEDURE getIsVisited(
    IN runIdIn INT,
    IN requestHashIn INT UNSIGNED
) BEGIN
    SELECT 1 FROM Visited
    WHERE requestHash = requestHashIn AND runId = runIdIn;
END $$

DELIMITER ;