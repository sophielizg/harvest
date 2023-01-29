use harvest;

DROP PROCEDURE IF EXISTS getCookiesForHost;

DELIMITER $$

CREATE PROCEDURE getCookiesForHost(
    IN runIdIn INT,
    IN hostIn VARCHAR(1024)
) BEGIN
    SELECT value FROM Cookies
    WHERE runId = runIdIn AND host = hostIn;
END $$

DELIMITER ;