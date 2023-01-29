use harvest;

DROP PROCEDURE IF EXISTS getParsersForCrawl;

DELIMITER $$

CREATE PROCEDURE getParsersForCrawl(
    IN crawlIdIn INT
) BEGIN
    SELECT 
        p.*,
        ty.name AS type,
        GROUP_CONCAT(tg.name) AS tags
    FROM Parser p
    INNER JOIN ParserType ty ON p.typeId = ty.parserTypeId
    LEFT JOIN ParserTag tg ON p.parserId = tg.parserId
    WHERE p.crawlId = crawlIdIn
    GROUP BY p.parserId;
END $$

DELIMITER ;