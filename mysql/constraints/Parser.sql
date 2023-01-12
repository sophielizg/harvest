use harvest;

CALL addForeignKeyIfNotExists(
    'Parser', 
    'FK_Parser_ParserType_parserTypeId',
    'typeId',
    'ParserType(parserTypeId)',
    'ON DELETE RESTRICT ON UPDATE CASCADE');

CALL addForeignKeyIfNotExists(
    'Parser', 
    'FK_Parser_Crawl_crawlId',
    'crawlId',
    'Crawl(crawlId)',
    'ON DELETE RESTRICT ON UPDATE CASCADE');

CALL addForeignKeyIfNotExists(
    'Parser', 
    'FK_Parser_Crawl_enqueueCrawlId',
    'enqueueCrawlId',
    'Crawl(crawlId)',
    'ON DELETE RESTRICT ON UPDATE CASCADE');
