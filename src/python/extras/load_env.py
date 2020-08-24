import sys, os
from __main__ import app



@app.route("/env")
def env():
    the_variable=os.getenv("THE_VAR","default value")
    return f"<div style='text-align: center;'> <h1 >The Variable is:<br> {the_variable} </h1><br>- Python </div>"
