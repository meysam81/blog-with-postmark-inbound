var protocol = window.location.protocol === "https:" ? "wss" : "ws";
var hostname = window.location.hostname;
var port = window.location.port

var VITE_WEBSOCKET_URL =
  import.meta.env.VITE_WEBSOCKET_URL || `${protocol}://${hostname}:${port}/ws`;

var VITE_INBOUND_EMAIL_ADDRESS =
  import.meta.env.VITE_INBOUND_EMAIL_ADDRESS ||
  "blog@publish.tarzan.meysam.io";

export { VITE_WEBSOCKET_URL, VITE_INBOUND_EMAIL_ADDRESS };
