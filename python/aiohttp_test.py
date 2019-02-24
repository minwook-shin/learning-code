import aiohttp
import asyncio
import json
from typing import Dict, Tuple

loop = asyncio.get_event_loop()

async def get(session, url: str) -> Tuple:
    async with session.get(url) as response:
        return (await response.text(), response.status)

async def getHttpData():
    async with aiohttp.ClientSession() as session:
        source: Tuple = await get(session, 'http://localhost:3000/user?id=0')
        dictSource = json.loads(source[0])
        return dictSource

result: Dict = loop.run_until_complete(getHttpData())
print(result)