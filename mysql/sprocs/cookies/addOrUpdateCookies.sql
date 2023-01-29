use harvest;

DROP PROCEDURE IF EXISTS addOrUpdateCookies;

DELIMITER $$

CREATE PROCEDURE addOrUpdateCookies(
    IN runIdIn INT,
    IN hostIn VARCHAR(1024),
    IN valueIn TEXT
) BEGIN
    INSERT INTO Cookies
        (runId, host, value)
    VALUES
        (runIdIn, hostIn, valueIn)
    ON DUPLICATE KEY UPDATE
        value = valueIn;
END $$

DELIMITER ;