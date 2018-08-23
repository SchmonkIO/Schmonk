/*
 * schmonk client implementation
 * created: 2018-08-20
 */



class EnemyPlayer {
  constructor(opts) {
    this.id = opts._id;
    this.name = opts.name;
    this.color = opts.color;
    this.posx = opts.posx;
    this.posy = opts.posy;

    this.entityInstance = null; // crafty instance 
    this.textEntityInstance = null; //

    // global game engine
    this.engine = opts.engine;
    this.spawn();
  }

  spawn() {
    let Engine = this.engine;
    this.entityInstance = Engine.e("2D, DOM, Color, Solid")
      .attr({x: this.posx, y: this.posy, w: 50, h:50})
      .color(this.color);
    this.textEntityInstance = Engine.e("2D, DOM, Text")
      .attr({ x: this.posx - 20, y: this.posy})
      .textAlign('center')
      .textFont({ size: '14px', weight: 'bold' })
      .text(this.name);
  }
  destroy() {
    this.entityInstance.destroy();
    this.textEntityInstance.destroy();
  }
  move(posx, posy) {
    this.entityInstance.x = posx;
    this.entityInstance.y = posy;
    this.textEntityInstance.x = posx;
    this.textEntityInstance.y = posy - 20;
  }

}

/*******************************************/


class SchmonkClient {
  constructor(opts) {
    this.domElId = opts.domElId;
    this.statusDomElId = opts.statusDomElId;
    this.socketUrl = opts.socketUrl;
    this.playerName = opts.playerName || "UnnamedPlayer";

    this.connection = null;
    this.engine = null;

    // enemy player instances
    this.enemys = [];

    // spin up engine, connection and scene
    this.initEngine(this.domElId);
    this.initConnection(this.socketUrl, this.playerName)
    this.initScene();

  }

  initEngine(domElId) {
    this.engine = Crafty.init(1000,500, document.getElementById(domElId));
  }

  initConnection(socketUrl, playerName) {
    this.connection = new WebSocket(socketUrl);
    this.connection.onopen = () => {
      document.getElementById(this.statusDomElId).innerHTML = "Connected"
      console.log("connected");
      let payload = {action: "join", name: playerName};
      this.connection.send(JSON.stringify(payload));
      this.spawnLocalPlayer();
    }
    this.connection.onerror = (error) => {
      console.log(this.statusDomElId);
      document.getElementById(this.statusDomElId).innerHTML = "Can't connect to game server :/";
    }
    this.connection.onmessage = (ev) => {
      let data = JSON.parse(ev.data);
      if(data && data.action) { 
        // seems like we received valid json w/ action
        switch(data.action) {
          case "tick":
            this.handleTickUpdate(data);
          case "chatMessage":
            // TODO: handle chat messages
            break;
          default:
            // ..we don't care :D
            break;
        }
      }
    };
  }

  initScene() {
    let Engine = this.engine;

    // TODO: main menu scene

    Engine.defineScene("play", () => {     
      Engine.background("#dee2e6");

      // the black bottom bar
      Engine.e('Floor, 2D, Canvas, Color, Solid, WiredHitBox')
        .attr({x: 0, y: 450, w: 1000, h: 50})
        .color('black');
    });

    // ..just straight into play scene
    Engine.enterScene("play");
  }

  spawnLocalPlayer() {
    let Engine = this.engine;
    let connection = this.connection;
    // our player entity
    let player = Engine.e("2D, DOM, Color, Twoway, Collision, Gravity, WiredHitBox")
      .attr({x: 0, y: 0, w: 50, h:50})
      .color("#cc5de8")
      .twoway(200, 800)
      .gravityConst(2500)
      .gravity('Floor')
      .bind("Move", () => {
        let payload = {action: "move", posx: player.x.toString(), posy: player.y.toString()};
        connection.send(JSON.stringify(payload));
      })
  }

  handleTickUpdate(data) {
    // check for new enemys and update existing ones
    Object.keys(data.players).forEach((playerId) => {
      let player = data.players[playerId];
      if(this.enemys[playerId]) {
        // instance is there, so just update the pos
        this.enemys[playerId].move(player.posx, player.posy);
      } else {
        if(playerId === data._id) { // dont render if its us
          return;
        }
        // he's new so bring him to life .d
        this.enemys[playerId] = new EnemyPlayer({
          _id: player._id,
          name: player.name,
          posx: player.posx,
          posy: player.psy,
          color: player.color,
          engine: this.engine
        });
      }

   });

   // check for disconnected players
   Object.keys(this.enemys).forEach((playerId) => {
     if(!data.players[playerId]) {
       this.enemys[playerId].destroy(); // kill him!
       delete data.players[playerId]; // remove from player register
     }
   })
  }
}