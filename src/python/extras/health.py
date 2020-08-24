import json,os
from __main__ import app



@app.route("/health")
def health():
    print("i'm healthy")
    return "healthy"







