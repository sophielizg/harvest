use harvest;

CALL addForeignKeyIfNotExists(
    'Scrape', 
    'FK_Scrape_CrawlRun_crawlRunId',
    'crawlRunId',
    'CrawlRun(crawlRunId)',
    'ON DELETE CASCADE ON UPDATE CASCADE');
