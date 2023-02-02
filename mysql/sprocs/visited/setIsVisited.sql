use harvest;

DROP PROCEDURE IF EXISTS setIsVisited;

DELIMITER $$

CREATE PROCEDURE setIsVisited(
    IN runIdIn INT,
    IN requestHashIn INT UNSIGNED
) BEGIN
    INSERT INTO Visited (runId, requestHash)
    VALUES (runIdIn, requestHashIn);
END $$

DELIMITER ;