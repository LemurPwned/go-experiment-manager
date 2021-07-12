from pymongo.errors import _format_detailed_error
from pydantic import BaseModel
from pydantic.types import Json
from typing import Any, Optional, Dict


class Experiment(BaseModel):
    name: str
    created: str
    data: Dict[str, Any]
    assets: Optional[bytes]
