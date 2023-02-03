use harvest;

DROP PROCEDURE IF EXISTS getParsersForScraper;

DELIMITER $$

CREATE PROCEDURE getParsersForScraper(
    IN scraperIdIn INT
) BEGIN
    SELECT 
        p.*,
        ty.name AS type,
        GROUP_CONCAT(tg.name) AS tags
    FROM Parser p
    INNER JOIN ParserType ty ON p.parserTypeId = ty.parserTypeId
    LEFT JOIN ParserTag tg ON p.parserId = tg.parserId
    WHERE p.scraperId = scraperIdIn
    GROUP BY p.parserId;
END $$

DELIMITER ;