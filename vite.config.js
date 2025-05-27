import { defineConfig, loadEnv } from "vite";
import { ViteMinifyPlugin } from "vite-plugin-minify";
import tailwindcss from "@tailwindcss/vite";
import vue from "@vitejs/plugin-vue";
import { visualizer } from "rollup-plugin-visualizer";
import { compression } from "vite-plugin-compression2";
import path from "path";

export default defineConfig(function configureVite({ mode, command }) {
  var env = loadEnv(mode, process.cwd(), "");
  var isProduction = mode === "production";

  return {
    build: {
      target: "es2015",
      minify: "terser",
      cssMinify: true,
      reportCompressedSize: true,
      chunkSizeWarningLimit: 1000,
      terserOptions: {
        compress: {
          drop_console: isProduction,
          drop_debugger: isProduction,
          pure_funcs: isProduction ? ["console.log", "console.info"] : [],
        },
        format: {
          comments: false,
        },
      },
      rollupOptions: {
        output: {
          manualChunks: function manualChunks(id) {
            if (id.includes("node_modules")) {
              if (id.includes("vue")) {
                return "vue-vendor";
              }
              if (id.includes("tailwindcss") || id.includes("postcss")) {
                return "css-vendor";
              }
              return "vendor";
            }
          },
          entryFileNames: function entryFileNames() {
            return isProduction
              ? "assets/js/[name]-[hash].js"
              : "assets/js/[name].js";
          },
          assetFileNames: function assetFileNames(assetInfo) {
            var extType = assetInfo.name.split(".").at(1);
            if (/png|jpe?g|svg|gif|tiff|bmp|ico/i.test(extType)) {
              extType = "img";
            }
            return isProduction
              ? `assets/${extType}/[name]-[hash][extname]`
              : `assets/${extType}/[name][extname]`;
          },
        },
      },
      sourcemap: isProduction ? "hidden" : true,
    },
    server: {
      host: "::",
      port: env.PORT || 8080,
      strictPort: false,
      open: false,
      cors: true,
      headers: {
        "X-Content-Type-Options": "nosniff",
        "X-Frame-Options": "DENY",
        "X-XSS-Protection": "1; mode=block",
      },
      proxy: {
        "/api": {
          target: env.VITE_API_URL || "http://localhost:8000",
          changeOrigin: true,
          secure: isProduction,
          ws: true,
        },
      },
    },
    preview: {
      host: "::",
      port: env.PREVIEW_PORT || 4173,
      strictPort: false,
      cors: true,
      headers: {
        "Cache-Control": "public, max-age=31536000",
        "X-Content-Type-Options": "nosniff",
        "X-Frame-Options": "DENY",
        "X-XSS-Protection": "1; mode=block",
      },
    },
    resolve: {
      alias: {
        "@": path.resolve(__dirname, "./src"),
        "~": path.resolve(__dirname, "./src"),
        "@components": path.resolve(__dirname, "./src/components"),
        "@assets": path.resolve(__dirname, "./src/assets"),
        "@utils": path.resolve(__dirname, "./src/utils"),
        "@stores": path.resolve(__dirname, "./src/stores"),
        "@views": path.resolve(__dirname, "./src/views"),
      },
    },
    css: {
      devSourcemap: !isProduction,
      preprocessorOptions: {
        scss: {
          additionalData: '@import "@/styles/variables.scss";',
        },
      },
    },
    optimizeDeps: {
      include: ["vue", "@vue/runtime-core"],
      exclude: ["vue-demi"],
    },
    plugins: [
      vue({
        template: {
          compilerOptions: {
            isCustomElement: function isCustomElement(tag) {
              return tag.startsWith("ion-");
            },
          },
        },
      }),
      tailwindcss(),
      isProduction && ViteMinifyPlugin({}),
      isProduction &&
        compression({
          algorithm: "gzip",
          exclude: [/\.(br)$ /, /\.(gz)$/],
        }),
      isProduction &&
        compression({
          algorithm: "brotliCompress",
          exclude: [/\.(br)$ /, /\.(gz)$/],
          filename: "[path][base].br",
        }),
      command === "build" &&
        visualizer({
          filename: "dist/stats.html",
          open: false,
          gzipSize: true,
          brotliSize: true,
        }),
    ].filter(Boolean),
    define: {
      __VUE_OPTIONS_API__: "true",
      __VUE_PROD_DEVTOOLS__: "false",
      __DEV__: !isProduction,
    },
    esbuild: {
      drop: isProduction ? ["console", "debugger"] : [],
      legalComments: "none",
    },
  };
});
