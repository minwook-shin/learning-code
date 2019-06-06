import logging
import logging.config

logging.config.fileConfig('logging.conf')

logger = logging.getLogger('file config')

logger.debug('debug text')
logger.info('info text')
logger.warning('warning text')
logger.error('error text')
logger.critical('critical text')
