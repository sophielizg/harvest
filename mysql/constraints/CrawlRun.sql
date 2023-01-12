use harvest;

CALL addForeignKeyIfNotExists(
    'CrawlRun', 
    'FK_CrawlRun_Crawl_crawlId',
    'crawlId',
    'Crawl(crawlId)',
    'ON DELETE CASCADE ON UPDATE CASCADE');

CALL addIndexIfNotExists('CrawlRun', 'IX_CrawlRun_startTimestamp', 'startTimestamp');
