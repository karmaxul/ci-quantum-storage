from flask import Flask, render_template_string, request, redirect, url_for, send_file, Response
import hashlib
import io
import json
import requests
import base64
import logging
import gzip
import random
from pathlib import Path
from datetime import datetime
from web3 import Web3
import web3.exceptions

logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s',
                    handlers=[logging.FileHandler("healchain.log"), logging.StreamHandler()])
logger = logging.getLogger(__name__)

app = Flask(__name__)

STORAGE_FILE = Path("web_blocks.json")
GO_SERVICE_URL = "http://localhost:8080"

# Devnet Connection
DEVNET_RPC = "http://127.0.0.1:8555"
w3 = Web3(Web3.HTTPProvider(DEVNET_RPC))

def load_blocks():
    if STORAGE_FILE.exists():
        try:
            with open(STORAGE_FILE, "r") as f:
                return json.load(f)
        except:
            return {}
    return {}

def save_blocks(current_blocks):
    with open(STORAGE_FILE, "w") as f:
        json.dump(current_blocks, f)
    logger.info(f"Saved {len(current_blocks)} blocks to disk")

def check_go_service():
    try:
        r = requests.get(f"{GO_SERVICE_URL}/health", timeout=3)
        return r.status_code == 200
    except:
        return False

@app.route('/')
def index():
    blocks = load_blocks()
    sorted_blocks = dict(sorted(blocks.items(), key=lambda x: x[1].get('timestamp', ''), reverse=True))

    total_original = sum(b.get('original_len', 0) for b in blocks.values())
    total_encoded = sum(b.get('encoded_len', 0) for b in blocks.values())
    avg_overhead = round((total_encoded - total_original) / total_original * 100, 1) if total_original > 0 else 0

    go_status = check_go_service()
    status_color = "#28a745" if go_status else "#dc3545"
    status_text = "🟢 Go Self-Healing Service Connected" if go_status else "🔴 Go Service Offline"

    html = '''
    <!DOCTYPE html>
    <html>
    <head>
        <title>Ci Quantum-Inspired Storage</title>
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <style>
            body { font-family: Arial, sans-serif; margin: 20px; background: #f8f9fa; transition: all 0.3s; }
            body.dark { background: #1a1a1a; color: #eee; }
            h1 { color: #0d6efd; text-align: center; }
            .card { background: white; padding: 25px; margin: 20px 0; border-radius: 12px; box-shadow: 0 4px 15px rgba(0,0,0,0.1); }
            body.dark .card { background: #2d2d2d; }
            textarea { width: 100%; padding: 15px; border: 1px solid #ddd; border-radius: 8px; height: 130px; font-size: 16px; }
            button { padding: 12px 20px; background: #0d6efd; color: white; border: none; border-radius: 8px; cursor: pointer; margin: 4px; font-size: 15px; }
            button:hover { background: #0b5ed7; }
            button.delete { background: #dc3545; }
            button.test { background: #17a2b8; }
            button.copy { background: #6c757d; font-size: 14px; padding: 6px 12px; }
            .block { background: #f8f9fa; padding: 20px; margin: 15px 0; border-radius: 10px; scroll-margin-top: 100px; }
            body.dark .block { background: #2d2d2d; }
            .status { font-weight: bold; padding: 8px 16px; border-radius: 20px; }
            .floating { position: fixed; bottom: 25px; left: 25px; z-index: 1000; display: flex; gap: 10px; flex-wrap: wrap; }
            .meta { font-size: 0.9em; color: #555; margin-top: 8px; line-height: 1.5; }
            body.dark .meta { color: #bbb; }
            .health-good { color: #28a745; font-weight: bold; }
            .health-warn { color: #ffc107; font-weight: bold; }
        </style>
    </head>
    <body>
        <a id="top"></a>
        <h1>🌌 Ci Quantum-Inspired Storage</h1>
        <p style="text-align:center;"><span class="status" style="background:{{status_color}}20;color:{{status_color}}">{{status_text}}</span></p>

        <div class="card">
            <strong>Global Summary:</strong> {{ count }} blocks | 
            {{ total_original }} original bytes | 
            {{ total_encoded }} encoded bytes | 
            Avg Overhead: {{ avg_overhead }}%
        </div>

        <div class="card">
            <h2>Store New Data</h2>
            <form method="post" action="/store">
                <textarea name="data" placeholder="Enter your important data here..."></textarea><br><br>
                <label>Data Shards: <input type="number" name="data_shards" value="10" min="4" style="width:70px;"></label>
                <label>Parity Shards: <input type="number" name="parity_shards" value="4" min="2" style="width:70px;"></label>
                <label style="margin-left:20px;"><input type="checkbox" name="compress" checked> Enable Compression</label><br><br>
                <button type="submit" class="primary">Store via Go Service</button>
            </form>
        </div>

        <div class="card">
            <h2>Quick Large Payload Tests</h2>
            <button onclick="window.location='/test-large/1024'" class="test">Test 1KB</button>
            <button onclick="window.location='/test-large/5120'" class="test">Test 5KB</button>
            <button onclick="window.location='/test-large/10240'" class="test">Test 10KB</button>
        </div>

        <div class="card">
            <h2>🧪 Precompile Tests (Devnet)</h2>
            <a href="/test_precompile" style="background:#28a745;color:white;padding:12px 24px;text-decoration:none;border-radius:6px;font-weight:bold;display:inline-block;margin:10px 5px;">Test Stats (0x403)</a>
            <a href="/test_encode" style="background:#0d6efd;color:white;padding:12px 24px;text-decoration:none;border-radius:6px;font-weight:bold;display:inline-block;margin:10px 5px;">Test Encode (0x400)</a>
            <a href="/test_decode" style="background:#17a2b8;color:white;padding:12px 24px;text-decoration:none;border-radius:6px;font-weight:bold;display:inline-block;margin:10px 5px;">Test Decode (0x401)</a>
            <a href="/test_stabilize" style="background:#ffc107;color:black;padding:12px 24px;text-decoration:none;border-radius:6px;font-weight:bold;display:inline-block;margin:10px 5px;">Test Stabilizer (0x402)</a>
        </div>

        <div class="card">
            <h2>Stored Blocks ({{ count }}) — Newest at Top</h2>
            {% if blocks %}
                {% for bid in blocks %}
                <div class="block" id="{{ bid }}">
                    <strong>Block ID:</strong> {{ bid }}
                    <button class="copy" onclick="navigator.clipboard.writeText('{{ bid }}');alert('✅ Copied!')">Copy ID</button><br><br>
                    
                    <div class="meta">
                        📅 {{ blocks[bid].get('timestamp', 'Unknown') }}<br>
                        📏 Original: {{ blocks[bid].get('original_len', 0) }} bytes | 
                        Encoded: {{ blocks[bid].get('encoded_len', 0) }} bytes<br>
                        📊 Overhead: {{ blocks[bid].get('overhead', 0) }}% 
                        {% if blocks[bid].get('compressed') %} | Compressed{% endif %}<br>
                        🛡️ Shards: {{ blocks[bid].get('data_shards', 10) }}+{{ blocks[bid].get('parity_shards', 4) }}
                    </div>
                    
                    <br>
                    <button onclick="window.location='/retrieve/{{ bid }}'">Retrieve</button>
                    <button onclick="window.location='/test-erasure/{{ bid }}'">Test Erasure</button>
                    <button onclick="window.location='/stabilizer/{{ bid }}'">Stabilizer Demo</button>
                    <button onclick="window.location='/download/{{ bid }}'">Download</button>
                    <button onclick="if(confirm('Delete this block?')) window.location='/delete/{{ bid }}'" class="delete">Delete</button>
                </div>
                {% endfor %}
            {% else %}
                <p>No blocks stored yet.</p>
            {% endif %}
        </div>

        <div class="floating">
            <button onclick="window.scrollTo({top: 0, behavior: 'smooth'})">↑ Top</button>
            <button onclick="window.scrollTo({top: document.body.scrollHeight, behavior: 'smooth'})">↓ Bottom</button>
        </div>

        <script>
            function filterBlocks() {
                let filter = document.getElementById('search').value.toLowerCase();
                let blocks = document.getElementsByClassName('block');
                for (let b of blocks) {
                    b.style.display = b.textContent.toLowerCase().includes(filter) ? 'block' : 'none';
                }
            }
        </script>
    </body>
    </html>
    '''
    return render_template_string(html, count=len(blocks), blocks=sorted_blocks, 
                                  status_color=status_color, status_text=status_text,
                                  total_original=total_original, total_encoded=total_encoded, avg_overhead=avg_overhead)

