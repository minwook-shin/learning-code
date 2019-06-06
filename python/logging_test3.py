import logging

logging.basicConfig(
    format='%(levelname)s[%(asctime)s]:%(message)s ', level=logging.DEBUG)
logging.debug('debug text')
logging.info('information text')
logging.warning('warning text')
