import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    hmr: {
      clientPort: 8000,
    },
    watch: {
      usePolling: true,
    },
    host: '0.0.0.0',
    port: 5173,
    strictPort: true,
  },
});
