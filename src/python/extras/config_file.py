import json,os
from __main__ import app


def read_properties(config_file="example.json"):
    with open(config_file, 'r') as f:
        #config = json.load(f)
        lines=f.readlines()
        return lines


@app.route("/config")
def read_config():
    config_file=os.getenv("CONFIG_FILE","example.json")
    resp=""
    for line in read_properties(config_file):
        resp+=f"""
        {line}
        <br>
        """
        print(line)
    return f"<div style='text-align: center;'> <h1 >All from json config file:<br> </h1> {resp}  <br>- Python </div>"







