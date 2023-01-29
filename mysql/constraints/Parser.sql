use harvest;

CALL addForeignKeyIfNotExists(
    'Parser', 
    'FK_Parser_ParserType_parserTypeId',
    'parserTypeId',
    'ParserType(parserTypeId)',
    'ON DELETE RESTRICT ON UPDATE CASCADE');

CALL addForeignKeyIfNotExists(
    'Parser', 
    'FK_Parser_Scraper_scraperId',
    'scraperId',
    'Scraper(scraperId)',
    'ON DELETE RESTRICT ON UPDATE CASCADE');

CALL addForeignKeyIfNotExists(
    'Parser', 
    'FK_Parser_Scraper_enqueueScraperId',
    'enqueueScraperId',
    'Scraper(scraperId)',
    'ON DELETE RESTRICT ON UPDATE CASCADE');
