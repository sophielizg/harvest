use harvest;

CALL addForeignKeyIfNotExists(
    'Scrape', 
    'FK_Scrape_Crawl_crawlId',
    'crawlId',
    'Crawl(crawlId)',
    'ON DELETE CASCADE ON UPDATE CASCADE');
