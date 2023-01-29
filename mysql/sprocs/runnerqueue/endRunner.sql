use harvest;

DROP PROCEDURE IF EXISTS endRunner;

DELIMITER $$

CREATE PROCEDURE endRunner(
    IN runnerIdIn INT
)
BEGIN
    UPDATE RunnerQueue SET
        endTimestamp = NOW()
    WHERE runnerId = runnerIdIn;
END $$

DELIMITER ;