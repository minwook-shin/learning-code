import logging

logging.basicConfig(filename='./file.log', filemode='w', level=logging.DEBUG)
logging.debug('write debug text')
logging.info('write information text')
logging.warning('write warning text')
