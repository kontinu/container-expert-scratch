import sys, os
from __main__ import app
from __main__ import kontinu_msg

def get_next_msg():
    import requests
    next_msg=""
    if os.getenv("NEXT_URL"):
        r=requests.get(os.getenv("NEXT_URL"))
        next_msg=r.json()["msg"]
    return " "+next_msg + " "

@app.route('/api')
def api():
    from flask import jsonify
    global kontinu_msg
    final_msg=kontinu_msg+get_next_msg()
    print(f"{final_msg}")
    return jsonify({"msg":f"{final_msg}", "hostname": f"{os.getenv('HOSTNAME')}"})
    #return """<div style='text-align: center;'><h1>{}</h1><br><h3> - Python</h3></div>""".format(final_msg)