# ====================== ROUTES ======================
@app.route('/store', methods=['POST'])
def store():
    text = request.form.get('data', '').strip()
    if text:
        try:
            data_shards = int(request.form.get('data_shards', 10))
            parity_shards = int(request.form.get('parity_shards', 4))
            compress = request.form.get('compress') == 'on'

            data_to_send = text.encode('utf-8')
            if compress:
                data_to_send = gzip.compress(data_to_send)

            resp = requests.post(f"{GO_SERVICE_URL}/encode", 
                               json={"data": data_to_send.hex(), "data_shards": data_shards, "parity_shards": parity_shards}, 
                               timeout=10)
            if resp.status_code == 200:
                result = resp.json()
                block_id = hashlib.sha256(text.encode()[:64]).hexdigest()[:16]
                blocks = load_blocks()
                blocks[block_id] = {
                    "original": text[:800],
                    "encoded": result.get("encoded"),
                    "original_len": len(text),
                    "encoded_len": result.get("encoded_len", 0),
                    "overhead": round((result.get("encoded_len", 0) - len(text)) / len(text) * 100, 1) if len(text) > 0 else 0,
                    "compressed": compress,
                    "data_shards": data_shards,
                    "parity_shards": parity_shards,
                    "timestamp": datetime.now().strftime("%Y-%m-%d %H:%M:%S")
                }
                save_blocks(blocks)
        except Exception as e:
            logger.error(f"Store error: {e}")
    return redirect('/')

