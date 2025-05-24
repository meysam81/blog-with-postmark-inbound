import { defineConfig, loadEnv } from "vite";
import { ViteMinifyPlugin } from "vite-plugin-minify";
import tailwindcss from "@tailwindcss/vite";
import path from "path";

export default defineConfig(function configureVite({ mode }) {
  var env = loadEnv(mode, process.cwd(), "");
  return {
    build: {
      minify: true,
    },
    server: {
      port: env.PORT || 8080,
      proxy: {
        "/api": "http://localhost:8000",
      },
    },
    resolve: {
      alias: {
        "@": path.resolve(__dirname, "./src"),
      },
    },
    plugins: [tailwindcss(), ViteMinifyPlugin({})],
  };
});
