from flask import Flask
from flask import Flask, make_response, render_template
import os, platform

app = Flask(__name__)

#? BACK

kontinu_msg=os.getenv("MSG"," 🐳 Hello")

# Use environment variables
#import extras.load_env

# load json config files
#import extras.config_file

# "/health"
import extras.health

# "/kontinu" in case
import extras.kontinu

# redis conn
from extras.redis_connect import nocache, get_hit_count

@app.route('/')
@nocache
def root():
    count = get_hit_count()
    host=platform.node()
    foo=os.getenv('FOO', 'unset')
    REDIS_HOST=os.getenv("REDIS_HOST",'')
    DOCKER_SERVICE_NAME=os.getenv('DOCKER_SERVICE_NAME', host)
    if "{{" in DOCKER_SERVICE_NAME or "}}" in DOCKER_SERVICE_NAME:
        DOCKER_SERVICE_NAME="N/A"
    #print(f"Getting visits! {count}")
    global kontinu_msg
    try:
        kontinu_msg=kontinu_msg+extras.kontinu.get_next_msg()
    except:
        pass

    return render_template('index.html',visit_counts=count, hostname=host, DOCKER_SERVICE_NAME=DOCKER_SERVICE_NAME, FOO=foo , redis_host=REDIS_HOST, greeting=kontinu_msg)



if __name__ == '__main__':
    print("==== Starting Python Server ====")
    app.run(host='0.0.0.0',port=8080)
