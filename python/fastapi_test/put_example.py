from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI()


class Item(BaseModel):
    name: str
    number: float


@app.get("/")
def read_root():
    return {"Hello": "World"}


@app.get("/search/{search_id}")
def read_search(search_id: int, q: str = None):
    return {"search_id": search_id, "q": q}


@app.put("/search/{search_id}")
def create_search(search_id: int, item: Item):
    return {"item_name": item.name, "search_id": search_id}
