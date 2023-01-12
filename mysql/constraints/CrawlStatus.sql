use harvest;

CALL addForeignKeyIfNotExists(
    'CrawlStatus', 
    'FK_CrawlStatus_Crawl_crawlId',
    'crawlId',
    'Crawl(crawlId)',
    'ON DELETE CASCADE ON UPDATE CASCADE');

CALL addForeignKeyIfNotExists(
    'CrawlStatus', 
    'FK_CrawlStatus_Scrape_scrapeId',
    'scrapeId',
    'Scrape(scrapeId)',
    'ON DELETE SET NULL ON UPDATE CASCADE');
