# app.py
from flask import Flask, request, jsonify, session
from flask_cors import CORS

app = Flask(__name__)
CORS(app)
app.secret_key = 'supersecret'
users = {}

@app.route('/register', methods=['POST'])
def register():
    data = request.json
    username, password = data['username'], data['password']
    if username in users:
        return jsonify({'error': 'User exists'}), 400
    users[username] = {'password': password, 'tasks': []}
    return jsonify({'message': 'Registered'})

@app.route('/login', methods=['POST'])
def login():
    data = request.json
    username, password = data['username'], data['password']
    user = users.get(username)
    if not user or user['password'] != password:
        return jsonify({'error': 'Invalid credentials'}), 401
    session['user'] = username
    return jsonify({'message': 'Logged in'})

@app.route('/tasks', methods=['GET', 'POST'])
def tasks():
    username = session.get('user')
    if not username:
        return jsonify({'error': 'Not logged in'}), 401
    if request.method == 'POST':
        task = request.json['task']
        users[username]['tasks'].append({'text': task, 'completed': False})
    return jsonify(users[username]['tasks'])

if __name__ == '__main__':
    app.run(debug=True)
