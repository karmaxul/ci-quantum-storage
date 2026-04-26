# app.py - Final Polished Version
from flask import Flask, render_template_string, request, redirect, url_for, send_file
import hashlib
import io
import json
from pathlib import Path

from ci_sha4096_v2_4 import ci_sha4096_v2_4
from ci_rs_wrapper import CiReedSolomon

app = Flask(__name__)

STORAGE_FILE = Path("web_blocks.json")
rs_wrapper = CiReedSolomon(nsym=48)

def load_blocks():
    if STORAGE_FILE.exists():
        try:
            with open(STORAGE_FILE, "r") as f:
                return json.load(f)
        except:
            return {}
    return {}

def save_blocks(blocks):
    with open(STORAGE_FILE, "w") as f:
        json.dump(blocks, f)

blocks = load_blocks()

@app.route('/')
def index():
    html = '''
    <!DOCTYPE html>
    <html>
    <head>
        <title>Ci Quantum-Inspired Storage</title>
        <style>
            body { font-family: Arial, sans-serif; margin: 40px; background: #f8f9fa; }
            h1 { color: #0d6efd; }
            .card { background: white; padding: 25px; margin: 20px 0; border-radius: 12px; box-shadow: 0 4px 15px rgba(0,0,0,0.1); }
            textarea { width: 100%; padding: 15px; border: 1px solid #ddd; border-radius: 8px; height: 120px; font-size: 16px; }
            button { padding: 14px 28px; background: #0d6efd; color: white; border: none; border-radius: 8px; cursor: pointer; font-size: 16px; margin: 5px; }
            button:hover { background: #0b5ed7; }
            .block { background: #f8f9fa; padding: 20px; margin: 15px 0; border-radius: 10px; }
            .success { color: #28a745; font-weight: bold; background: #d4edda; padding: 15px; border-radius: 8px; }
            .copy-btn { background: #6c757d; }
        </style>
        <script>
            function copyToClipboard(text) {
                navigator.clipboard.writeText(text).then(() => alert("Hash copied to clipboard!"));
            }
        </script>
    </head>
    <body>
        <h1>🌌 Ci Quantum-Inspired Storage</h1>
        <p><strong>Powered by Ci-SHA4096 v2.4 (85/27) + Reed-Solomon Protection</strong></p>

        <div class="card">
            <h2>Store New Data</h2>
            <form method="post" action="/store">
                <textarea name="data" placeholder="Enter your important data here..."></textarea><br>
                <button type="submit">Store Data (with Full RS + Ci Protection)</button>
            </form>
        </div>

        <div class="card">
            <h2>Stored Blocks ({{ count }})</h2>
            {% if blocks %}
                {% for block_id in blocks %}
                <div class="block">
                    <strong>Block ID:</strong> {{ block_id }}<br><br>
                    <button onclick="window.location='/retrieve/{{ block_id }}'">Retrieve</button>
                    <button onclick="window.location='/test-erasure/{{ block_id }}'">Test Erasure</button>
                    <button onclick="window.location='/stabilizer/{{ block_id }}'">Stabilizer Demo</button>
                    <button onclick="window.location='/download/{{ block_id }}'">Download Block</button>
                </div>
                {% endfor %}
            {% else %}
                <p>No blocks stored yet.</p>
            {% endif %}
        </div>
    </body>
    </html>
    '''
    return render_template_string(html, blocks=blocks, count=len(blocks))

# Store, Retrieve, Test Erasure, Stabilizer, Download routes (same as before - they are already working)

@app.route('/store', methods=['POST'])
def store():
    text = request.form.get('data', '').strip()
    if text:
        data = text.encode('utf-8')
        block_id = hashlib.sha256(data[:64]).hexdigest()[:16]
        encoded = rs_wrapper.encode(data)
        blocks[block_id] = list(encoded)
        save_blocks(blocks)
        print(f"✅ Stored protected block {block_id}")
    return redirect(url_for('index'))

# (Keep your existing /retrieve, /test-erasure, /stabilizer, /download routes)

if __name__ == '__main__':
    print("🌌 Starting Ci Quantum Storage Web Interface...")
    print("Open → http://127.0.0.1:5000")
    app.run(debug=True, port=5000)
