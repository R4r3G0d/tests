from flask import Flask, render_template, request, abort, send_file
from os import getcwd
from os.path import exists

app = Flask(__name__)
working_dir = getcwd()


@app.route('/')
def root():
    return render_template("index.html")


@app.route('/static')
def get_static():
    file_path = f"{working_dir}/static/{request.args.get('file')}"

    if exists(file_path):
        return send_file(file_path)
    else:
        abort(404)


def main():
    app.run()


if __name__ == '__main__':
    main()
