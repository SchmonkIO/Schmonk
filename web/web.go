package web

import (
	"html/template"
	"net/http"
)

func Load(w http.ResponseWriter, r *http.Request) {
	UITemplate.Execute(w, "ws://"+r.Host+"/ws")
}

var UITemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
	var input = document.getElementById("input");
	var nameInput = document.getElementById("name");
	var xInput = document.getElementById("posx");
	var yInput = document.getElementById("posy");
    var ws;
    var print = function(message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };
    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };
    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
		print("SEND: " + input.value);
		let payload = {
			action: "chat",
			message: input.value
		}
        ws.send(JSON.stringify(payload));
        return false;
	};
	document.getElementById("set").onclick = function(evt) {
        if (!ws) {
            return false;
        }
		print("SEND: " + nameInput.value);
		let payload = {
			action: "join",
			name: nameInput.value
		}
        ws.send(JSON.stringify(payload));
        return false;
	};
	document.getElementById("sendPos").onclick = function(evt) {
        if (!ws) {
            return false;
        }
		let payload = {
			action: "move",
			posx: xInput.value,
			posy: yInput.value
		}
        ws.send(JSON.stringify(payload));
        return false;
	};
	document.getElementById("startGame").onclick = function(evt) {
        if (!ws) {
            return false;
        }
		let payload = {
			action: "startGame"
		}
        ws.send(JSON.stringify(payload));
        return false;
	};
	document.getElementById("stopGame").onclick = function(evt) {
        if (!ws) {
            return false;
        }
		let payload = {
			action: "stopGame"
		}
        ws.send(JSON.stringify(payload));
        return false;
    };
    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Schmonk Web Testing Interface, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<button id="startGame">Start Game</button>
<button id="stopGame">Stop Game</button>
<p><input id="name" type="text" value="Tester">
<button id="set">Set Name</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send Message</button>
<p><input id="posx" type="text" value="12">
<input id="posy" type="text" value="12">
<button id="sendPos">Move</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))
