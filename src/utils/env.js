var VITE_WEBSOCKET_URL =
  import.meta.env.VITE_WEBSOCKET_URL || "ws://localhost:8000/ws";

var VITE_INBOUND_EMAIL_ADDRESS =
  import.meta.env.VITE_INBOUND_EMAIL_ADDRESS ||
  "blog@publish.tarzan.meysam.io";

export { VITE_WEBSOCKET_URL, VITE_INBOUND_EMAIL_ADDRESS };
