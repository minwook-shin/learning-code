import logging

logger = logging.getLogger('hello_logger')
logger.setLevel(logging.DEBUG)

ch = logging.StreamHandler()
ch.setLevel(logging.DEBUG)

ch.setFormatter(
    logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s'))

logger.addHandler(ch)

logger.debug('debug text')
logger.info('information text')
logger.warning('warning text')
logger.error('error text')
logger.critical('critical text')
