use harvest;

DROP PROCEDURE IF EXISTS getCrawlByName;

DELIMITER $$

CREATE PROCEDURE getCrawlByName(
    IN nameIn VARCHAR(255)
) BEGIN
    SELECT * FROM Crawl WHERE name = nameIn;
END $$

DELIMITER ;