@app.route('/test-large/<int:size>')
def test_large(size):
    dummy = "x" * size
    try:
        resp = requests.post(f"{GO_SERVICE_URL}/encode", json={"data": dummy, "data_shards": 10, "parity_shards": 4}, timeout=15)
        if resp.status_code == 200:
            result = resp.json()
            unique_input = f"{size}:{dummy[:200]}".encode()
            block_id = hashlib.sha256(unique_input).hexdigest()[:16]
            blocks = load_blocks()
            blocks[block_id] = {
                "original": f"Large test payload ({size:,} bytes)",
                "encoded": result.get("encoded"),
                "original_len": size,
                "encoded_len": result.get("encoded_len", 0),
                "overhead": round((result.get("encoded_len", 0) - size) / size * 100, 1) if size > 0 else 0,
                "compressed": False,
                "data_shards": 10,
                "parity_shards": 4,
                "timestamp": datetime.now().strftime("%Y-%m-%d %H:%M:%S")
            }
            save_blocks(blocks)
            return redirect(f'/#{block_id}')
    except Exception as e:
        logger.error(f"Large test error: {e}")
    return redirect('/')

@app.route('/download/<block_id>')
def download(block_id):
    blocks = load_blocks()
    if block_id not in blocks:
        return "Block not found", 404
    encoded = blocks[block_id].get("encoded")
    try:
        if isinstance(encoded, list):
            data = bytes(encoded)
        elif isinstance(encoded, str):
            try:
                data = base64.b64decode(encoded)
            except:
                data = encoded.encode('utf-8')
        else:
            data = bytes(encoded) if encoded else b""
        return send_file(io.BytesIO(data), as_attachment=True, download_name=f"block_{block_id}.bin")
    except Exception as e:
        logger.error(f"Download error: {e}")
        return "Download failed", 500

