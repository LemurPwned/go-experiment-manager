from bson.objectid import ObjectId
import pymongo
from models import Experiment
from fastapi import FastAPI
from typing import Optional
from pymongo.collection import Collection

from pymongo import MongoClient


conn_str = "mongodb://localhost:27017"
client = MongoClient(conn_str)
db = client['experiments']

app = FastAPI()


@app.post("/experiments")
def post_experiment(experiment: Experiment, collection: Optional[str]):
    collection: Collection = db.get_collection(collection)
    eid = collection.insert_one(
        experiment.dict()
    )
    return {"message": f"inserted the document under: {eid}"}


@app.get("/experiments")
def get_experiments(experiment_id: Optional[str] = None,
                    name: Optional[str] = None,
                    collection: Optional[str] = "default"):
    collection = db.get_collection(collection)
    if name:
        experiment = collection.find({
            "name": name
        })
        return Experiment(**experiment)
    elif experiment_id:
        experiment = collection.find({
            "_id": ObjectId(experiment_id)
        })
        return Experiment(**experiment)
    else:
        all_experiments = collection.find({})
        return [Experiment(**exp) for exp in all_experiments]
