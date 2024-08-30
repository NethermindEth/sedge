from fastapi import FastAPI, HTTPException
import os

app = FastAPI()

@app.get("/option/{option_target}")
def get_option(option_target: str):
    # Get option from env
    option = os.getenv(option_target)
    if not option:
        raise HTTPException(status_code=404, detail="Option target not found")
    return {"option_value": option}