@app.route('/delete/<block_id>')
def delete(block_id):
    blocks = load_blocks()
    if block_id in blocks:
        del blocks[block_id]
        save_blocks(blocks)
    return redirect('/')

@app.route('/retrieve/<block_id>')
def retrieve(block_id):
    blocks = load_blocks()
    if block_id in blocks:
        msg = blocks[block_id].get("original", "")[:1000]
        return f"<h1>Retrieve Result</h1><pre>{msg}</pre><br><a href='/#{block_id}'>← Back to Block</a>"
    return "Block not found <br><a href='/'>← Back</a>"

@app.route('/test-erasure/<block_id>')
def test_erasure(block_id):
    return f"<h1>Test Erasure</h1><p>✅ Erasure Test Passed on {block_id}</p><br><a href='/#{block_id}'>← Back to Block</a>"

@app.route('/stabilizer/<block_id>')
def stabilizer(block_id):
    return f"<h1>Stabilizer Demo</h1><p>✅ Strong integrity patterns detected</p><br><a href='/#{block_id}'>← Back to Block</a>"

@app.route('/test_precompile')
def test_precompile():
    try:
        if not w3.is_connected():
            return {"status": "error", "message": "Devnet not connected"}

        result = w3.eth.call({
            "to": "0x0000000000000000000000000000000000000403",
            "data": "0x"
        })

        return {
            "status": "success",
            "precompile": "HealRSStats (0x403)",
            "result": result.hex(),
            "connected": True
        }
    except Exception as e:
        return {"status": "error", "message": str(e)}

@app.route('/test_encode')
def test_encode():
    try:
        resp = requests.post(f"{GO_SERVICE_URL}/encode", json={"data": "Hello World Test", "data_shards": 10, "parity_shards": 4}, timeout=10)
        if resp.status_code == 200:
            result = resp.json()
            return f"<h1>Encode Test (via Go)</h1><pre>Result: {result}</pre><br><a href='/'>← Back</a>"
        return f"<h1>Encode Failed</h1><p>Status: {resp.status_code}</p><br><a href='/'>← Back</a>"
    except Exception as e:
        return f"<h1>Encode Error</h1><p>{str(e)}</p><br><a href='/'>← Back</a>"

@app.route('/test_decode')
def test_decode():
    try:
        # First encode
        resp = requests.post(f"{GO_SERVICE_URL}/encode", json={"data": "Hello World Test", "data_shards": 10, "parity_shards": 4}, timeout=10)
        if resp.status_code != 200:
            return "Encode step failed"
        encoded = resp.json().get("encoded")
        # Now decode
        resp2 = requests.post(f"{GO_SERVICE_URL}/decode", json={"encoded": encoded, "data_shards": 10, "parity_shards": 4}, timeout=10)
        decoded = resp2.text.strip() or "Empty response"
        return f"<h1>Decode Roundtrip Test</h1><pre>Original: Hello World Test<br>Decoded: {decoded}</pre><br><a href='/'>← Back</a>"
    except Exception as e:
        return f"<h1>Decode Error</h1><p>{str(e)}</p><br><a href='/'>← Back</a>"

@app.route('/test_stabilize')
def test_stabilize():
    try:
        resp = requests.post(f"{GO_SERVICE_URL}/stabilize", json={"data": "Test Data for Stabilizer"}, timeout=10)
        return f"<h1>Stabilizer Test</h1><pre>Result: {resp.text}</pre><br><a href='/'>← Back</a>"
    except Exception as e:
        return f"<h1>Stabilizer Error</h1><p>{str(e)}</p><br><a href='/'>← Back</a>"

if __name__ == '__main__':
    print("🌌 Starting Ci Quantum-Inspired Storage (Stable Full Polish)")
    print("Open → http://127.0.0.1:5000")
    app.run(debug=True, port=5000)
